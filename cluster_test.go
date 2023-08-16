package test

import (
	"crypto/tls"
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/terraform"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// TestCluster is an automated test which:
// 1. Deploys all infra defined in terraform
// 2. Checks the number of nodes in the cluster
// 3. Checks External access to Traefik via an external LoadBalancer
// 4. Checks there is a single nginx pod running in the default namespace
func TestCluster(t *testing.T) {
	// We use Gomega to make assertions, this is not required and can be replaced by the standard testing lib
	g := NewGomegaWithT(t)

	// Terraform options to deploy resrouces from the current dir
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./",
	})
	// Ensure that resources are destroyed at the end of the test
	defer terraform.Destroy(t, terraformOptions)

	// Deploy all infra defined in terraform
	terraform.InitAndApply(t, terraformOptions)

	// Use the local Kubeconfig file to connect to the cluster for testing
	clusterOptions := k8s.NewKubectlOptions("", "./kubeconfig", "default")

	// Check the number of nodes in the cluster
	g.Expect(len(k8s.GetNodes(t, clusterOptions))).Should(Equal(3))
	g.Expect(k8s.AreAllNodesReady(t, clusterOptions)).Should(BeTrue())

	// Check External access to Traefik
	clusterOptions.Namespace = "traefik"
	svc := k8s.GetService(t, clusterOptions, "traefik")
	g.Expect(len(svc.Status.LoadBalancer.Ingress)).Should(Equal(1))

	svcAddress := svc.Status.LoadBalancer.Ingress[0].IP
	httpCode, _ := http_helper.HttpGet(t, fmt.Sprintf("http://%s", svcAddress), &tls.Config{})
	// We expect a 404 because we have not deployed any services / ingress to
	// the cluster. A failure to connect would indicate a problem with the
	// resources deployed
	g.Expect(httpCode).Should(Equal(404))

	httpCode, _ = http_helper.HttpGet(t, fmt.Sprintf("https://%s", svcAddress), &tls.Config{InsecureSkipVerify: true})
	// We expect a 404 because we have not deployed any services / ingress to
	// the cluster. A failure to connect would indicate a problem with the
	// resources deployed
	g.Expect(httpCode).Should(Equal(404))

	// Check there is a single nginx pod running in the default namespace
	clusterOptions.Namespace = "default"
	nginxDeployment := k8s.ListPods(t, clusterOptions, metav1.ListOptions{LabelSelector: "nginx=nginx"})
	g.Expect(len(nginxDeployment)).Should(Equal(1))
}

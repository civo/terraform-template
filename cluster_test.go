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

func TestCluster(t *testing.T) {
	g := NewGomegaWithT(t)
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "./",
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	clusterOptions := k8s.NewKubectlOptions("", "./kubeconfig", "default")

	// Check nodes
	g.Expect(len(k8s.GetNodes(t, clusterOptions))).Should(Equal(3))
	g.Expect(k8s.AreAllNodesReady(t, clusterOptions)).Should(BeTrue())

	// Check External access to Traefik
	clusterOptions.Namespace = "traefik"
	svc := k8s.GetService(t, clusterOptions, "traefik")
	g.Expect(len(svc.Status.LoadBalancer.Ingress)).Should(Equal(1))

	svcAddress := svc.Status.LoadBalancer.Ingress[0].IP
	httpCode, _ := http_helper.HttpGet(t, fmt.Sprintf("http://%s", svcAddress), &tls.Config{})
	g.Expect(httpCode).Should(Equal(404))

	httpCode, _ = http_helper.HttpGet(t, fmt.Sprintf("https://%s", svcAddress), &tls.Config{InsecureSkipVerify: true})
	g.Expect(httpCode).Should(Equal(404))

	// Check Nginix Deployment
	clusterOptions.Namespace = "default"
	nginxDeployment := k8s.ListPods(t, clusterOptions, metav1.ListOptions{LabelSelector: "nginx=nginx"})
	g.Expect(len(nginxDeployment)).Should(Equal(1))
}

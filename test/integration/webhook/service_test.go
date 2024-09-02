/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package webhook

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	coreapi "github.com/inftyai/llmaz/api/core/v1alpha1"
	inferenceapi "github.com/inftyai/llmaz/api/inference/v1alpha1"
	"github.com/inftyai/llmaz/test/util"
	"github.com/inftyai/llmaz/test/util/wrapper"
)

var _ = ginkgo.Describe("service default and validation", func() {
	// Each test runs in a separate namespace.
	var ns *corev1.Namespace

	ginkgo.BeforeEach(func() {
		// Create test namespace before each test.
		ns = &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "test-ns-",
			},
		}
		gomega.Expect(k8sClient.Create(ctx, ns)).To(gomega.Succeed())
	})
	ginkgo.AfterEach(func() {
		gomega.Expect(k8sClient.Delete(ctx, ns)).To(gomega.Succeed())
	})

	type testValidatingCase struct {
		service func() *inferenceapi.Service
		failed  bool
	}

	ginkgo.DescribeTable("test validating",
		func(tc *testValidatingCase) {
			if tc.failed {
				gomega.Expect(k8sClient.Create(ctx, tc.service())).To(gomega.HaveOccurred())
			} else {
				gomega.Expect(k8sClient.Create(ctx, tc.service())).To(gomega.Succeed())
			}
		},
		ginkgo.Entry("normal Service creation", &testValidatingCase{
			service: func() *inferenceapi.Service {
				return util.MockASampleService(ns.Name)
			},
			failed: false,
		}),
		ginkgo.Entry("invalid name", &testValidatingCase{
			service: func() *inferenceapi.Service {
				return wrapper.MakeService("service-0.5b", ns.Name).WorkerTemplate().Obj()
			},
			failed: true,
		}),
		ginkgo.Entry("model-runner container doesn't exist", &testValidatingCase{
			service: func() *inferenceapi.Service {
				return wrapper.MakeService("service-llama3-8b", ns.Name).
					ModelsClaim([]string{"llama3-8b"}, coreapi.Standard, nil).
					WorkerTemplate().
					ContainerName("model-runner-fake").
					Obj()
			},
			failed: true,
		}),
	)
})

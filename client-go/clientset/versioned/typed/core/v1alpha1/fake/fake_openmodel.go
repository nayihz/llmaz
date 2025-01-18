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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/inftyai/llmaz/api/core/v1alpha1"
	corev1alpha1 "github.com/inftyai/llmaz/client-go/applyconfiguration/core/v1alpha1"
	typedcorev1alpha1 "github.com/inftyai/llmaz/client-go/clientset/versioned/typed/core/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeOpenModels implements OpenModelInterface
type fakeOpenModels struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.OpenModel, *v1alpha1.OpenModelList, *corev1alpha1.OpenModelApplyConfiguration]
	Fake *FakeLlmazV1alpha1
}

func newFakeOpenModels(fake *FakeLlmazV1alpha1, namespace string) typedcorev1alpha1.OpenModelInterface {
	return &fakeOpenModels{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.OpenModel, *v1alpha1.OpenModelList, *corev1alpha1.OpenModelApplyConfiguration](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("openmodels"),
			v1alpha1.SchemeGroupVersion.WithKind("OpenModel"),
			func() *v1alpha1.OpenModel { return &v1alpha1.OpenModel{} },
			func() *v1alpha1.OpenModelList { return &v1alpha1.OpenModelList{} },
			func(dst, src *v1alpha1.OpenModelList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.OpenModelList) []*v1alpha1.OpenModel { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.OpenModelList, items []*v1alpha1.OpenModel) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package seed

import (
	"testing"

	kubermaticv1 "github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestValidate(t *testing.T) {
	fakeProviderSpec := kubermaticv1.DatacenterSpec{
		Fake: &kubermaticv1.DatacenterSpecFake{},
	}

	testCases := []struct {
		name             string
		existingSeeds    map[string]*kubermaticv1.Seed
		seedToValidate   *kubermaticv1.Seed
		existingClusters []runtime.Object
		isDelete         bool
		errExpected      bool
	}{
		{
			name:           "Adding an empty seed should be possible",
			seedToValidate: &kubermaticv1.Seed{},
		},
		{
			name: "Adding a seed with a single datacenter and valid provider should succeed",
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "new-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"dc1": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
		},
		{
			name: "No changes, no error",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"dc1": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
		},
		{
			name: "Clusters from other seeds should have no effect on new empty seeds",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"europe-west3-c": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "europe-west3-c",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"do-fra1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			existingClusters: []runtime.Object{
				&kubermaticv1.Cluster{
					Spec: kubermaticv1.ClusterSpec{
						Cloud: kubermaticv1.CloudSpec{
							DatacenterName: "do-fra1",
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "asia-south1-a",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{},
				},
			},
		},
		{
			name: "Clusters from other seeds should have no effect when deleting seeds",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"europe-west3-c": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "europe-west3-c",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"do-fra1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
				"asia-south1-a": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "asia-south1-a",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"aws-asia-south1-a": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			existingClusters: []runtime.Object{
				&kubermaticv1.Cluster{
					Spec: kubermaticv1.ClusterSpec{
						Cloud: kubermaticv1.CloudSpec{
							DatacenterName: "do-fra1",
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "asia-south1-a",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"aws-asia-south1-a": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
			isDelete: true,
		},
		{
			name: "Adding new datacenter should be possible",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"dc1": {
							Spec: fakeProviderSpec,
						},
						"dc2": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
		},
		{
			name: "Should be able to remove unused datacenters from a seed",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{},
				},
			},
		},
		{
			name: "Datacenters must have a provider defined",
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "myseed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {},
					},
				},
			},
			errExpected: true,
		},
		{
			name: "Datacenters cannot have multiple providers",
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "myseed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {
							Spec: kubermaticv1.DatacenterSpec{
								AWS:   &kubermaticv1.DatacenterSpecAWS{},
								Azure: &kubermaticv1.DatacenterSpecAzure{},
							},
						},
					},
				},
			},
			errExpected: true,
		},
		{
			name: "It should not be possible to change a datacenter's provider",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"dc1": {
							Spec: kubermaticv1.DatacenterSpec{
								AWS: &kubermaticv1.DatacenterSpecAWS{},
							},
						},
					},
				},
			},
			errExpected: true,
		},
		{
			name: "Datacenter names are unique across all seeds",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"in-use": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "new-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"in-use": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
			errExpected: true,
		},
		{
			name: "Cannot remove datacenters that are used by clusters",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			existingClusters: []runtime.Object{
				&kubermaticv1.Cluster{
					Spec: kubermaticv1.ClusterSpec{
						Cloud: kubermaticv1.CloudSpec{
							DatacenterName: "dc1",
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
			},
			errExpected: true,
		},
		{
			name:           "Shuld be able to delete empty seeds",
			seedToValidate: &kubermaticv1.Seed{},
			isDelete:       true,
		},
		{
			name: "Shuld be able to delete seeds with no used datacenters",
			existingSeeds: map[string]*kubermaticv1.Seed{
				"existing-seed": {
					ObjectMeta: metav1.ObjectMeta{
						Name: "existing-seed",
					},
					Spec: kubermaticv1.SeedSpec{
						Datacenters: map[string]kubermaticv1.Datacenter{
							"dc1": {
								Spec: fakeProviderSpec,
							},
						},
					},
				},
			},
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "existing-seed",
				},
				Spec: kubermaticv1.SeedSpec{
					Datacenters: map[string]kubermaticv1.Datacenter{
						"dc1": {
							Spec: fakeProviderSpec,
						},
					},
				},
			},
			isDelete: true,
		},
		{
			name: "Cannot delete a seed when there are still clusters left",
			seedToValidate: &kubermaticv1.Seed{
				ObjectMeta: metav1.ObjectMeta{
					Name: "myseed",
				},
			},
			existingClusters: []runtime.Object{
				&kubermaticv1.Cluster{},
			},
			isDelete:    true,
			errExpected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sv := &seedValidator{
				listOpts: &ctrlruntimeclient.ListOptions{},
			}

			err := sv.validate(tc.seedToValidate,
				fakectrlruntimeclient.NewFakeClientWithScheme(scheme.Scheme, tc.existingClusters...),
				tc.existingSeeds, tc.isDelete)

			if (err != nil) != tc.errExpected {
				t.Fatalf("Expected err: %t, but got err: %v", tc.errExpected, err)
			}
		})
	}

}

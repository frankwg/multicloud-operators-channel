// Copyright 2019 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package webhook

import (
	"context"
	"os"
	"path/filepath"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	admissionv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	mgr "sigs.k8s.io/controller-runtime/pkg/manager"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	chv1 "github.com/open-cluster-management/multicloud-operators-channel/pkg/apis/apps/v1"
)

var _ = Describe("test channel validation logic", func() {
	Context("given an exist namespace channel in a namespace", func() {
		var (
			chkey  = types.NamespacedName{Name: "ch1", Namespace: "default"}
			chnIns = chv1.Channel{
				ObjectMeta: metav1.ObjectMeta{
					Name:      chkey.Name,
					Namespace: chkey.Namespace},
				Spec: chv1.ChannelSpec{
					Type:     chv1.ChannelType(chv1.ChannelTypeNamespace),
					Pathname: chkey.Namespace,
				},
			}
		)

		BeforeEach(func() {
			// Create the Channel object and expect the Reconcile
			Expect(k8sClient.Create(context.TODO(), chnIns.DeepCopy())).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), &chnIns)).Should(Succeed())
		})

		It("should create git channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeGit
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("should not create 2nd  namespace channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("should not create 2nd objectbucket channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeObjectBucket
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("should not create 2nd  helm channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeHelmRepo
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})
	})

	Context("given exist git channels in a namespace", func() {
		var (
			chkey  = types.NamespacedName{Name: "ch1", Namespace: "default"}
			chnIns = chv1.Channel{
				ObjectMeta: metav1.ObjectMeta{
					Name:      chkey.Name,
					Namespace: chkey.Namespace},
				Spec: chv1.ChannelSpec{
					Type:     chv1.ChannelType(chv1.ChannelTypeGit),
					Pathname: chkey.Namespace,
				},
			}
		)

		BeforeEach(func() {
			// Create the Channel object and expect the Reconcile
			Expect(k8sClient.Create(context.TODO(), chnIns.DeepCopy())).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), &chnIns)).Should(Succeed())
		})

		It("should create 2nd git channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeGit
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("should create 2nd git channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = "GitHub"
			dupChn.SetName("dup-chn1-1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("should create 2nd  namespace channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeNamespace
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("should create 2nd objectbucket channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeObjectBucket
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()

		})

		It("should create 2nd  helm channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeHelmRepo
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})
	})

	Context("given an exist objectbucket channel in a namespace", func() {
		var (
			chkey  = types.NamespacedName{Name: "ch1", Namespace: "default"}
			chnIns = chv1.Channel{
				ObjectMeta: metav1.ObjectMeta{
					Name:      chkey.Name,
					Namespace: chkey.Namespace},
				Spec: chv1.ChannelSpec{
					Type:     chv1.ChannelType(chv1.ChannelTypeObjectBucket),
					Pathname: chkey.Namespace,
				},
			}
		)

		BeforeEach(func() {
			// Create the Channel object and expect the Reconcile
			Expect(k8sClient.Create(context.TODO(), chnIns.DeepCopy())).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), &chnIns)).Should(Succeed())
		})

		It("should create 2nd git channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeGit
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("shouldn't create 2nd  namespace channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeNamespace
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("shouldn't create 2nd objectbucket channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeObjectBucket
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("shouldn't create 2nd  helm channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeHelmRepo
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})
	})

	Context("given an exist helm channel in a namespace", func() {
		var (
			chkey  = types.NamespacedName{Name: "ch1", Namespace: "default"}
			chnIns = chv1.Channel{
				ObjectMeta: metav1.ObjectMeta{
					Name:      chkey.Name,
					Namespace: chkey.Namespace},
				Spec: chv1.ChannelSpec{
					Type:     chv1.ChannelType(chv1.ChannelTypeHelmRepo),
					Pathname: chkey.Namespace,
				},
			}
		)

		BeforeEach(func() {
			// Create the Channel object and expect the Reconcile
			Expect(k8sClient.Create(context.TODO(), chnIns.DeepCopy())).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			Expect(k8sClient.Delete(context.TODO(), &chnIns)).Should(Succeed())
		})

		It("should create 2nd git channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeGit
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).Should(Succeed())
			defer func() {
				Expect(k8sClient.Delete(context.TODO(), dupChn)).Should(Succeed())
			}()
		})

		It("shouldn't create 2nd  namespace channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeNamespace
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("shouldn't create 2nd objectbucket channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeObjectBucket
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})

		It("shouldn't create 2nd  helm channel", func() {
			dupChn := chnIns.DeepCopy()
			dupChn.Spec.Type = chv1.ChannelTypeHelmRepo
			dupChn.SetName("dup-chn1")

			Expect(k8sClient.Create(context.TODO(), dupChn)).ShouldNot(Succeed())
		})
	})

	// somehow this test only fail on travis
	// make sure this one runs at the end, otherwise, we might register this
	// webhook before the default one, which cause unexpected results.
	PContext("given a k8s env, it create svc and validating webhook config", func() {
		var (
			lMgr    mgr.Manager
			certDir string
			testNs  string
			caCert  []byte
			err     error
			sstop   chan struct{}
		)

		It("should create a service and ValidatingWebhookConfiguration", func() {
			lMgr, err = mgr.New(testEnv.Config, mgr.Options{MetricsBindAddress: "0"})
			Expect(err).Should(BeNil())

			sstop = make(chan struct{})
			defer close(sstop)
			go func() {
				Expect(lMgr.Start(sstop)).Should(Succeed())
			}()

			certDir = filepath.Join(os.TempDir(), "k8s-webhook-server", "serving-certs")
			testNs = "default"
			os.Setenv("POD_NAMESPACE", testNs)

			caCert, err = GenerateWebhookCerts(certDir)
			Expect(err).Should(BeNil())
			validatorName := "test-validator"
			wbhSvcNm := "ch-wbh-svc"
			WireUpWebhookSupplymentryResource(lMgr, stop, wbhSvcNm, validatorName, certDir, caCert)

			ns, err := findEnvVariable(podNamespaceEnvVar)
			Expect(err).Should(BeNil())

			time.Sleep(3 * time.Second)
			wbhSvc := &corev1.Service{}
			svcKey := types.NamespacedName{Name: wbhSvcNm, Namespace: ns}
			Expect(lMgr.GetClient().Get(context.TODO(), svcKey, wbhSvc)).Should(Succeed())
			defer func() {
				Expect(lMgr.GetClient().Delete(context.TODO(), wbhSvc)).Should(Succeed())
			}()

			wbhCfg := &admissionv1.ValidatingWebhookConfiguration{}
			cfgKey := types.NamespacedName{Name: validatorName}
			Expect(lMgr.GetClient().Get(context.TODO(), cfgKey, wbhCfg)).Should(Succeed())

			defer func() {
				Expect(lMgr.GetClient().Delete(context.TODO(), wbhCfg)).Should(Succeed())
			}()
		})
	})
})

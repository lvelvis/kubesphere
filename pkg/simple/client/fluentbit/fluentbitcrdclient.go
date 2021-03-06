/*
Copyright 2018 The KubeSphere Authors.
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
package fluentbitclient

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

const (
	CRDPlural   string = "fluentbits"
	CRDGroup    string = "logging.kubesphere.io"
	CRDVersion  string = "v1alpha1"
	FullCRDName string = CRDPlural + "." + CRDGroup
)

// FluentBitList auto generated by the sdk
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FluentBitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []FluentBit `json:"items"`
}

// FluentBit auto generated by the sdk
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type FluentBit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              FluentBitSpec   `json:"spec"`
	Status            FluentBitStatus `json:"status,omitempty"`
}

// FluentBitSpec holds the spec for the operator
type FluentBitSpec struct {
	Service  []Plugin `json:"service"`
	Input    []Plugin `json:"input"`
	Filter   []Plugin `json:"filter"`
	Output   []Plugin `json:"output"`
	Settings []Plugin `json:"settings"`
}

// FluentBitStatus holds the status info for the operator
type FluentBitStatus struct {
	// Fill me
}

// Plugin struct for fluent-bit plugins
type Plugin struct {
	Type       string      `json:"type" description:"output plugin type, eg. fluentbit-output-es"`
	Name       string      `json:"name" description:"output plugin name, eg. fluentbit-output-es"`
	Parameters []Parameter `json:"parameters" description:"output plugin configuration parameters"`
}

// Fluent-bit output plugins
type OutputPlugin struct {
	Plugin
	Id         string    `json:"id,omitempty" description:"output uuid"`
	Enable     bool      `json:"enable" description:"active status, one of true, false"`
	Updatetime time.Time `json:"updatetime,omitempty" description:"last updatetime"`
}

// Parameter generic parameter type to handle values from different sources
type Parameter struct {
	Name      string     `json:"name" description:"configuration parameter key, eg. Name. refer to Fluent bit's Output Plugins Section for more configuration parameters."`
	ValueFrom *ValueFrom `json:"valueFrom,omitempty"`
	Value     string     `json:"value" description:"configuration parameter value, eg. es. refer to Fluent bit's Output Plugins Section for more configuration parameters."`
}

// ValueFrom generic type to determine value origin
type ValueFrom struct {
	SecretKeyRef KubernetesSecret `json:"secretKeyRef"`
}

// KubernetesSecret is a ValueFrom type
type KubernetesSecret struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	Namespace string `json:"namespace"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentBit) DeepCopyInto(out *FluentBit) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentBit.
func (in *FluentBit) DeepCopy() *FluentBit {
	if in == nil {
		return nil
	}
	out := new(FluentBit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FluentBit) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentBitList) DeepCopyInto(out *FluentBitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FluentBit, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentBitList.
func (in *FluentBitList) DeepCopy() *FluentBitList {
	if in == nil {
		return nil
	}
	out := new(FluentBitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FluentBitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentBitSpec) DeepCopyInto(out *FluentBitSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentBitSpec.
func (in *FluentBitSpec) DeepCopy() *FluentBitSpec {
	if in == nil {
		return nil
	}
	out := new(FluentBitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FluentBitStatus) DeepCopyInto(out *FluentBitStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FluentBitStatus.
func (in *FluentBitStatus) DeepCopy() *FluentBitStatus {
	if in == nil {
		return nil
	}
	out := new(FluentBitStatus)
	in.DeepCopyInto(out)
	return out
}

// Create a  Rest client with the new CRD Schema
var SchemeGroupVersion = schema.GroupVersion{Group: CRDGroup, Version: CRDVersion}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&FluentBit{},
		&FluentBitList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func NewFluentbitCRDClient(cfg *rest.Config) (*rest.RESTClient, *runtime.Scheme, error) {
	scheme := runtime.NewScheme()
	SchemeBuilder := runtime.NewSchemeBuilder(addKnownTypes)
	if err := SchemeBuilder.AddToScheme(scheme); err != nil {
		return nil, nil, err
	}
	config := *cfg
	config.GroupVersion = &SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.NewCodecFactory(runtime.NewScheme()).WithoutConversion()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, nil, err
	}
	return client, scheme, nil
}

// This file implement all the (CRUD) client methods we need to access our CRD object

func CrdClient(cl *rest.RESTClient, scheme *runtime.Scheme, namespace string) *crdclient {
	return &crdclient{cl: cl, ns: namespace, plural: CRDPlural,
		codec: runtime.NewParameterCodec(scheme)}
}

type crdclient struct {
	cl     *rest.RESTClient
	ns     string
	plural string
	codec  runtime.ParameterCodec
}

func (f *crdclient) Create(obj *FluentBit) (*FluentBit, error) {
	var result FluentBit
	err := f.cl.Post().
		Namespace(f.ns).Resource(f.plural).
		Body(obj).Do().Into(&result)
	return &result, err
}

func (f *crdclient) Update(name string, obj *FluentBit) (*FluentBit, error) {
	var result FluentBit
	err := f.cl.Put().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Body(obj).Do().Into(&result)
	return &result, err
}

func (f *crdclient) Delete(name string, options *metav1.DeleteOptions) error {
	return f.cl.Delete().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Body(options).Do().
		Error()
}

func (f *crdclient) Get(name string) (*FluentBit, error) {
	var result FluentBit
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Do().Into(&result)
	return &result, err
}

func (f *crdclient) List(opts metav1.ListOptions) (*FluentBitList, error) {
	var result FluentBitList
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		VersionedParams(&opts, f.codec).
		Do().Into(&result)
	return &result, err
}

// Create a new List watch for our TPR
func (f *crdclient) NewListWatch() *cache.ListWatch {
	return cache.NewListWatchFromClient(f.cl, f.plural, f.ns, fields.Everything())
}

// return rest config, if path not specified assume in cluster config
func GetClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

package awsekslegacy

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsekslegacy/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Represents a resource within the Kubernetes system.
//
// Alternatively, you can use `cluster.addResource(resource[, resource, ...])`
// to define resources on this cluster.
//
// Applies/deletes the resources using `kubectl` in sync with the resource.
//
// Example:
//   var cluster cluster
//   appLabel := map[string]*string{
//   	"app": jsii.String("hello-kubernetes"),
//   }
//
//   deployment := map[string]interface{}{
//   	"apiVersion": jsii.String("apps/v1"),
//   	"kind": jsii.String("Deployment"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"replicas": jsii.Number(3),
//   		"selector": map[string]map[string]*string{
//   			"matchLabels": appLabel,
//   		},
//   		"template": map[string]map[string]map[string]*string{
//   			"metadata": map[string]map[string]*string{
//   				"labels": appLabel,
//   			},
//   			"spec": map[string][]map[string]interface{}{
//   				"containers": []map[string]interface{}{
//   					map[string]interface{}{
//   						"name": jsii.String("hello-kubernetes"),
//   						"image": jsii.String("paulbouwer/hello-kubernetes:1.5"),
//   						"ports": []map[string]*f64{
//   							map[string]*f64{
//   								"containerPort": jsii.Number(8080),
//   							},
//   						},
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
//   service := map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Service"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("hello-kubernetes"),
//   	},
//   	"spec": map[string]interface{}{
//   		"type": jsii.String("LoadBalancer"),
//   		"ports": []map[string]*f64{
//   			map[string]*f64{
//   				"port": jsii.Number(80),
//   				"targetPort": jsii.Number(8080),
//   			},
//   		},
//   		"selector": appLabel,
//   	},
//   }
//   // option 1: use a construct
//   // option 1: use a construct
//   eks.NewKubernetesResource(this, jsii.String("hello-kub"), &KubernetesResourceProps{
//   	Cluster: Cluster,
//   	Manifest: []interface{}{
//   		deployment,
//   		service,
//   	},
//   })
//
//   // or, option2: use `addResource`
//   cluster.AddResource(jsii.String("hello-kub"), service, deployment)
//
// Experimental.
type KubernetesResource interface {
	awscdk.Construct
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	OnPrepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	OnSynthesize(session constructs.ISynthesisSession)
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	OnValidate() *[]*string
	// Perform final modifications before synthesis.
	//
	// This method can be implemented by derived constructs in order to perform
	// final changes before synthesis. prepare() will be called after child
	// constructs have been prepared.
	//
	// This is an advanced framework feature. Only use this if you
	// understand the implications.
	// Experimental.
	Prepare()
	// Allows this construct to emit artifacts into the cloud assembly during synthesis.
	//
	// This method is usually implemented by framework-level constructs such as `Stack` and `Asset`
	// as they participate in synthesizing the cloud assembly.
	// Experimental.
	Synthesize(session awscdk.ISynthesisSession)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Validate the current construct.
	//
	// This method can be implemented by derived constructs in order to perform
	// validation logic. It is called on all constructs before synthesis.
	//
	// Returns: An array of validation error messages, or an empty array if the construct is valid.
	// Experimental.
	Validate() *[]*string
}

// The jsii proxy struct for KubernetesResource
type jsiiProxy_KubernetesResource struct {
	internal.Type__awscdkConstruct
}

func (j *jsiiProxy_KubernetesResource) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesResource(scope awscdk.Construct, id *string, props *KubernetesResourceProps) KubernetesResource {
	_init_.Initialize()

	if err := validateNewKubernetesResourceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesResource{}

	_jsii_.Create(
		"monocdk.aws_eks_legacy.KubernetesResource",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesResource_Override(k KubernetesResource, scope awscdk.Construct, id *string, props *KubernetesResourceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_eks_legacy.KubernetesResource",
		[]interface{}{scope, id, props},
		k,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func KubernetesResource_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesResource_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.aws_eks_legacy.KubernetesResource",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesResource_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"monocdk.aws_eks_legacy.KubernetesResource",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesResource) OnPrepare() {
	_jsii_.InvokeVoid(
		k,
		"onPrepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesResource) OnSynthesize(session constructs.ISynthesisSession) {
	if err := k.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (k *jsiiProxy_KubernetesResource) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesResource) Prepare() {
	_jsii_.InvokeVoid(
		k,
		"prepare",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesResource) Synthesize(session awscdk.ISynthesisSession) {
	if err := k.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"synthesize",
		[]interface{}{session},
	)
}

func (k *jsiiProxy_KubernetesResource) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesResource) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


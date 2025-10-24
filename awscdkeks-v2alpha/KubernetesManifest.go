package awscdkeks-v2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a manifest within the Kubernetes system.
//
// Alternatively, you can use `cluster.addManifest(resource[, resource, ...])`
// to define resources on this cluster.
//
// Applies/deletes the manifest using `kubectl`.
//
// Example:
//   var cluster Cluster
//
//   namespace := cluster.addManifest(jsii.String("my-namespace"), map[string]interface{}{
//   	"apiVersion": jsii.String("v1"),
//   	"kind": jsii.String("Namespace"),
//   	"metadata": map[string]*string{
//   		"name": jsii.String("my-app"),
//   	},
//   })
//
//   service := cluster.addManifest(jsii.String("my-service"), map[string]interface{}{
//   	"metadata": map[string]*string{
//   		"name": jsii.String("myservice"),
//   		"namespace": jsii.String("my-app"),
//   	},
//   	"spec": map[string]interface{}{
//   	},
//   })
//
//   service.Node.AddDependency(namespace)
//
// Experimental.
type KubernetesManifest interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesManifest
type jsiiProxy_KubernetesManifest struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KubernetesManifest) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesManifest(scope constructs.Construct, id *string, props *KubernetesManifestProps) KubernetesManifest {
	_init_.Initialize()

	if err := validateNewKubernetesManifestParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesManifest{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifest",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesManifest_Override(k KubernetesManifest, scope constructs.Construct, id *string, props *KubernetesManifestProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifest",
		[]interface{}{scope, id, props},
		k,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func KubernetesManifest_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesManifest_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifest",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesManifest_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesManifest",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesManifest) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


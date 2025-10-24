package awscdkeks-v2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkeks-v2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Construct for installing the AWS ALB Contoller on EKS clusters.
//
// Use the factory functions `get` and `getOrCreate` to obtain/create instances of this controller.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import eks_v2_alpha "github.com/aws/aws-cdk-go/awscdkeks-v2alpha"
//
//   var albControllerVersion AlbControllerVersion
//   var cluster Cluster
//   var policy interface{}
//
//   albController := eks_v2_alpha.NewAlbController(this, jsii.String("MyAlbController"), &AlbControllerProps{
//   	Cluster: cluster,
//   	Version: albControllerVersion,
//
//   	// the properties below are optional
//   	Policy: policy,
//   	Repository: jsii.String("repository"),
//   })
//
// See: https://kubernetes-sigs.github.io/aws-load-balancer-controller
//
// Experimental.
type AlbController interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AlbController
type jsiiProxy_AlbController struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_AlbController) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlbController(scope constructs.Construct, id *string, props *AlbControllerProps) AlbController {
	_init_.Initialize()

	if err := validateNewAlbControllerParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlbController{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AlbController",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlbController_Override(a AlbController, scope constructs.Construct, id *string, props *AlbControllerProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.AlbController",
		[]interface{}{scope, id, props},
		a,
	)
}

// Create the controller construct associated with this cluster and scope.
//
// Singleton per stack/cluster.
// Experimental.
func AlbController_Create(scope constructs.Construct, props *AlbControllerProps) AlbController {
	_init_.Initialize()

	if err := validateAlbController_CreateParameters(scope, props); err != nil {
		panic(err)
	}
	var returns AlbController

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.AlbController",
		"create",
		[]interface{}{scope, props},
		&returns,
	)

	return returns
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
func AlbController_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAlbController_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.AlbController",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlbController) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


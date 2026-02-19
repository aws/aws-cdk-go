package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a value of a specific object deployed in the cluster.
//
// Use this to fetch any information available by the `kubectl get` command.
//
// Example:
//   var cluster Cluster
//
//   // query the load balancer address
//   myServiceAddress := eks.NewKubernetesObjectValue(this, jsii.String("LoadBalancerAttribute"), &KubernetesObjectValueProps{
//   	Cluster: cluster,
//   	ObjectType: jsii.String("service"),
//   	ObjectName: jsii.String("my-service"),
//   	JsonPath: jsii.String(".status.loadBalancer.ingress[0].hostname"),
//   })
//
//   // pass the address to a lambda function
//   proxyFunction := lambda.NewFunction(this, jsii.String("ProxyFunction"), &FunctionProps{
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromInline(jsii.String("my-code")),
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Environment: map[string]*string{
//   		"myServiceAddress": myServiceAddress.value,
//   	},
//   })
//
// Experimental.
type KubernetesObjectValue interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The value as a string token.
	// Experimental.
	Value() *string
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for KubernetesObjectValue
type jsiiProxy_KubernetesObjectValue struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_KubernetesObjectValue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesObjectValue) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubernetesObjectValue(scope constructs.Construct, id *string, props *KubernetesObjectValueProps) KubernetesObjectValue {
	_init_.Initialize()

	if err := validateNewKubernetesObjectValueParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesObjectValue{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubernetesObjectValue_Override(k KubernetesObjectValue, scope constructs.Construct, id *string, props *KubernetesObjectValueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValue",
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
func KubernetesObjectValue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesObjectValue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesObjectValue_RESOURCE_TYPE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@aws-cdk/aws-eks-v2-alpha.KubernetesObjectValue",
		"RESOURCE_TYPE",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesObjectValue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesObjectValue) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		k,
		"with",
		args,
		&returns,
	)

	return returns
}


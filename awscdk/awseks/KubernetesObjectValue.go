package awseks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awseks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a value of a specific object deployed in the cluster.
//
// Use this to fetch any information available by the `kubectl get` command.
//
// Example:
//   var cluster cluster
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
type KubernetesObjectValue interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// The value as a string token.
	Value() *string
	// Returns a string representation of this construct.
	ToString() *string
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


func NewKubernetesObjectValue(scope constructs.Construct, id *string, props *KubernetesObjectValueProps) KubernetesObjectValue {
	_init_.Initialize()

	if err := validateNewKubernetesObjectValueParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesObjectValue{}

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.KubernetesObjectValue",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewKubernetesObjectValue_Override(k KubernetesObjectValue, scope constructs.Construct, id *string, props *KubernetesObjectValueProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_eks.KubernetesObjectValue",
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
func KubernetesObjectValue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesObjectValue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_eks.KubernetesObjectValue",
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
		"aws-cdk-lib.aws_eks.KubernetesObjectValue",
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


package awscdkeksv2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdkeksv2alpha/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Implementation of Kubectl Lambda.
//
// Example:
//   handlerRole := iam.Role_FromRoleArn(this, jsii.String("HandlerRole"), jsii.String("arn:aws:iam::123456789012:role/lambda-role"))
//   // get the serivceToken from the custom resource provider
//   functionArn := lambda.Function_FromFunctionName(this, jsii.String("ProviderOnEventFunc"), jsii.String("ProviderframeworkonEvent-XXX")).FunctionArn
//   kubectlProvider := eks.KubectlProvider_FromKubectlProviderAttributes(this, jsii.String("KubectlProvider"), &KubectlProviderAttributes{
//   	ServiceToken: functionArn,
//   	Role: handlerRole,
//   })
//
//   cluster := eks.Cluster_FromClusterAttributes(this, jsii.String("Cluster"), &ClusterAttributes{
//   	ClusterName: jsii.String("cluster"),
//   	KubectlProvider: KubectlProvider,
//   })
//
// Experimental.
type KubectlProvider interface {
	constructs.Construct
	IKubectlProvider
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// The IAM execution role of the handler.
	// Experimental.
	Role() awsiam.IRole
	// The custom resource provider's service token.
	// Experimental.
	ServiceToken() *string
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

// The jsii proxy struct for KubectlProvider
type jsiiProxy_KubectlProvider struct {
	internal.Type__constructsConstruct
	jsiiProxy_IKubectlProvider
}

func (j *jsiiProxy_KubectlProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubectlProvider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewKubectlProvider(scope constructs.Construct, id *string, props *KubectlProviderProps) KubectlProvider {
	_init_.Initialize()

	if err := validateNewKubectlProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubectlProvider{}

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKubectlProvider_Override(k KubectlProvider, scope constructs.Construct, id *string, props *KubectlProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		[]interface{}{scope, id, props},
		k,
	)
}

// Import an existing provider.
// Experimental.
func KubectlProvider_FromKubectlProviderAttributes(scope constructs.Construct, id *string, attrs *KubectlProviderAttributes) IKubectlProvider {
	_init_.Initialize()

	if err := validateKubectlProvider_FromKubectlProviderAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IKubectlProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		"fromKubectlProviderAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Take existing provider on cluster.
// Experimental.
func KubectlProvider_GetKubectlProvider(scope constructs.Construct, cluster ICluster) IKubectlProvider {
	_init_.Initialize()

	if err := validateKubectlProvider_GetKubectlProviderParameters(scope, cluster); err != nil {
		panic(err)
	}
	var returns IKubectlProvider

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		"getKubectlProvider",
		[]interface{}{scope, cluster},
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
func KubectlProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubectlProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-eks-v2-alpha.KubectlProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubectlProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubectlProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
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


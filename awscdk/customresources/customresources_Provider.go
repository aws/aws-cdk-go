package customresources

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awscloudformation"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/customresources/internal"
	"github.com/aws/constructs-go/constructs/v3"
)

// Defines an AWS CloudFormation custom resource provider.
//
// Example:
//   var onEvent function
//   var isComplete function
//   var myRole role
//
//   myProvider := cr.NewProvider(this, jsii.String("MyProvider"), &ProviderProps{
//   	OnEventHandler: onEvent,
//   	IsCompleteHandler: isComplete,
//   	LogRetention: logs.RetentionDays_ONE_DAY,
//   	Role: myRole,
//   	ProviderFunctionName: jsii.String("the-lambda-name"),
//   })
//
// Experimental.
type Provider interface {
	awscdk.Construct
	awscloudformation.ICustomResourceProvider
	// The user-defined AWS Lambda function which is invoked asynchronously in order to determine if the operation is complete.
	// Experimental.
	IsCompleteHandler() awslambda.IFunction
	// The construct tree node associated with this construct.
	// Experimental.
	Node() awscdk.ConstructNode
	// The user-defined AWS Lambda function which is invoked for all resource lifecycle operations (CREATE/UPDATE/DELETE).
	// Experimental.
	OnEventHandler() awslambda.IFunction
	// The service token to use in order to define custom resources that are backed by this provider.
	// Experimental.
	ServiceToken() *string
	// Called by `CustomResource` which uses this provider.
	// Deprecated: use `provider.serviceToken` instead
	Bind(_scope awscdk.Construct) *awscloudformation.CustomResourceProviderConfig
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

// The jsii proxy struct for Provider
type jsiiProxy_Provider struct {
	internal.Type__awscdkConstruct
	internal.Type__awscloudformationICustomResourceProvider
}

func (j *jsiiProxy_Provider) IsCompleteHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"isCompleteHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) Node() awscdk.ConstructNode {
	var returns awscdk.ConstructNode
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) OnEventHandler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"onEventHandler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Provider) ServiceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceToken",
		&returns,
	)
	return returns
}


// Experimental.
func NewProvider(scope constructs.Construct, id *string, props *ProviderProps) Provider {
	_init_.Initialize()

	if err := validateNewProviderParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Provider{}

	_jsii_.Create(
		"monocdk.custom_resources.Provider",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewProvider_Override(p Provider, scope constructs.Construct, id *string, props *ProviderProps) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.custom_resources.Provider",
		[]interface{}{scope, id, props},
		p,
	)
}

// Return whether the given object is a Construct.
// Experimental.
func Provider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"monocdk.custom_resources.Provider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Bind(_scope awscdk.Construct) *awscloudformation.CustomResourceProviderConfig {
	if err := p.validateBindParameters(_scope); err != nil {
		panic(err)
	}
	var returns *awscloudformation.CustomResourceProviderConfig

	_jsii_.Invoke(
		p,
		"bind",
		[]interface{}{_scope},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) OnPrepare() {
	_jsii_.InvokeVoid(
		p,
		"onPrepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Provider) OnSynthesize(session constructs.ISynthesisSession) {
	if err := p.validateOnSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"onSynthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Provider) OnValidate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"onValidate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Prepare() {
	_jsii_.InvokeVoid(
		p,
		"prepare",
		nil, // no parameters
	)
}

func (p *jsiiProxy_Provider) Synthesize(session awscdk.ISynthesisSession) {
	if err := p.validateSynthesizeParameters(session); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"synthesize",
		[]interface{}{session},
	)
}

func (p *jsiiProxy_Provider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_Provider) Validate() *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"validate",
		nil, // no parameters
		&returns,
	)

	return returns
}


package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

// A Lambda function Url.
type IFunctionUrl interface {
	awscdk.IResource
	interfacesawslambda.IUrlRef
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant
	// The authType of the function URL, used for access control.
	AuthType() FunctionUrlAuthType
	// The ARN of the function this URL refers to.
	FunctionArn() *string
	// The url of the Lambda function.
	Url() *string
}

// The jsii proxy for IFunctionUrl
type jsiiProxy_IFunctionUrl struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawslambdaIUrlRef
}

func (i *jsiiProxy_IFunctionUrl) GrantInvokeUrl(identity awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeUrlParameters(identity); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeUrl",
		[]interface{}{identity},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IFunctionUrl) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IFunctionUrl) AuthType() FunctionUrlAuthType {
	var returns FunctionUrlAuthType
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionUrl) UrlRef() *interfacesawslambda.UrlReference {
	var returns *interfacesawslambda.UrlReference
	_jsii_.Get(
		j,
		"urlRef",
		&returns,
	)
	return returns
}


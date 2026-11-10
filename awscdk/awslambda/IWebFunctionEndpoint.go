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

// Represents a WebFunction endpoint, either created in this stack or imported.
type IWebFunctionEndpoint interface {
	awscdk.IResource
	awsiam.IResourceWithPolicyV2
	interfacesawslambda.IWebFunctionEndpointRef
	// How the endpoint authenticates requests.
	AuthType() WebFunctionAuthType
	// The domain name of the endpoint: the host clients send requests to.
	DomainName() *string
	// The ARN of the endpoint.
	EndpointArn() *string
	// The endpoint type.
	EndpointType() WebFunctionEndpointType
	// Grant methods for the endpoint.
	Grants() WebFunctionEndpointGrants
	// The ARN of the WebFunction the endpoint serves.
	WebFunctionArn() *string
}

// The jsii proxy for IWebFunctionEndpoint
type jsiiProxy_IWebFunctionEndpoint struct {
	internal.Type__awscdkIResource
	internal.Type__awsiamIResourceWithPolicyV2
	internal.Type__interfacesawslambdaIWebFunctionEndpointRef
}

func (i *jsiiProxy_IWebFunctionEndpoint) AddToResourcePolicy(statement awsiam.PolicyStatement) *awsiam.AddToResourcePolicyResult {
	if err := i.validateAddToResourcePolicyParameters(statement); err != nil {
		panic(err)
	}
	var returns *awsiam.AddToResourcePolicyResult

	_jsii_.Invoke(
		i,
		"addToResourcePolicy",
		[]interface{}{statement},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IWebFunctionEndpoint) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IWebFunctionEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) AuthType() WebFunctionAuthType {
	var returns WebFunctionAuthType
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) EndpointArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) EndpointType() WebFunctionEndpointType {
	var returns WebFunctionEndpointType
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) Grants() WebFunctionEndpointGrants {
	var returns WebFunctionEndpointGrants
	_jsii_.Get(
		j,
		"grants",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) WebFunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"webFunctionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWebFunctionEndpoint) WebFunctionEndpointRef() *interfacesawslambda.WebFunctionEndpointReference {
	var returns *interfacesawslambda.WebFunctionEndpointReference
	_jsii_.Get(
		j,
		"webFunctionEndpointRef",
		&returns,
	)
	return returns
}


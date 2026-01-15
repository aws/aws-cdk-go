package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents an ApiGatewayV2 ApiMapping resource.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigatewayv2-apimapping.html
//
type IApiMapping interface {
	interfacesawsapigatewayv2.IApiMappingRef
	awscdk.IResource
	// ID of the api mapping.
	ApiMappingId() *string
}

// The jsii proxy for IApiMapping
type jsiiProxy_IApiMapping struct {
	internal.Type__interfacesawsapigatewayv2IApiMappingRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IApiMapping) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (j *jsiiProxy_IApiMapping) ApiMappingId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiMappingId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMapping) ApiMappingRef() *interfacesawsapigatewayv2.ApiMappingReference {
	var returns *interfacesawsapigatewayv2.ApiMappingReference
	_jsii_.Get(
		j,
		"apiMappingRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMapping) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiMapping) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


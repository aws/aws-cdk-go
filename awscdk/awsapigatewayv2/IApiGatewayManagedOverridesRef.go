package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiGatewayManagedOverrides.
// Experimental.
type IApiGatewayManagedOverridesRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ApiGatewayManagedOverrides resource.
	// Experimental.
	ApiGatewayManagedOverridesRef() *ApiGatewayManagedOverridesReference
}

// The jsii proxy for IApiGatewayManagedOverridesRef
type jsiiProxy_IApiGatewayManagedOverridesRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) ApiGatewayManagedOverridesRef() *ApiGatewayManagedOverridesReference {
	var returns *ApiGatewayManagedOverridesReference
	_jsii_.Get(
		j,
		"apiGatewayManagedOverridesRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiGatewayManagedOverridesRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


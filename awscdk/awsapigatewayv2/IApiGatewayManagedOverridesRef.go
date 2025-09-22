package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiGatewayManagedOverrides.
// Experimental.
type IApiGatewayManagedOverridesRef interface {
	constructs.IConstruct
	// A reference to a ApiGatewayManagedOverrides resource.
	// Experimental.
	ApiGatewayManagedOverridesRef() *ApiGatewayManagedOverridesReference
}

// The jsii proxy for IApiGatewayManagedOverridesRef
type jsiiProxy_IApiGatewayManagedOverridesRef struct {
	internal.Type__constructsIConstruct
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


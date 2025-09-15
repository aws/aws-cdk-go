package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IntegrationResponse.
// Experimental.
type IIntegrationResponseRef interface {
	constructs.IConstruct
	// A reference to a IntegrationResponse resource.
	// Experimental.
	IntegrationResponseRef() *IntegrationResponseReference
}

// The jsii proxy for IIntegrationResponseRef
type jsiiProxy_IIntegrationResponseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIntegrationResponseRef) IntegrationResponseRef() *IntegrationResponseReference {
	var returns *IntegrationResponseReference
	_jsii_.Get(
		j,
		"integrationResponseRef",
		&returns,
	)
	return returns
}


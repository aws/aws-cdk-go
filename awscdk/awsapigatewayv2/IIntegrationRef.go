package awsapigatewayv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Integration.
// Experimental.
type IIntegrationRef interface {
	constructs.IConstruct
	// A reference to a Integration resource.
	// Experimental.
	IntegrationRef() *IntegrationReference
}

// The jsii proxy for IIntegrationRef
type jsiiProxy_IIntegrationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIntegrationRef) IntegrationRef() *IntegrationReference {
	var returns *IntegrationReference
	_jsii_.Get(
		j,
		"integrationRef",
		&returns,
	)
	return returns
}


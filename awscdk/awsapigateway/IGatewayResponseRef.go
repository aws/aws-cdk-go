package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayResponse.
// Experimental.
type IGatewayResponseRef interface {
	constructs.IConstruct
	// A reference to a GatewayResponse resource.
	// Experimental.
	GatewayResponseRef() *GatewayResponseReference
}

// The jsii proxy for IGatewayResponseRef
type jsiiProxy_IGatewayResponseRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGatewayResponseRef) GatewayResponseRef() *GatewayResponseReference {
	var returns *GatewayResponseReference
	_jsii_.Get(
		j,
		"gatewayResponseRef",
		&returns,
	)
	return returns
}


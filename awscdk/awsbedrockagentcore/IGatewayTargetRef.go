package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrockagentcore/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GatewayTarget.
// Experimental.
type IGatewayTargetRef interface {
	constructs.IConstruct
	// A reference to a GatewayTarget resource.
	// Experimental.
	GatewayTargetRef() *GatewayTargetReference
}

// The jsii proxy for IGatewayTargetRef
type jsiiProxy_IGatewayTargetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGatewayTargetRef) GatewayTargetRef() *GatewayTargetReference {
	var returns *GatewayTargetReference
	_jsii_.Get(
		j,
		"gatewayTargetRef",
		&returns,
	)
	return returns
}


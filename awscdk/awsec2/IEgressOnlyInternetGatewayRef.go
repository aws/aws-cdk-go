package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EgressOnlyInternetGateway.
// Experimental.
type IEgressOnlyInternetGatewayRef interface {
	constructs.IConstruct
	// A reference to a EgressOnlyInternetGateway resource.
	// Experimental.
	EgressOnlyInternetGatewayRef() *EgressOnlyInternetGatewayReference
}

// The jsii proxy for IEgressOnlyInternetGatewayRef
type jsiiProxy_IEgressOnlyInternetGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEgressOnlyInternetGatewayRef) EgressOnlyInternetGatewayRef() *EgressOnlyInternetGatewayReference {
	var returns *EgressOnlyInternetGatewayReference
	_jsii_.Get(
		j,
		"egressOnlyInternetGatewayRef",
		&returns,
	)
	return returns
}


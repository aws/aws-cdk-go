package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CustomerGateway.
// Experimental.
type ICustomerGatewayRef interface {
	constructs.IConstruct
	// A reference to a CustomerGateway resource.
	// Experimental.
	CustomerGatewayRef() *CustomerGatewayReference
}

// The jsii proxy for ICustomerGatewayRef
type jsiiProxy_ICustomerGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICustomerGatewayRef) CustomerGatewayRef() *CustomerGatewayReference {
	var returns *CustomerGatewayReference
	_jsii_.Get(
		j,
		"customerGatewayRef",
		&returns,
	)
	return returns
}


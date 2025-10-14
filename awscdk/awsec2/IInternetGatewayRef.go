package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a InternetGateway.
// Experimental.
type IInternetGatewayRef interface {
	constructs.IConstruct
	// A reference to a InternetGateway resource.
	// Experimental.
	InternetGatewayRef() *InternetGatewayReference
}

// The jsii proxy for IInternetGatewayRef
type jsiiProxy_IInternetGatewayRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IInternetGatewayRef) InternetGatewayRef() *InternetGatewayReference {
	var returns *InternetGatewayReference
	_jsii_.Get(
		j,
		"internetGatewayRef",
		&returns,
	)
	return returns
}


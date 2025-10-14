package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Subnet.
// Experimental.
type ISubnetRef interface {
	constructs.IConstruct
	// A reference to a Subnet resource.
	// Experimental.
	SubnetRef() *SubnetReference
}

// The jsii proxy for ISubnetRef
type jsiiProxy_ISubnetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISubnetRef) SubnetRef() *SubnetReference {
	var returns *SubnetReference
	_jsii_.Get(
		j,
		"subnetRef",
		&returns,
	)
	return returns
}


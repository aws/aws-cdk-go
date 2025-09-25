package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapprunner/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcConnector.
// Experimental.
type IVpcConnectorRef interface {
	constructs.IConstruct
	// A reference to a VpcConnector resource.
	// Experimental.
	VpcConnectorRef() *VpcConnectorReference
}

// The jsii proxy for IVpcConnectorRef
type jsiiProxy_IVpcConnectorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcConnectorRef) VpcConnectorRef() *VpcConnectorReference {
	var returns *VpcConnectorReference
	_jsii_.Get(
		j,
		"vpcConnectorRef",
		&returns,
	)
	return returns
}


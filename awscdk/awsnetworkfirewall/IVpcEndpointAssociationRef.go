package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpointAssociation.
// Experimental.
type IVpcEndpointAssociationRef interface {
	constructs.IConstruct
	// A reference to a VpcEndpointAssociation resource.
	// Experimental.
	VpcEndpointAssociationRef() *VpcEndpointAssociationReference
}

// The jsii proxy for IVpcEndpointAssociationRef
type jsiiProxy_IVpcEndpointAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVpcEndpointAssociationRef) VpcEndpointAssociationRef() *VpcEndpointAssociationReference {
	var returns *VpcEndpointAssociationReference
	_jsii_.Get(
		j,
		"vpcEndpointAssociationRef",
		&returns,
	)
	return returns
}


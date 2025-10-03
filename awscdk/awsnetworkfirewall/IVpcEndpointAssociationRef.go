package awsnetworkfirewall

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnetworkfirewall/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VpcEndpointAssociation.
// Experimental.
type IVpcEndpointAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVpcEndpointAssociationRef
type jsiiProxy_IVpcEndpointAssociationRef struct {
	internal.Type__constructsIConstruct
}


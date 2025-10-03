package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCDHCPOptionsAssociation.
// Experimental.
type IVPCDHCPOptionsAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVPCDHCPOptionsAssociationRef
type jsiiProxy_IVPCDHCPOptionsAssociationRef struct {
	internal.Type__constructsIConstruct
}


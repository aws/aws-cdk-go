package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCDHCPOptionsAssociation.
// Experimental.
type IVPCDHCPOptionsAssociationRef interface {
	constructs.IConstruct
	// A reference to a VPCDHCPOptionsAssociation resource.
	// Experimental.
	VpcdhcpOptionsAssociationRef() *VPCDHCPOptionsAssociationReference
}

// The jsii proxy for IVPCDHCPOptionsAssociationRef
type jsiiProxy_IVPCDHCPOptionsAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVPCDHCPOptionsAssociationRef) VpcdhcpOptionsAssociationRef() *VPCDHCPOptionsAssociationReference {
	var returns *VPCDHCPOptionsAssociationReference
	_jsii_.Get(
		j,
		"vpcdhcpOptionsAssociationRef",
		&returns,
	)
	return returns
}


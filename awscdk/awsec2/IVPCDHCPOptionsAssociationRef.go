package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a VPCDHCPOptionsAssociation.
// Experimental.
type IVPCDHCPOptionsAssociationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a VPCDHCPOptionsAssociation resource.
	// Experimental.
	VpcdhcpOptionsAssociationRef() *VPCDHCPOptionsAssociationReference
}

// The jsii proxy for IVPCDHCPOptionsAssociationRef
type jsiiProxy_IVPCDHCPOptionsAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IVPCDHCPOptionsAssociationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVPCDHCPOptionsAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


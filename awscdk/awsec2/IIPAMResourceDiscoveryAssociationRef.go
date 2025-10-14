package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMResourceDiscoveryAssociation.
// Experimental.
type IIPAMResourceDiscoveryAssociationRef interface {
	constructs.IConstruct
	// A reference to a IPAMResourceDiscoveryAssociation resource.
	// Experimental.
	IpamResourceDiscoveryAssociationRef() *IPAMResourceDiscoveryAssociationReference
}

// The jsii proxy for IIPAMResourceDiscoveryAssociationRef
type jsiiProxy_IIPAMResourceDiscoveryAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMResourceDiscoveryAssociationRef) IpamResourceDiscoveryAssociationRef() *IPAMResourceDiscoveryAssociationReference {
	var returns *IPAMResourceDiscoveryAssociationReference
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryAssociationRef",
		&returns,
	)
	return returns
}


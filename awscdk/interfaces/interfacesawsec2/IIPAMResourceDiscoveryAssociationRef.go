package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMResourceDiscoveryAssociation.
// Experimental.
type IIPAMResourceDiscoveryAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a IPAMResourceDiscoveryAssociation resource.
	// Experimental.
	IpamResourceDiscoveryAssociationRef() *IPAMResourceDiscoveryAssociationReference
}

// The jsii proxy for IIPAMResourceDiscoveryAssociationRef
type jsiiProxy_IIPAMResourceDiscoveryAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IIPAMResourceDiscoveryAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IIPAMResourceDiscoveryAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


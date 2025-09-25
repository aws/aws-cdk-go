package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a IPAMResourceDiscovery.
// Experimental.
type IIPAMResourceDiscoveryRef interface {
	constructs.IConstruct
	// A reference to a IPAMResourceDiscovery resource.
	// Experimental.
	IpamResourceDiscoveryRef() *IPAMResourceDiscoveryReference
}

// The jsii proxy for IIPAMResourceDiscoveryRef
type jsiiProxy_IIPAMResourceDiscoveryRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IIPAMResourceDiscoveryRef) IpamResourceDiscoveryRef() *IPAMResourceDiscoveryReference {
	var returns *IPAMResourceDiscoveryReference
	_jsii_.Get(
		j,
		"ipamResourceDiscoveryRef",
		&returns,
	)
	return returns
}


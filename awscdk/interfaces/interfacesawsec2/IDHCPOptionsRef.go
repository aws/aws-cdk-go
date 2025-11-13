package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DHCPOptions.
// Experimental.
type IDHCPOptionsRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DHCPOptions resource.
	// Experimental.
	DhcpOptionsRef() *DHCPOptionsReference
}

// The jsii proxy for IDHCPOptionsRef
type jsiiProxy_IDHCPOptionsRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IDHCPOptionsRef) DhcpOptionsRef() *DHCPOptionsReference {
	var returns *DHCPOptionsReference
	_jsii_.Get(
		j,
		"dhcpOptionsRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDHCPOptionsRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDHCPOptionsRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


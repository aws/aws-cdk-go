package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DHCPOptions.
// Experimental.
type IDHCPOptionsRef interface {
	constructs.IConstruct
	// A reference to a DHCPOptions resource.
	// Experimental.
	DhcpOptionsRef() *DHCPOptionsReference
}

// The jsii proxy for IDHCPOptionsRef
type jsiiProxy_IDHCPOptionsRef struct {
	internal.Type__constructsIConstruct
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


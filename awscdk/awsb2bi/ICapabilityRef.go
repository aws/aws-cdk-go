package awsb2bi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Capability.
// Experimental.
type ICapabilityRef interface {
	constructs.IConstruct
	// A reference to a Capability resource.
	// Experimental.
	CapabilityRef() *CapabilityReference
}

// The jsii proxy for ICapabilityRef
type jsiiProxy_ICapabilityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICapabilityRef) CapabilityRef() *CapabilityReference {
	var returns *CapabilityReference
	_jsii_.Get(
		j,
		"capabilityRef",
		&returns,
	)
	return returns
}


package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MulticastGroup.
// Experimental.
type IMulticastGroupRef interface {
	constructs.IConstruct
	// A reference to a MulticastGroup resource.
	// Experimental.
	MulticastGroupRef() *MulticastGroupReference
}

// The jsii proxy for IMulticastGroupRef
type jsiiProxy_IMulticastGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMulticastGroupRef) MulticastGroupRef() *MulticastGroupReference {
	var returns *MulticastGroupReference
	_jsii_.Get(
		j,
		"multicastGroupRef",
		&returns,
	)
	return returns
}


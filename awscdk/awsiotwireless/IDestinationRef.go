package awsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Destination.
// Experimental.
type IDestinationRef interface {
	constructs.IConstruct
	// A reference to a Destination resource.
	// Experimental.
	DestinationRef() *DestinationReference
}

// The jsii proxy for IDestinationRef
type jsiiProxy_IDestinationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IDestinationRef) DestinationRef() *DestinationReference {
	var returns *DestinationReference
	_jsii_.Get(
		j,
		"destinationRef",
		&returns,
	)
	return returns
}


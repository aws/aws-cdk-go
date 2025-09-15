package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ApiDestination.
// Experimental.
type IApiDestinationRef interface {
	constructs.IConstruct
	// A reference to a ApiDestination resource.
	// Experimental.
	ApiDestinationRef() *ApiDestinationReference
}

// The jsii proxy for IApiDestinationRef
type jsiiProxy_IApiDestinationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IApiDestinationRef) ApiDestinationRef() *ApiDestinationReference {
	var returns *ApiDestinationReference
	_jsii_.Get(
		j,
		"apiDestinationRef",
		&returns,
	)
	return returns
}


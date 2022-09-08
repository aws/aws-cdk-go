package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsevents/internal"
)

// Interface for API Destinations.
// Experimental.
type IApiDestination interface {
	awscdk.IResource
	// The ARN of the Api Destination created.
	// Experimental.
	ApiDestinationArn() *string
	// The Name of the Api Destination created.
	// Experimental.
	ApiDestinationName() *string
}

// The jsii proxy for IApiDestination
type jsiiProxy_IApiDestination struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IApiDestination) ApiDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiDestination) ApiDestinationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDestinationName",
		&returns,
	)
	return returns
}


package awsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents/internal"
)

// Interface for API Destinations.
type IApiDestination interface {
	awscdk.IResource
	// The ARN of the Api Destination created.
	ApiDestinationArn() *string
	// The Name of the Api Destination created.
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


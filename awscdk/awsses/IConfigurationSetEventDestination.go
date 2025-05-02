package awsses

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsses/internal"
)

// A configuration set event destination.
type IConfigurationSetEventDestination interface {
	awscdk.IResource
	// The ID of the configuration set event destination.
	ConfigurationSetEventDestinationId() *string
}

// The jsii proxy for IConfigurationSetEventDestination
type jsiiProxy_IConfigurationSetEventDestination struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IConfigurationSetEventDestination) ConfigurationSetEventDestinationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationSetEventDestinationId",
		&returns,
	)
	return returns
}


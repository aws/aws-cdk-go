package awscdkivsalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkivsalpha/v2/internal"
)

// Represents the IVS Recording configuration.
// Experimental.
type IRecordingConfiguration interface {
	awscdk.IResource
	// The ARN of the Recording configuration.
	// Experimental.
	RecordingConfigurationArn() *string
	// The ID of the Recording configuration.
	// Experimental.
	RecordingConfigurationId() *string
}

// The jsii proxy for IRecordingConfiguration
type jsiiProxy_IRecordingConfiguration struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRecordingConfiguration) RecordingConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordingConfiguration) RecordingConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recordingConfigurationId",
		&returns,
	)
	return returns
}


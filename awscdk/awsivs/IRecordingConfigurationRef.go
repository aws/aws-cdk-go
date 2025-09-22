package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordingConfiguration.
// Experimental.
type IRecordingConfigurationRef interface {
	constructs.IConstruct
	// A reference to a RecordingConfiguration resource.
	// Experimental.
	RecordingConfigurationRef() *RecordingConfigurationReference
}

// The jsii proxy for IRecordingConfigurationRef
type jsiiProxy_IRecordingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRecordingConfigurationRef) RecordingConfigurationRef() *RecordingConfigurationReference {
	var returns *RecordingConfigurationReference
	_jsii_.Get(
		j,
		"recordingConfigurationRef",
		&returns,
	)
	return returns
}


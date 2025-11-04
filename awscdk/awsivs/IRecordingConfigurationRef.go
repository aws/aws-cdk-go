package awsivs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsivs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RecordingConfiguration.
// Experimental.
type IRecordingConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a RecordingConfiguration resource.
	// Experimental.
	RecordingConfigurationRef() *RecordingConfigurationReference
}

// The jsii proxy for IRecordingConfigurationRef
type jsiiProxy_IRecordingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IRecordingConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecordingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


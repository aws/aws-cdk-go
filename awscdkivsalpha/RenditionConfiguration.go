package awscdkivsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkivsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Rendition configuration for IVS Recording configuration.
//
// Example:
//   var recordingBucket bucket
//
//
//   recordingConfiguration := ivs.NewRecordingConfiguration(this, jsii.String("RecordingConfiguration"), &RecordingConfigurationProps{
//   	Bucket: recordingBucket,
//
//   	// set rendition configuration
//   	RenditionConfiguration: ivs.RenditionConfiguration_Custom([]resolution{
//   		ivs.*resolution_HD,
//   		ivs.*resolution_SD,
//   	}),
//   })
//
// Experimental.
type RenditionConfiguration interface {
	// A list of which renditions are recorded for a stream.
	//
	// If you do not specify this property, no resolution is selected.
	// Experimental.
	Renditions() *[]Resolution
	// The set of renditions are recorded for a stream.
	// Experimental.
	RenditionSelection() RenditionSelection
}

// The jsii proxy struct for RenditionConfiguration
type jsiiProxy_RenditionConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_RenditionConfiguration) Renditions() *[]Resolution {
	var returns *[]Resolution
	_jsii_.Get(
		j,
		"renditions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RenditionConfiguration) RenditionSelection() RenditionSelection {
	var returns RenditionSelection
	_jsii_.Get(
		j,
		"renditionSelection",
		&returns,
	)
	return returns
}


// Record all available renditions.
// Experimental.
func RenditionConfiguration_All() RenditionConfiguration {
	_init_.Initialize()

	var returns RenditionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.RenditionConfiguration",
		"all",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Record a subset of video renditions.
// Experimental.
func RenditionConfiguration_Custom(renditions *[]Resolution) RenditionConfiguration {
	_init_.Initialize()

	if err := validateRenditionConfiguration_CustomParameters(renditions); err != nil {
		panic(err)
	}
	var returns RenditionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.RenditionConfiguration",
		"custom",
		[]interface{}{renditions},
		&returns,
	)

	return returns
}

// Does not record any video.
// Experimental.
func RenditionConfiguration_None() RenditionConfiguration {
	_init_.Initialize()

	var returns RenditionConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.RenditionConfiguration",
		"none",
		nil, // no parameters
		&returns,
	)

	return returns
}


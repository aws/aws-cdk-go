package awscdkivsalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkivsalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Thumbnail configuration for IVS Recording configuration.
//
// Example:
//   var recordingBucket Bucket
//
//
//   recordingConfiguration := ivs.NewRecordingConfiguration(this, jsii.String("RecordingConfiguration"), &RecordingConfigurationProps{
//   	Bucket: recordingBucket,
//
//   	// set thumbnail settings
//   	ThumbnailConfiguration: ivs.ThumbnailConfiguration_Interval(ivs.Resolution_HD, []Storage{
//   		ivs.Storage_LATEST,
//   		ivs.Storage_SEQUENTIAL,
//   	}, awscdk.Duration_Seconds(jsii.Number(30))),
//   })
//
// Experimental.
type ThumbnailConfiguration interface {
	// Thumbnail recording mode.
	//
	// If you do not specify this property, `ThumbnailRecordingMode.INTERVAL` is set.
	// Experimental.
	RecordingMode() RecordingMode
	// The desired resolution of recorded thumbnails for a stream.
	//
	// If you do not specify this property, same resolution as Input stream is used.
	// Experimental.
	Resolution() Resolution
	// The format in which thumbnails are recorded for a stream.
	//
	// If you do not specify this property, `ThumbnailStorage.SEQUENTIAL` is set.
	// Experimental.
	Storage() *[]Storage
	// The targeted thumbnail-generation interval.
	//
	// Must be between 1 and 60 seconds. If you do not specify this property, `Duration.seconds(60)` is set.
	// Experimental.
	TargetInterval() awscdk.Duration
}

// The jsii proxy struct for ThumbnailConfiguration
type jsiiProxy_ThumbnailConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ThumbnailConfiguration) RecordingMode() RecordingMode {
	var returns RecordingMode
	_jsii_.Get(
		j,
		"recordingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ThumbnailConfiguration) Resolution() Resolution {
	var returns Resolution
	_jsii_.Get(
		j,
		"resolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ThumbnailConfiguration) Storage() *[]Storage {
	var returns *[]Storage
	_jsii_.Get(
		j,
		"storage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ThumbnailConfiguration) TargetInterval() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"targetInterval",
		&returns,
	)
	return returns
}


// Disable the generation of thumbnails for recorded video.
// Experimental.
func ThumbnailConfiguration_Disable() ThumbnailConfiguration {
	_init_.Initialize()

	var returns ThumbnailConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.ThumbnailConfiguration",
		"disable",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Enable the generation of thumbnails for recorded video at a time interval.
// Experimental.
func ThumbnailConfiguration_Interval(resolution Resolution, storage *[]Storage, targetInterval awscdk.Duration) ThumbnailConfiguration {
	_init_.Initialize()

	var returns ThumbnailConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-ivs-alpha.ThumbnailConfiguration",
		"interval",
		[]interface{}{resolution, storage, targetInterval},
		&returns,
	)

	return returns
}


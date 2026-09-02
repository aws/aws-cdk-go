package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The method MediaLive uses to synchronise pipelines for pipeline output locking.
//
// Example:
//   // Video-aligned pipeline locking — useful when sources lack reliable timecodes
//   locking := medialive.OutputLocking_Pipeline(&PipelineOutputLockingProps{
//   	Method: medialive.PipelineLockingMethod_VIDEO_ALIGNMENT(),
//   })
//
// Experimental.
type PipelineLockingMethod interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for PipelineLockingMethod
type jsiiProxy_PipelineLockingMethod struct {
	_ byte // padding
}

func (j *jsiiProxy_PipelineLockingMethod) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func PipelineLockingMethod_Of(value *string) PipelineLockingMethod {
	_init_.Initialize()

	if err := validatePipelineLockingMethod_OfParameters(value); err != nil {
		panic(err)
	}
	var returns PipelineLockingMethod

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.PipelineLockingMethod",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func PipelineLockingMethod_SOURCE_TIMECODE() PipelineLockingMethod {
	_init_.Initialize()
	var returns PipelineLockingMethod
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.PipelineLockingMethod",
		"SOURCE_TIMECODE",
		&returns,
	)
	return returns
}

func PipelineLockingMethod_VIDEO_ALIGNMENT() PipelineLockingMethod {
	_init_.Initialize()
	var returns PipelineLockingMethod
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.PipelineLockingMethod",
		"VIDEO_ALIGNMENT",
		&returns,
	)
	return returns
}


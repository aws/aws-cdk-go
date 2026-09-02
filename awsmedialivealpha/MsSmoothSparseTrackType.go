package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth sparse track type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothSparseTrackType := medialive_alpha.MsSmoothSparseTrackType_NONE()
//
// Experimental.
type MsSmoothSparseTrackType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothSparseTrackType
type jsiiProxy_MsSmoothSparseTrackType struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothSparseTrackType) Value() *string {
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
func MsSmoothSparseTrackType_Of(value *string) MsSmoothSparseTrackType {
	_init_.Initialize()

	if err := validateMsSmoothSparseTrackType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothSparseTrackType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSparseTrackType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothSparseTrackType_NONE() MsSmoothSparseTrackType {
	_init_.Initialize()
	var returns MsSmoothSparseTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSparseTrackType",
		"NONE",
		&returns,
	)
	return returns
}

func MsSmoothSparseTrackType_SCTE_35() MsSmoothSparseTrackType {
	_init_.Initialize()
	var returns MsSmoothSparseTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSparseTrackType",
		"SCTE_35",
		&returns,
	)
	return returns
}

func MsSmoothSparseTrackType_SCTE_35_WITHOUT_SEGMENTATION() MsSmoothSparseTrackType {
	_init_.Initialize()
	var returns MsSmoothSparseTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSparseTrackType",
		"SCTE_35_WITHOUT_SEGMENTATION",
		&returns,
	)
	return returns
}


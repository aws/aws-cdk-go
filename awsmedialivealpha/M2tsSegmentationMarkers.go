package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The type of segmentation markers to insert.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsSegmentationMarkers := medialive_alpha.M2tsSegmentationMarkers_EBP()
//
// Experimental.
type M2tsSegmentationMarkers interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsSegmentationMarkers
type jsiiProxy_M2tsSegmentationMarkers struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsSegmentationMarkers) Value() *string {
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
func M2tsSegmentationMarkers_Of(value *string) M2tsSegmentationMarkers {
	_init_.Initialize()

	if err := validateM2tsSegmentationMarkers_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsSegmentationMarkers

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsSegmentationMarkers_EBP() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"EBP",
		&returns,
	)
	return returns
}

func M2tsSegmentationMarkers_EBP_LEGACY() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"EBP_LEGACY",
		&returns,
	)
	return returns
}

func M2tsSegmentationMarkers_NONE() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"NONE",
		&returns,
	)
	return returns
}

func M2tsSegmentationMarkers_PSI_SEGSTART() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"PSI_SEGSTART",
		&returns,
	)
	return returns
}

func M2tsSegmentationMarkers_RAI_ADAPT() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"RAI_ADAPT",
		&returns,
	)
	return returns
}

func M2tsSegmentationMarkers_RAI_SEGSTART() M2tsSegmentationMarkers {
	_init_.Initialize()
	var returns M2tsSegmentationMarkers
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsSegmentationMarkers",
		"RAI_SEGSTART",
		&returns,
	)
	return returns
}


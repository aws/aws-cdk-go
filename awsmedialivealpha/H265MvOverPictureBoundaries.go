package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 motion vector over picture boundaries.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265MvOverPictureBoundaries := medialive_alpha.H265MvOverPictureBoundaries_Of(jsii.String("value"))
//
// Experimental.
type H265MvOverPictureBoundaries interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265MvOverPictureBoundaries
type jsiiProxy_H265MvOverPictureBoundaries struct {
	_ byte // padding
}

func (j *jsiiProxy_H265MvOverPictureBoundaries) Value() *string {
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
func H265MvOverPictureBoundaries_Of(value *string) H265MvOverPictureBoundaries {
	_init_.Initialize()

	if err := validateH265MvOverPictureBoundaries_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265MvOverPictureBoundaries

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265MvOverPictureBoundaries",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265MvOverPictureBoundaries_DISABLED() H265MvOverPictureBoundaries {
	_init_.Initialize()
	var returns H265MvOverPictureBoundaries
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265MvOverPictureBoundaries",
		"DISABLED",
		&returns,
	)
	return returns
}

func H265MvOverPictureBoundaries_ENABLED() H265MvOverPictureBoundaries {
	_init_.Initialize()
	var returns H265MvOverPictureBoundaries
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265MvOverPictureBoundaries",
		"ENABLED",
		&returns,
	)
	return returns
}


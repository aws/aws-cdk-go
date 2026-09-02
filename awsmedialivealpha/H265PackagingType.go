package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 packaging type for HLS/MS Smooth outputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265PackagingType := medialive_alpha.H265PackagingType_Of(jsii.String("value"))
//
// Experimental.
type H265PackagingType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265PackagingType
type jsiiProxy_H265PackagingType struct {
	_ byte // padding
}

func (j *jsiiProxy_H265PackagingType) Value() *string {
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
func H265PackagingType_Of(value *string) H265PackagingType {
	_init_.Initialize()

	if err := validateH265PackagingType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265PackagingType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265PackagingType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265PackagingType_HEV1() H265PackagingType {
	_init_.Initialize()
	var returns H265PackagingType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265PackagingType",
		"HEV1",
		&returns,
	)
	return returns
}

func H265PackagingType_HVC1() H265PackagingType {
	_init_.Initialize()
	var returns H265PackagingType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265PackagingType",
		"HVC1",
		&returns,
	)
	return returns
}


package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Scan type for the output video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   scanType := medialive_alpha.ScanType_Of(jsii.String("value"))
//
// Experimental.
type ScanType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ScanType
type jsiiProxy_ScanType struct {
	_ byte // padding
}

func (j *jsiiProxy_ScanType) Value() *string {
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
func ScanType_Of(value *string) ScanType {
	_init_.Initialize()

	if err := validateScanType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ScanType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ScanType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ScanType_INTERLACED() ScanType {
	_init_.Initialize()
	var returns ScanType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ScanType",
		"INTERLACED",
		&returns,
	)
	return returns
}

func ScanType_PROGRESSIVE() ScanType {
	_init_.Initialize()
	var returns ScanType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ScanType",
		"PROGRESSIVE",
		&returns,
	)
	return returns
}


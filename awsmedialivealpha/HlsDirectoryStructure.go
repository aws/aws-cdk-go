package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS directory structure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsDirectoryStructure := medialive_alpha.HlsDirectoryStructure_Of(jsii.String("value"))
//
// Experimental.
type HlsDirectoryStructure interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsDirectoryStructure
type jsiiProxy_HlsDirectoryStructure struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsDirectoryStructure) Value() *string {
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
func HlsDirectoryStructure_Of(value *string) HlsDirectoryStructure {
	_init_.Initialize()

	if err := validateHlsDirectoryStructure_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsDirectoryStructure

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsDirectoryStructure",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsDirectoryStructure_SINGLE_DIRECTORY() HlsDirectoryStructure {
	_init_.Initialize()
	var returns HlsDirectoryStructure
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsDirectoryStructure",
		"SINGLE_DIRECTORY",
		&returns,
	)
	return returns
}

func HlsDirectoryStructure_SUBDIRECTORY_PER_STREAM() HlsDirectoryStructure {
	_init_.Initialize()
	var returns HlsDirectoryStructure
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsDirectoryStructure",
		"SUBDIRECTORY_PER_STREAM",
		&returns,
	)
	return returns
}


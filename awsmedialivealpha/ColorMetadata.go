package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Color metadata inclusion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   colorMetadata := medialive_alpha.ColorMetadata_Of(jsii.String("value"))
//
// Experimental.
type ColorMetadata interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for ColorMetadata
type jsiiProxy_ColorMetadata struct {
	_ byte // padding
}

func (j *jsiiProxy_ColorMetadata) Value() *string {
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
func ColorMetadata_Of(value *string) ColorMetadata {
	_init_.Initialize()

	if err := validateColorMetadata_OfParameters(value); err != nil {
		panic(err)
	}
	var returns ColorMetadata

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.ColorMetadata",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func ColorMetadata_IGNORE() ColorMetadata {
	_init_.Initialize()
	var returns ColorMetadata
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorMetadata",
		"IGNORE",
		&returns,
	)
	return returns
}

func ColorMetadata_INSERT() ColorMetadata {
	_init_.Initialize()
	var returns ColorMetadata
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.ColorMetadata",
		"INSERT",
		&returns,
	)
	return returns
}


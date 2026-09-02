package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether to use chunked transfer encoding for an HLS CDN connection (Akamai, WebDAV).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   httpTransferMode := medialive_alpha.HttpTransferMode_Of(jsii.String("value"))
//
// Experimental.
type HttpTransferMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HttpTransferMode
type jsiiProxy_HttpTransferMode struct {
	_ byte // padding
}

func (j *jsiiProxy_HttpTransferMode) Value() *string {
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
func HttpTransferMode_Of(value *string) HttpTransferMode {
	_init_.Initialize()

	if err := validateHttpTransferMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HttpTransferMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HttpTransferMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HttpTransferMode_CHUNKED() HttpTransferMode {
	_init_.Initialize()
	var returns HttpTransferMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HttpTransferMode",
		"CHUNKED",
		&returns,
	)
	return returns
}

func HttpTransferMode_NON_CHUNKED() HttpTransferMode {
	_init_.Initialize()
	var returns HttpTransferMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HttpTransferMode",
		"NON_CHUNKED",
		&returns,
	)
	return returns
}


package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import mediapackagev2_alpha "github.com/aws/aws-cdk-go/awsmediapackagev2alpha"
//
//   headersCMSD := mediapackagev2_alpha.HeadersCMSD_Of(jsii.String("value"))
//
// Experimental.
type HeadersCMSD interface {
	// The string value of the CMSD header.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HeadersCMSD
type jsiiProxy_HeadersCMSD struct {
	_ byte // padding
}

func (j *jsiiProxy_HeadersCMSD) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Escape hatch for new CMSD headers not yet supported by the construct.
// Experimental.
func HeadersCMSD_Of(value *string) HeadersCMSD {
	_init_.Initialize()

	if err := validateHeadersCMSD_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HeadersCMSD

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.HeadersCMSD",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HeadersCMSD_MQCS() HeadersCMSD {
	_init_.Initialize()
	var returns HeadersCMSD
	_jsii_.StaticGet(
		"@aws-cdk/aws-mediapackagev2-alpha.HeadersCMSD",
		"MQCS",
		&returns,
	)
	return returns
}


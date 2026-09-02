package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CMAF Ingest KLV behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   klvBehavior := medialive_alpha.KlvBehavior_Of(jsii.String("value"))
//
// Experimental.
type KlvBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for KlvBehavior
type jsiiProxy_KlvBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_KlvBehavior) Value() *string {
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
func KlvBehavior_Of(value *string) KlvBehavior {
	_init_.Initialize()

	if err := validateKlvBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns KlvBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.KlvBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func KlvBehavior_NO_PASSTHROUGH() KlvBehavior {
	_init_.Initialize()
	var returns KlvBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.KlvBehavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func KlvBehavior_PASSTHROUGH() KlvBehavior {
	_init_.Initialize()
	var returns KlvBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.KlvBehavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}


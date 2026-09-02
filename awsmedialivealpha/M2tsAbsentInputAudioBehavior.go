package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Behavior when the selected input audio stream is removed from the input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsAbsentInputAudioBehavior := medialive_alpha.M2tsAbsentInputAudioBehavior_Of(jsii.String("value"))
//
// Experimental.
type M2tsAbsentInputAudioBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsAbsentInputAudioBehavior
type jsiiProxy_M2tsAbsentInputAudioBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsAbsentInputAudioBehavior) Value() *string {
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
func M2tsAbsentInputAudioBehavior_Of(value *string) M2tsAbsentInputAudioBehavior {
	_init_.Initialize()

	if err := validateM2tsAbsentInputAudioBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsAbsentInputAudioBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsAbsentInputAudioBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsAbsentInputAudioBehavior_DROP() M2tsAbsentInputAudioBehavior {
	_init_.Initialize()
	var returns M2tsAbsentInputAudioBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAbsentInputAudioBehavior",
		"DROP",
		&returns,
	)
	return returns
}

func M2tsAbsentInputAudioBehavior_ENCODE_SILENCE() M2tsAbsentInputAudioBehavior {
	_init_.Initialize()
	var returns M2tsAbsentInputAudioBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsAbsentInputAudioBehavior",
		"ENCODE_SILENCE",
		&returns,
	)
	return returns
}


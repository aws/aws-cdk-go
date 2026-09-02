package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AV1 timecode insertion.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1TimecodeInsertion := medialive_alpha.Av1TimecodeInsertion_Of(jsii.String("value"))
//
// Experimental.
type Av1TimecodeInsertion interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Av1TimecodeInsertion
type jsiiProxy_Av1TimecodeInsertion struct {
	_ byte // padding
}

func (j *jsiiProxy_Av1TimecodeInsertion) Value() *string {
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
func Av1TimecodeInsertion_Of(value *string) Av1TimecodeInsertion {
	_init_.Initialize()

	if err := validateAv1TimecodeInsertion_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Av1TimecodeInsertion

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1TimecodeInsertion",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Av1TimecodeInsertion_DISABLED() Av1TimecodeInsertion {
	_init_.Initialize()
	var returns Av1TimecodeInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1TimecodeInsertion",
		"DISABLED",
		&returns,
	)
	return returns
}

func Av1TimecodeInsertion_METADATA_OBU() Av1TimecodeInsertion {
	_init_.Initialize()
	var returns Av1TimecodeInsertion
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Av1TimecodeInsertion",
		"METADATA_OBU",
		&returns,
	)
	return returns
}


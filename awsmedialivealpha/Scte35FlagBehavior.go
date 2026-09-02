package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// How to handle SCTE-35 regional blackout and web delivery flags.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   scte35FlagBehavior := medialive_alpha.Scte35FlagBehavior_Of(jsii.String("value"))
//
// Experimental.
type Scte35FlagBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Scte35FlagBehavior
type jsiiProxy_Scte35FlagBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_Scte35FlagBehavior) Value() *string {
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
func Scte35FlagBehavior_Of(value *string) Scte35FlagBehavior {
	_init_.Initialize()

	if err := validateScte35FlagBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Scte35FlagBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Scte35FlagBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Scte35FlagBehavior_FOLLOW() Scte35FlagBehavior {
	_init_.Initialize()
	var returns Scte35FlagBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35FlagBehavior",
		"FOLLOW",
		&returns,
	)
	return returns
}

func Scte35FlagBehavior_IGNORE() Scte35FlagBehavior {
	_init_.Initialize()
	var returns Scte35FlagBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Scte35FlagBehavior",
		"IGNORE",
		&returns,
	)
	return returns
}


package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Timed metadata passthrough behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsTimedMetadataBehavior := medialive_alpha.M2tsTimedMetadataBehavior_Of(jsii.String("value"))
//
// Experimental.
type M2tsTimedMetadataBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsTimedMetadataBehavior
type jsiiProxy_M2tsTimedMetadataBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsTimedMetadataBehavior) Value() *string {
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
func M2tsTimedMetadataBehavior_Of(value *string) M2tsTimedMetadataBehavior {
	_init_.Initialize()

	if err := validateM2tsTimedMetadataBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsTimedMetadataBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsTimedMetadataBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsTimedMetadataBehavior_NO_PASSTHROUGH() M2tsTimedMetadataBehavior {
	_init_.Initialize()
	var returns M2tsTimedMetadataBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsTimedMetadataBehavior",
		"NO_PASSTHROUGH",
		&returns,
	)
	return returns
}

func M2tsTimedMetadataBehavior_PASSTHROUGH() M2tsTimedMetadataBehavior {
	_init_.Initialize()
	var returns M2tsTimedMetadataBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsTimedMetadataBehavior",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}


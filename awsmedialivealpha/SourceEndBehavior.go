package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The source end behavior for file-based inputs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   sourceEndBehavior := medialive_alpha.SourceEndBehavior_Of(jsii.String("value"))
//
// Experimental.
type SourceEndBehavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SourceEndBehavior
type jsiiProxy_SourceEndBehavior struct {
	_ byte // padding
}

func (j *jsiiProxy_SourceEndBehavior) Value() *string {
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
func SourceEndBehavior_Of(value *string) SourceEndBehavior {
	_init_.Initialize()

	if err := validateSourceEndBehavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SourceEndBehavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SourceEndBehavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SourceEndBehavior_CONTINUE() SourceEndBehavior {
	_init_.Initialize()
	var returns SourceEndBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SourceEndBehavior",
		"CONTINUE",
		&returns,
	)
	return returns
}

func SourceEndBehavior_LOOP() SourceEndBehavior {
	_init_.Initialize()
	var returns SourceEndBehavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SourceEndBehavior",
		"LOOP",
		&returns,
	)
	return returns
}


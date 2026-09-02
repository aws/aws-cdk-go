package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// ID3 metadata insertion behavior (CMAF Ingest and MediaPackage V2 output groups).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   id3Behavior := medialive_alpha.Id3Behavior_Of(jsii.String("value"))
//
// Experimental.
type Id3Behavior interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Id3Behavior
type jsiiProxy_Id3Behavior struct {
	_ byte // padding
}

func (j *jsiiProxy_Id3Behavior) Value() *string {
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
func Id3Behavior_Of(value *string) Id3Behavior {
	_init_.Initialize()

	if err := validateId3Behavior_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Id3Behavior

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Id3Behavior",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Id3Behavior_DISABLED() Id3Behavior {
	_init_.Initialize()
	var returns Id3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Id3Behavior",
		"DISABLED",
		&returns,
	)
	return returns
}

func Id3Behavior_ENABLED() Id3Behavior {
	_init_.Initialize()
	var returns Id3Behavior
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Id3Behavior",
		"ENABLED",
		&returns,
	)
	return returns
}


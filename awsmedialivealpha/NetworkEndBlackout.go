package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Network end blackout state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   networkEndBlackout := medialive_alpha.NetworkEndBlackout_Of(jsii.String("value"))
//
// Experimental.
type NetworkEndBlackout interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NetworkEndBlackout
type jsiiProxy_NetworkEndBlackout struct {
	_ byte // padding
}

func (j *jsiiProxy_NetworkEndBlackout) Value() *string {
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
func NetworkEndBlackout_Of(value *string) NetworkEndBlackout {
	_init_.Initialize()

	if err := validateNetworkEndBlackout_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NetworkEndBlackout

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NetworkEndBlackout",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NetworkEndBlackout_DISABLED() NetworkEndBlackout {
	_init_.Initialize()
	var returns NetworkEndBlackout
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NetworkEndBlackout",
		"DISABLED",
		&returns,
	)
	return returns
}

func NetworkEndBlackout_ENABLED() NetworkEndBlackout {
	_init_.Initialize()
	var returns NetworkEndBlackout
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NetworkEndBlackout",
		"ENABLED",
		&returns,
	)
	return returns
}


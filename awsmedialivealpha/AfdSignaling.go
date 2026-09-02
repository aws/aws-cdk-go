package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AFD signaling mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   afdSignaling := medialive_alpha.AfdSignaling_AUTO()
//
// Experimental.
type AfdSignaling interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AfdSignaling
type jsiiProxy_AfdSignaling struct {
	_ byte // padding
}

func (j *jsiiProxy_AfdSignaling) Value() *string {
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
func AfdSignaling_Of(value *string) AfdSignaling {
	_init_.Initialize()

	if err := validateAfdSignaling_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AfdSignaling

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AfdSignaling",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AfdSignaling_AUTO() AfdSignaling {
	_init_.Initialize()
	var returns AfdSignaling
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AfdSignaling",
		"AUTO",
		&returns,
	)
	return returns
}

func AfdSignaling_FIXED() AfdSignaling {
	_init_.Initialize()
	var returns AfdSignaling
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AfdSignaling",
		"FIXED",
		&returns,
	)
	return returns
}

func AfdSignaling_NONE() AfdSignaling {
	_init_.Initialize()
	var returns AfdSignaling
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AfdSignaling",
		"NONE",
		&returns,
	)
	return returns
}


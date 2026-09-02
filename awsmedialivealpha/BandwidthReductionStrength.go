package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Bandwidth reduction filter strength.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   bandwidthReductionStrength := medialive_alpha.BandwidthReductionStrength_AUTO()
//
// Experimental.
type BandwidthReductionStrength interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for BandwidthReductionStrength
type jsiiProxy_BandwidthReductionStrength struct {
	_ byte // padding
}

func (j *jsiiProxy_BandwidthReductionStrength) Value() *string {
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
func BandwidthReductionStrength_Of(value *string) BandwidthReductionStrength {
	_init_.Initialize()

	if err := validateBandwidthReductionStrength_OfParameters(value); err != nil {
		panic(err)
	}
	var returns BandwidthReductionStrength

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func BandwidthReductionStrength_AUTO() BandwidthReductionStrength {
	_init_.Initialize()
	var returns BandwidthReductionStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"AUTO",
		&returns,
	)
	return returns
}

func BandwidthReductionStrength_STRENGTH_1() BandwidthReductionStrength {
	_init_.Initialize()
	var returns BandwidthReductionStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"STRENGTH_1",
		&returns,
	)
	return returns
}

func BandwidthReductionStrength_STRENGTH_2() BandwidthReductionStrength {
	_init_.Initialize()
	var returns BandwidthReductionStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"STRENGTH_2",
		&returns,
	)
	return returns
}

func BandwidthReductionStrength_STRENGTH_3() BandwidthReductionStrength {
	_init_.Initialize()
	var returns BandwidthReductionStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"STRENGTH_3",
		&returns,
	)
	return returns
}

func BandwidthReductionStrength_STRENGTH_4() BandwidthReductionStrength {
	_init_.Initialize()
	var returns BandwidthReductionStrength
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionStrength",
		"STRENGTH_4",
		&returns,
	)
	return returns
}


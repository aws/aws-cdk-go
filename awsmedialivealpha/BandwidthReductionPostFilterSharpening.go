package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Post-filter sharpening for bandwidth reduction filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   bandwidthReductionPostFilterSharpening := medialive_alpha.BandwidthReductionPostFilterSharpening_DISABLED()
//
// Experimental.
type BandwidthReductionPostFilterSharpening interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for BandwidthReductionPostFilterSharpening
type jsiiProxy_BandwidthReductionPostFilterSharpening struct {
	_ byte // padding
}

func (j *jsiiProxy_BandwidthReductionPostFilterSharpening) Value() *string {
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
func BandwidthReductionPostFilterSharpening_Of(value *string) BandwidthReductionPostFilterSharpening {
	_init_.Initialize()

	if err := validateBandwidthReductionPostFilterSharpening_OfParameters(value); err != nil {
		panic(err)
	}
	var returns BandwidthReductionPostFilterSharpening

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionPostFilterSharpening",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func BandwidthReductionPostFilterSharpening_DISABLED() BandwidthReductionPostFilterSharpening {
	_init_.Initialize()
	var returns BandwidthReductionPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionPostFilterSharpening",
		"DISABLED",
		&returns,
	)
	return returns
}

func BandwidthReductionPostFilterSharpening_SHARPENING_1() BandwidthReductionPostFilterSharpening {
	_init_.Initialize()
	var returns BandwidthReductionPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionPostFilterSharpening",
		"SHARPENING_1",
		&returns,
	)
	return returns
}

func BandwidthReductionPostFilterSharpening_SHARPENING_2() BandwidthReductionPostFilterSharpening {
	_init_.Initialize()
	var returns BandwidthReductionPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionPostFilterSharpening",
		"SHARPENING_2",
		&returns,
	)
	return returns
}

func BandwidthReductionPostFilterSharpening_SHARPENING_3() BandwidthReductionPostFilterSharpening {
	_init_.Initialize()
	var returns BandwidthReductionPostFilterSharpening
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.BandwidthReductionPostFilterSharpening",
		"SHARPENING_3",
		&returns,
	)
	return returns
}


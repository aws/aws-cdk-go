package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 motion vector temporal predictor.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265MvTemporalPredictor := medialive_alpha.H265MvTemporalPredictor_Of(jsii.String("value"))
//
// Experimental.
type H265MvTemporalPredictor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265MvTemporalPredictor
type jsiiProxy_H265MvTemporalPredictor struct {
	_ byte // padding
}

func (j *jsiiProxy_H265MvTemporalPredictor) Value() *string {
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
func H265MvTemporalPredictor_Of(value *string) H265MvTemporalPredictor {
	_init_.Initialize()

	if err := validateH265MvTemporalPredictor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265MvTemporalPredictor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265MvTemporalPredictor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265MvTemporalPredictor_DISABLED() H265MvTemporalPredictor {
	_init_.Initialize()
	var returns H265MvTemporalPredictor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265MvTemporalPredictor",
		"DISABLED",
		&returns,
	)
	return returns
}

func H265MvTemporalPredictor_ENABLED() H265MvTemporalPredictor {
	_init_.Initialize()
	var returns H265MvTemporalPredictor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265MvTemporalPredictor",
		"ENABLED",
		&returns,
	)
	return returns
}


package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The pixel aspect ratio (PAR) of the video.
//
// Use the predefined constants for standard ratios, or {@link PixelAspectRatio.of} for
// a custom ratio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   pixelAspectRatio := medialive_alpha.PixelAspectRatio_Of(jsii.Number(123), jsii.Number(123))
//
// Experimental.
type PixelAspectRatio interface {
	// Returns the string value in `numerator:denominator` form.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PixelAspectRatio
type jsiiProxy_PixelAspectRatio struct {
	_ byte // padding
}

// Define a pixel aspect ratio.
// Experimental.
func PixelAspectRatio_Of(numerator *float64, denominator *float64) PixelAspectRatio {
	_init_.Initialize()

	if err := validatePixelAspectRatio_OfParameters(numerator, denominator); err != nil {
		panic(err)
	}
	var returns PixelAspectRatio

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.PixelAspectRatio",
		"of",
		[]interface{}{numerator, denominator},
		&returns,
	)

	return returns
}

func PixelAspectRatio_SQUARE() PixelAspectRatio {
	_init_.Initialize()
	var returns PixelAspectRatio
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.PixelAspectRatio",
		"SQUARE",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PixelAspectRatio) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


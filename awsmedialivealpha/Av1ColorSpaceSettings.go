package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Color space settings for AV1 video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   av1ColorSpaceSettings := medialive_alpha.Av1ColorSpaceSettings_Hdr10(&Hdr10SettingsProps{
//   	MaxCll: jsii.Number(123),
//   	MaxFall: jsii.Number(123),
//   })
//
// Experimental.
type Av1ColorSpaceSettings interface {
}

// The jsii proxy struct for Av1ColorSpaceSettings
type jsiiProxy_Av1ColorSpaceSettings struct {
	_ byte // padding
}

// HDR10 color space.
// Experimental.
func Av1ColorSpaceSettings_Hdr10(props *Hdr10SettingsProps) Av1ColorSpaceSettings {
	_init_.Initialize()

	if err := validateAv1ColorSpaceSettings_Hdr10Parameters(props); err != nil {
		panic(err)
	}
	var returns Av1ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1ColorSpaceSettings",
		"hdr10",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// HLG 2020 color space.
// Experimental.
func Av1ColorSpaceSettings_Hlg2020() Av1ColorSpaceSettings {
	_init_.Initialize()

	var returns Av1ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1ColorSpaceSettings",
		"hlg2020",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Pass through the source color space with no conversion.
// Experimental.
func Av1ColorSpaceSettings_Passthrough() Av1ColorSpaceSettings {
	_init_.Initialize()

	var returns Av1ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1ColorSpaceSettings",
		"passthrough",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.601 color space.
// Experimental.
func Av1ColorSpaceSettings_Rec601() Av1ColorSpaceSettings {
	_init_.Initialize()

	var returns Av1ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1ColorSpaceSettings",
		"rec601",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.709 color space.
// Experimental.
func Av1ColorSpaceSettings_Rec709() Av1ColorSpaceSettings {
	_init_.Initialize()

	var returns Av1ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Av1ColorSpaceSettings",
		"rec709",
		nil, // no parameters
		&returns,
	)

	return returns
}


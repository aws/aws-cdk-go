package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Color space settings for H.264 video.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264ColorSpaceSettings := medialive_alpha.H264ColorSpaceSettings_Passthrough()
//
// Experimental.
type H264ColorSpaceSettings interface {
}

// The jsii proxy struct for H264ColorSpaceSettings
type jsiiProxy_H264ColorSpaceSettings struct {
	_ byte // padding
}

// Pass through the source color space with no conversion.
// Experimental.
func H264ColorSpaceSettings_Passthrough() H264ColorSpaceSettings {
	_init_.Initialize()

	var returns H264ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264ColorSpaceSettings",
		"passthrough",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.601 color space.
// Experimental.
func H264ColorSpaceSettings_Rec601() H264ColorSpaceSettings {
	_init_.Initialize()

	var returns H264ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264ColorSpaceSettings",
		"rec601",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Convert to Rec.709 color space.
// Experimental.
func H264ColorSpaceSettings_Rec709() H264ColorSpaceSettings {
	_init_.Initialize()

	var returns H264ColorSpaceSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264ColorSpaceSettings",
		"rec709",
		nil, // no parameters
		&returns,
	)

	return returns
}


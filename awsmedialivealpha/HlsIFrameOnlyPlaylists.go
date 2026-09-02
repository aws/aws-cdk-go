package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS I-frame only playlists.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsIFrameOnlyPlaylists := medialive_alpha.HlsIFrameOnlyPlaylists_Of(jsii.String("value"))
//
// Experimental.
type HlsIFrameOnlyPlaylists interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsIFrameOnlyPlaylists
type jsiiProxy_HlsIFrameOnlyPlaylists struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsIFrameOnlyPlaylists) Value() *string {
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
func HlsIFrameOnlyPlaylists_Of(value *string) HlsIFrameOnlyPlaylists {
	_init_.Initialize()

	if err := validateHlsIFrameOnlyPlaylists_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsIFrameOnlyPlaylists

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsIFrameOnlyPlaylists",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsIFrameOnlyPlaylists_DISABLED() HlsIFrameOnlyPlaylists {
	_init_.Initialize()
	var returns HlsIFrameOnlyPlaylists
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIFrameOnlyPlaylists",
		"DISABLED",
		&returns,
	)
	return returns
}

func HlsIFrameOnlyPlaylists_STANDARD() HlsIFrameOnlyPlaylists {
	_init_.Initialize()
	var returns HlsIFrameOnlyPlaylists
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIFrameOnlyPlaylists",
		"STANDARD",
		&returns,
	)
	return returns
}


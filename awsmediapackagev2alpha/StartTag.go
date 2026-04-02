package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Helper class for creating EXT-X-START tags in HLS playlists.
//
// The EXT-X-START tag indicates a preferred point at which to start playing a playlist.
// Configuration for EXT-X-START tag in HLS playlists.
//
// Use the `StartTag` helper class to create instances with validation.
//
// Example:
//   var channel Channel
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   			StartTag: awsmediapackagev2alpha.StartTag_Of(jsii.Number(10)),
//   		}),
//   	},
//   })
//
// Experimental.
type StartTag interface {
	// Specify the value for PRECISE within your EXT-X-START tag.
	//
	// Leave blank, or choose false, to use the default value NO. Choose true to use the value YES.
	// Experimental.
	Precise() *bool
	// Specify the value for TIME-OFFSET within your EXT-X-START tag.
	//
	// Enter a signed floating point value which, if positive, must be less than the configured
	// manifest duration minus three times the configured segment target duration. If negative,
	// the absolute value must be larger than three times the configured segment target duration,
	// and the absolute value must be smaller than the configured manifest duration.
	// Experimental.
	TimeOffset() *float64
}

// The jsii proxy struct for StartTag
type jsiiProxy_StartTag struct {
	_ byte // padding
}

func (j *jsiiProxy_StartTag) Precise() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"precise",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StartTag) TimeOffset() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeOffset",
		&returns,
	)
	return returns
}


// Experimental.
func NewStartTag() StartTag {
	_init_.Initialize()

	j := jsiiProxy_StartTag{}

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTag",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStartTag_Override(s StartTag) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTag",
		nil, // no parameters
		s,
	)
}

// Create a start tag with a time offset.
// Experimental.
func StartTag_Of(timeOffset *float64, options *StartTagOptions) StartTag {
	_init_.Initialize()

	if err := validateStartTag_OfParameters(timeOffset, options); err != nil {
		panic(err)
	}
	var returns StartTag

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTag",
		"of",
		[]interface{}{timeOffset, options},
		&returns,
	)

	return returns
}

// Create a start tag with precise positioning enabled (PRECISE=YES).
// Experimental.
func StartTag_WithPrecise(timeOffset *float64) StartTag {
	_init_.Initialize()

	if err := validateStartTag_WithPreciseParameters(timeOffset); err != nil {
		panic(err)
	}
	var returns StartTag

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.StartTag",
		"withPrecise",
		[]interface{}{timeOffset},
		&returns,
	)

	return returns
}


package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CDN settings for HLS output groups.
//
// Example:
//   var bucket IBucket
//   var video EncodeConfiguration
//   var audio EncodeConfiguration
//
//
//   // HLS to S3
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls_s3"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_ToBucket(bucket, jsii.String("live/stream")),
//   	},
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("hls_out"),
//   		},
//   	},
//   })
//
//   // HLS to an HTTPS CDN origin.
//   medialive.OutputGroupConfiguration_Hls(&HlsOutputGroupProps{
//   	Name: jsii.String("hls-http"),
//   	Destinations: []OutputDestination{
//   		medialive.OutputDestination_Url(jsii.String("https://203.0.113.10/ingest/stream")),
//   	},
//   	HlsCdnSettings: medialive.HlsCdnSettings_BasicPut(),
//   	Outputs: []HlsOutputDefinition{
//   		&HlsOutputDefinition{
//   			Encodes: []EncodeConfiguration{
//   				video,
//   				audio,
//   			},
//   			OutputName: jsii.String("hls_out"),
//   		},
//   	},
//   })
//
// Experimental.
type HlsCdnSettings interface {
}

// The jsii proxy struct for HlsCdnSettings
type jsiiProxy_HlsCdnSettings struct {
	_ byte // padding
}

// Use Akamai as the CDN for HLS output.
// Experimental.
func HlsCdnSettings_Akamai(props *HlsAkamaiCdnProps) HlsCdnSettings {
	_init_.Initialize()

	if err := validateHlsCdnSettings_AkamaiParameters(props); err != nil {
		panic(err)
	}
	var returns HlsCdnSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCdnSettings",
		"akamai",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use a basic HTTP PUT for HLS output.
// Experimental.
func HlsCdnSettings_BasicPut(props *HlsBasicPutCdnProps) HlsCdnSettings {
	_init_.Initialize()

	if err := validateHlsCdnSettings_BasicPutParameters(props); err != nil {
		panic(err)
	}
	var returns HlsCdnSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCdnSettings",
		"basicPut",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use Amazon S3 as the CDN for HLS output.
// Experimental.
func HlsCdnSettings_S3(props *HlsS3CdnProps) HlsCdnSettings {
	_init_.Initialize()

	if err := validateHlsCdnSettings_S3Parameters(props); err != nil {
		panic(err)
	}
	var returns HlsCdnSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCdnSettings",
		"s3",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Use WebDAV as the CDN for HLS output.
// Experimental.
func HlsCdnSettings_Webdav(props *HlsWebdavCdnProps) HlsCdnSettings {
	_init_.Initialize()

	if err := validateHlsCdnSettings_WebdavParameters(props); err != nil {
		panic(err)
	}
	var returns HlsCdnSettings

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsCdnSettings",
		"webdav",
		[]interface{}{props},
		&returns,
	)

	return returns
}


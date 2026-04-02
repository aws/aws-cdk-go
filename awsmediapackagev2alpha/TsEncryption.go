package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encryption configuration for TS segments.
//
// Use `TsEncryption.speke()` to create an instance.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("TsEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ts(&TsSegmentProps{
//   		Encryption: awsmediapackagev2alpha.TsEncryption_Speke(&TsSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.TsEncryptionMethod_SAMPLE_AES,
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Hls(&HlsManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type TsEncryption interface {
	EncryptionConfiguration
}

// The jsii proxy struct for TsEncryption
type jsiiProxy_TsEncryption struct {
	jsiiProxy_EncryptionConfiguration
}

// Create a SPEKE-based encryption configuration for TS segments.
// Experimental.
func TsEncryption_Speke(props *TsSpekeEncryptionProps) TsEncryption {
	_init_.Initialize()

	if err := validateTsEncryption_SpekeParameters(props); err != nil {
		panic(err)
	}
	var returns TsEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.TsEncryption",
		"speke",
		[]interface{}{props},
		&returns,
	)

	return returns
}


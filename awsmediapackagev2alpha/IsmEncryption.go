package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encryption configuration for ISM (Microsoft Smooth Streaming) segments.
//
// ISM only supports CENC encryption with PlayReady DRM.
// Audio and video presets are always SHARED.
//
// Use `IsmEncryption.speke()` to create an instance.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("IsmEndpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Ism(&IsmSegmentProps{
//   		Encryption: awsmediapackagev2alpha.IsmEncryption_Speke(&IsmSpekeEncryptionProps{
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   		}),
//   	}),
//   	Manifests: []Manifest{
//   		awsmediapackagev2alpha.Manifest_Mss(&MssManifestConfiguration{
//   			ManifestName: jsii.String("index"),
//   		}),
//   	},
//   })
//
// Experimental.
type IsmEncryption interface {
	EncryptionConfiguration
}

// The jsii proxy struct for IsmEncryption
type jsiiProxy_IsmEncryption struct {
	jsiiProxy_EncryptionConfiguration
}

// Create a SPEKE-based encryption configuration for ISM segments.
// Experimental.
func IsmEncryption_Speke(props *IsmSpekeEncryptionProps) IsmEncryption {
	_init_.Initialize()

	if err := validateIsmEncryption_SpekeParameters(props); err != nil {
		panic(err)
	}
	var returns IsmEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.IsmEncryption",
		"speke",
		[]interface{}{props},
		&returns,
	)

	return returns
}


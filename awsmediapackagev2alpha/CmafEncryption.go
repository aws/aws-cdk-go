package awsmediapackagev2alpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmediapackagev2alpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Encryption configuration for CMAF segments.
//
// Use `CmafEncryption.speke()` to create an instance.
//
// Example:
//   var channel Channel
//   var spekeRole IRole
//   var certificate ICertificate
//
//
//   awsmediapackagev2alpha.NewOriginEndpoint(this, jsii.String("Endpoint"), &OriginEndpointProps{
//   	Channel: Channel,
//   	Segment: awsmediapackagev2alpha.Segment_Cmaf(&CmafSegmentProps{
//   		Encryption: awsmediapackagev2alpha.CmafEncryption_Speke(&CmafSpekeEncryptionProps{
//   			Method: awsmediapackagev2alpha.CmafEncryptionMethod_CBCS,
//   			DrmSystems: []CmafDrmSystem{
//   				awsmediapackagev2alpha.CmafDrmSystem_FAIRPLAY,
//   			},
//   			ResourceId: jsii.String("my-content-id"),
//   			Url: jsii.String("https://example.com/speke"),
//   			Role: spekeRole,
//   			Certificate: *Certificate,
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
type CmafEncryption interface {
	EncryptionConfiguration
}

// The jsii proxy struct for CmafEncryption
type jsiiProxy_CmafEncryption struct {
	jsiiProxy_EncryptionConfiguration
}

// Create a SPEKE-based encryption configuration for CMAF segments.
// Experimental.
func CmafEncryption_Speke(props *CmafSpekeEncryptionProps) CmafEncryption {
	_init_.Initialize()

	if err := validateCmafEncryption_SpekeParameters(props); err != nil {
		panic(err)
	}
	var returns CmafEncryption

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-mediapackagev2-alpha.CmafEncryption",
		"speke",
		[]interface{}{props},
		&returns,
	)

	return returns
}


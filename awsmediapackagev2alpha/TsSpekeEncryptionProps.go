package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for TS SPEKE encryption configuration.
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
type TsSpekeEncryptionProps struct {
	// The encryption method to use.
	// Experimental.
	Method TsEncryptionMethod `field:"required" json:"method" yaml:"method"`
	// The unique identifier for the content.
	//
	// The service sends this identifier to the key server to identify the current endpoint.
	// How unique you make this identifier depends on how fine-grained you want access controls to be.
	// The service does not permit you to use the same ID for two simultaneous encryption processes.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// IAM role for accessing the key provider API.
	//
	// This role must have a trust policy that allows MediaPackage to assume the role,
	// and it must have sufficient permissions to access the key retrieval URL.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// URL of the SPEKE key provider.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Audio encryption preset.
	// Default: PresetSpeke20Audio.PRESET_AUDIO_1
	//
	// Experimental.
	AudioPreset PresetSpeke20Audio `field:"optional" json:"audioPreset" yaml:"audioPreset"`
	// The ARN of the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint.
	//
	// For this feature to work, your DRM key provider must support content key encryption.
	// Default: - no content key encryption.
	//
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// Constant initialization vector (32-character hex string).
	//
	// A 128-bit, 16-byte hex value represented by a 32-character string,
	// used in conjunction with the key for encrypting content.
	// Default: - MediaPackage generates the IV.
	//
	// Experimental.
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The DRM systems to use for content protection.
	// Default: - FairPlay for SAMPLE_AES, Clear Key AES 128 for AES_128.
	//
	// Experimental.
	DrmSystems *[]TsDrmSystem `field:"optional" json:"drmSystems" yaml:"drmSystems"`
	// Key rotation interval.
	// Default: - no rotation.
	//
	// Experimental.
	KeyRotationInterval awscdk.Duration `field:"optional" json:"keyRotationInterval" yaml:"keyRotationInterval"`
	// Video encryption preset.
	// Default: PresetSpeke20Video.PRESET_VIDEO_1
	//
	// Experimental.
	VideoPreset PresetSpeke20Video `field:"optional" json:"videoPreset" yaml:"videoPreset"`
}


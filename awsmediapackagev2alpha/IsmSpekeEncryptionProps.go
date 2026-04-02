package awsmediapackagev2alpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Properties for ISM SPEKE encryption configuration.
//
// ISM only supports CENC encryption with PlayReady DRM.
// Key rotation and constant initialization vectors are not supported.
// Audio and video presets default to SHARED.
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
type IsmSpekeEncryptionProps struct {
	// The unique identifier for the content.
	// Experimental.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// IAM role for accessing the key provider API.
	// Experimental.
	Role awsiam.IRole `field:"required" json:"role" yaml:"role"`
	// URL of the SPEKE key provider.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// The ARN of the certificate that you imported to AWS Certificate Manager to add content key encryption to this endpoint.
	// Default: - no content key encryption.
	//
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The DRM systems to use for content protection.
	// Default: - [IsmDrmSystem.PLAYREADY]
	//
	// Experimental.
	DrmSystems *[]IsmDrmSystem `field:"optional" json:"drmSystems" yaml:"drmSystems"`
}


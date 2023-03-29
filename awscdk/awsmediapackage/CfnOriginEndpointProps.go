package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOriginEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpointProps := &CfnOriginEndpointProps{
//   	ChannelId: jsii.String("channelId"),
//   	Id: jsii.String("id"),
//
//   	// the properties below are optional
//   	Authorization: &AuthorizationProperty{
//   		CdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		SecretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	CmafPackage: &CmafPackageProperty{
//   		Encryption: &CmafEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				Id: jsii.String("id"),
//
//   				// the properties below are optional
//   				AdMarkers: jsii.String("adMarkers"),
//   				AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				AdTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				PlaylistType: jsii.String("playlistType"),
//   				PlaylistWindowSeconds: jsii.Number(123),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				Url: jsii.String("url"),
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentPrefix: jsii.String("segmentPrefix"),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	DashPackage: &DashPackageProperty{
//   		AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		AdTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		Encryption: &DashEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//
//   			// the properties below are optional
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		ManifestLayout: jsii.String("manifestLayout"),
//   		ManifestWindowSeconds: jsii.Number(123),
//   		MinBufferTimeSeconds: jsii.Number(123),
//   		MinUpdatePeriodSeconds: jsii.Number(123),
//   		PeriodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		Profile: jsii.String("profile"),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   		SuggestedPresentationDelaySeconds: jsii.Number(123),
//   		UtcTiming: jsii.String("utcTiming"),
//   		UtcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	Description: jsii.String("description"),
//   	HlsPackage: &HlsPackageProperty{
//   		AdMarkers: jsii.String("adMarkers"),
//   		AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		AdTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		Encryption: &HlsEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   			KeyRotationIntervalSeconds: jsii.Number(123),
//   			RepeatExtXKey: jsii.Boolean(false),
//   		},
//   		IncludeDvbSubtitles: jsii.Boolean(false),
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		PlaylistType: jsii.String("playlistType"),
//   		PlaylistWindowSeconds: jsii.Number(123),
//   		ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   		UseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	ManifestName: jsii.String("manifestName"),
//   	MssPackage: &MssPackageProperty{
//   		Encryption: &MssEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				ResourceId: jsii.String("resourceId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				CertificateArn: jsii.String("certificateArn"),
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//   		},
//   		ManifestWindowSeconds: jsii.Number(123),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		StreamSelection: &StreamSelectionProperty{
//   			MaxVideoBitsPerSecond: jsii.Number(123),
//   			MinVideoBitsPerSecond: jsii.Number(123),
//   			StreamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	Origination: jsii.String("origination"),
//   	StartoverWindowSeconds: jsii.Number(123),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimeDelaySeconds: jsii.Number(123),
//   	Whitelist: []*string{
//   		jsii.String("whitelist"),
//   	},
//   }
//
type CfnOriginEndpointProps struct {
	// The ID of the channel associated with this endpoint.
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The manifest ID is required and must be unique within the OriginEndpoint.
	//
	// The ID can't be changed after the endpoint is created.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Parameters for CDN authorization.
	Authorization interface{} `field:"optional" json:"authorization" yaml:"authorization"`
	// Parameters for Common Media Application Format (CMAF) packaging.
	CmafPackage interface{} `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// Parameters for DASH packaging.
	DashPackage interface{} `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// Any descriptive information that you want to add to the endpoint for future identification purposes.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Parameters for Apple HLS packaging.
	HlsPackage interface{} `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// A short string that's appended to the end of the endpoint URL to create a unique path to this endpoint.
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage interface{} `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// Controls video origination from this endpoint.
	//
	// Valid values:
	//
	// - `ALLOW` - enables this endpoint to serve content to requesting devices.
	// - `DENY` - prevents this endpoint from serving content. Denying origination is helpful for harvesting live-to-VOD assets. For more information about harvesting and origination, see [Live-to-VOD Requirements](https://docs.aws.amazon.com/mediapackage/latest/ug/ltov-reqmts.html) .
	Origination *string `field:"optional" json:"origination" yaml:"origination"`
	// Maximum duration (seconds) of content to retain for startover playback.
	//
	// Omit this attribute or enter `0` to indicate that startover playback is disabled for this endpoint.
	StartoverWindowSeconds *float64 `field:"optional" json:"startoverWindowSeconds" yaml:"startoverWindowSeconds"`
	// The tags to assign to the endpoint.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Minimum duration (seconds) of delay to enforce on the playback of live content.
	//
	// Omit this attribute or enter `0` to indicate that there is no time delay in effect for this endpoint.
	TimeDelaySeconds *float64 `field:"optional" json:"timeDelaySeconds" yaml:"timeDelaySeconds"`
	// The IP addresses that can access this endpoint.
	Whitelist *[]*string `field:"optional" json:"whitelist" yaml:"whitelist"`
}


package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnOriginEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOriginEndpointProps := &cfnOriginEndpointProps{
//   	channelId: jsii.String("channelId"),
//   	id: jsii.String("id"),
//
//   	// the properties below are optional
//   	authorization: &authorizationProperty{
//   		cdnIdentifierSecret: jsii.String("cdnIdentifierSecret"),
//   		secretsRoleArn: jsii.String("secretsRoleArn"),
//   	},
//   	cmafPackage: &cmafPackageProperty{
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				id: jsii.String("id"),
//
//   				// the properties below are optional
//   				adMarkers: jsii.String("adMarkers"),
//   				adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   				adTriggers: []*string{
//   					jsii.String("adTriggers"),
//   				},
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				playlistType: jsii.String("playlistType"),
//   				playlistWindowSeconds: jsii.Number(123),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				url: jsii.String("url"),
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentPrefix: jsii.String("segmentPrefix"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	dashPackage: &dashPackageProperty{
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   		},
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		manifestLayout: jsii.String("manifestLayout"),
//   		manifestWindowSeconds: jsii.Number(123),
//   		minBufferTimeSeconds: jsii.Number(123),
//   		minUpdatePeriodSeconds: jsii.Number(123),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		profile: jsii.String("profile"),
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		suggestedPresentationDelaySeconds: jsii.Number(123),
//   		utcTiming: jsii.String("utcTiming"),
//   		utcTimingUri: jsii.String("utcTimingUri"),
//   	},
//   	description: jsii.String("description"),
//   	hlsPackage: &hlsPackageProperty{
//   		adMarkers: jsii.String("adMarkers"),
//   		adsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   		adTriggers: []*string{
//   			jsii.String("adTriggers"),
//   		},
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   			keyRotationIntervalSeconds: jsii.Number(123),
//   			repeatExtXKey: jsii.Boolean(false),
//   		},
//   		includeDvbSubtitles: jsii.Boolean(false),
//   		includeIframeOnlyStream: jsii.Boolean(false),
//   		playlistType: jsii.String("playlistType"),
//   		playlistWindowSeconds: jsii.Number(123),
//   		programDateTimeIntervalSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	manifestName: jsii.String("manifestName"),
//   	mssPackage: &mssPackageProperty{
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				resourceId: jsii.String("resourceId"),
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				certificateArn: jsii.String("certificateArn"),
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		manifestWindowSeconds: jsii.Number(123),
//   		segmentDurationSeconds: jsii.Number(123),
//   		streamSelection: &streamSelectionProperty{
//   			maxVideoBitsPerSecond: jsii.Number(123),
//   			minVideoBitsPerSecond: jsii.Number(123),
//   			streamOrder: jsii.String("streamOrder"),
//   		},
//   	},
//   	origination: jsii.String("origination"),
//   	startoverWindowSeconds: jsii.Number(123),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	timeDelaySeconds: jsii.Number(123),
//   	whitelist: []*string{
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


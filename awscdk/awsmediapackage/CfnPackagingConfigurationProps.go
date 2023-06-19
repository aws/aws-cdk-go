package awsmediapackage

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnPackagingConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackagingConfigurationProps := &CfnPackagingConfigurationProps{
//   	Id: jsii.String("id"),
//   	PackagingGroupId: jsii.String("packagingGroupId"),
//
//   	// the properties below are optional
//   	CmafPackage: &CmafPackageProperty{
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				AdMarkers: jsii.String("adMarkers"),
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &CmafEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	DashPackage: &DashPackageProperty{
//   		DashManifests: []interface{}{
//   			&DashManifestProperty{
//   				ManifestLayout: jsii.String("manifestLayout"),
//   				ManifestName: jsii.String("manifestName"),
//   				MinBufferTimeSeconds: jsii.Number(123),
//   				Profile: jsii.String("profile"),
//   				ScteMarkersSource: jsii.String("scteMarkersSource"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &DashEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//   		},
//   		IncludeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		IncludeIframeOnlyStream: jsii.Boolean(false),
//   		PeriodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   		SegmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	HlsPackage: &HlsPackageProperty{
//   		HlsManifests: []interface{}{
//   			&HlsManifestProperty{
//   				AdMarkers: jsii.String("adMarkers"),
//   				IncludeIframeOnlyStream: jsii.Boolean(false),
//   				ManifestName: jsii.String("manifestName"),
//   				ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   				RepeatExtXKey: jsii.Boolean(false),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &HlsEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//
//   			// the properties below are optional
//   			ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   			EncryptionMethod: jsii.String("encryptionMethod"),
//   		},
//   		IncludeDvbSubtitles: jsii.Boolean(false),
//   		SegmentDurationSeconds: jsii.Number(123),
//   		UseAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	MssPackage: &MssPackageProperty{
//   		MssManifests: []interface{}{
//   			&MssManifestProperty{
//   				ManifestName: jsii.String("manifestName"),
//   				StreamSelection: &StreamSelectionProperty{
//   					MaxVideoBitsPerSecond: jsii.Number(123),
//   					MinVideoBitsPerSecond: jsii.Number(123),
//   					StreamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		Encryption: &MssEncryptionProperty{
//   			SpekeKeyProvider: &SpekeKeyProviderProperty{
//   				RoleArn: jsii.String("roleArn"),
//   				SystemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				Url: jsii.String("url"),
//
//   				// the properties below are optional
//   				EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				},
//   			},
//   		},
//   		SegmentDurationSeconds: jsii.Number(123),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPackagingConfigurationProps struct {
	// Unique identifier that you assign to the packaging configuration.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The ID of the packaging group associated with this packaging configuration.
	PackagingGroupId *string `field:"required" json:"packagingGroupId" yaml:"packagingGroupId"`
	// Parameters for CMAF packaging.
	CmafPackage interface{} `field:"optional" json:"cmafPackage" yaml:"cmafPackage"`
	// Parameters for DASH-ISO packaging.
	DashPackage interface{} `field:"optional" json:"dashPackage" yaml:"dashPackage"`
	// Parameters for Apple HLS packaging.
	HlsPackage interface{} `field:"optional" json:"hlsPackage" yaml:"hlsPackage"`
	// Parameters for Microsoft Smooth Streaming packaging.
	MssPackage interface{} `field:"optional" json:"mssPackage" yaml:"mssPackage"`
	// The tags to assign to the packaging configuration.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


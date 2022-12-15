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
//   cfnPackagingConfigurationProps := &cfnPackagingConfigurationProps{
//   	id: jsii.String("id"),
//   	packagingGroupId: jsii.String("packagingGroupId"),
//
//   	// the properties below are optional
//   	cmafPackage: &cmafPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &cmafEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	dashPackage: &dashPackageProperty{
//   		dashManifests: []interface{}{
//   			&dashManifestProperty{
//   				manifestLayout: jsii.String("manifestLayout"),
//   				manifestName: jsii.String("manifestName"),
//   				minBufferTimeSeconds: jsii.Number(123),
//   				profile: jsii.String("profile"),
//   				scteMarkersSource: jsii.String("scteMarkersSource"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &dashEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		includeEncoderConfigurationInSegments: jsii.Boolean(false),
//   		periodTriggers: []*string{
//   			jsii.String("periodTriggers"),
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   		segmentTemplateFormat: jsii.String("segmentTemplateFormat"),
//   	},
//   	hlsPackage: &hlsPackageProperty{
//   		hlsManifests: []interface{}{
//   			&hlsManifestProperty{
//   				adMarkers: jsii.String("adMarkers"),
//   				includeIframeOnlyStream: jsii.Boolean(false),
//   				manifestName: jsii.String("manifestName"),
//   				programDateTimeIntervalSeconds: jsii.Number(123),
//   				repeatExtXKey: jsii.Boolean(false),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &hlsEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			constantInitializationVector: jsii.String("constantInitializationVector"),
//   			encryptionMethod: jsii.String("encryptionMethod"),
//   		},
//   		includeDvbSubtitles: jsii.Boolean(false),
//   		segmentDurationSeconds: jsii.Number(123),
//   		useAudioRenditionGroup: jsii.Boolean(false),
//   	},
//   	mssPackage: &mssPackageProperty{
//   		mssManifests: []interface{}{
//   			&mssManifestProperty{
//   				manifestName: jsii.String("manifestName"),
//   				streamSelection: &streamSelectionProperty{
//   					maxVideoBitsPerSecond: jsii.Number(123),
//   					minVideoBitsPerSecond: jsii.Number(123),
//   					streamOrder: jsii.String("streamOrder"),
//   				},
//   			},
//   		},
//
//   		// the properties below are optional
//   		encryption: &mssEncryptionProperty{
//   			spekeKeyProvider: &spekeKeyProviderProperty{
//   				roleArn: jsii.String("roleArn"),
//   				systemIds: []*string{
//   					jsii.String("systemIds"),
//   				},
//   				url: jsii.String("url"),
//
//   				// the properties below are optional
//   				encryptionContractConfiguration: &encryptionContractConfigurationProperty{
//   					presetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   					presetSpeke20Video: jsii.String("presetSpeke20Video"),
//   				},
//   			},
//   		},
//   		segmentDurationSeconds: jsii.Number(123),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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


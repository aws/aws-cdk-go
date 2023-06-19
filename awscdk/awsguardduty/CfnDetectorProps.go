package awsguardduty

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDetector`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDetectorProps := &CfnDetectorProps{
//   	Enable: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	DataSources: &CFNDataSourceConfigurationsProperty{
//   		Kubernetes: &CFNKubernetesConfigurationProperty{
//   			AuditLogs: &CFNKubernetesAuditLogsConfigurationProperty{
//   				Enable: jsii.Boolean(false),
//   			},
//   		},
//   		MalwareProtection: &CFNMalwareProtectionConfigurationProperty{
//   			ScanEc2InstanceWithFindings: &CFNScanEc2InstanceWithFindingsConfigurationProperty{
//   				EbsVolumes: jsii.Boolean(false),
//   			},
//   		},
//   		S3Logs: &CFNS3LogsConfigurationProperty{
//   			Enable: jsii.Boolean(false),
//   		},
//   	},
//   	Features: []interface{}{
//   		&FeatureConfigurationsProperty{
//   			AdditionalConfiguration: []interface{}{
//   				&FeatureAdditionalConfigurationProperty{
//   					Name: jsii.String("name"),
//   					Status: jsii.String("status"),
//   				},
//   			},
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//   		},
//   	},
//   	FindingPublishingFrequency: jsii.String("findingPublishingFrequency"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetectorProps struct {
	// Specifies whether the detector is to be enabled on creation.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Describes which data sources will be enabled for the detector.
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// A list of features that will be configured for the detector.
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// Specifies how frequently updated findings are exported.
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// Specifies tags added to a new detector resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// Currently, support is available only for creating and deleting a tag. No support exists for updating the tags.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


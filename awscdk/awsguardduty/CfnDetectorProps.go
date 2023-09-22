package awsguardduty


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
//   		&CFNFeatureConfigurationProperty{
//   			Name: jsii.String("name"),
//   			Status: jsii.String("status"),
//
//   			// the properties below are optional
//   			AdditionalConfiguration: []interface{}{
//   				&CFNFeatureAdditionalConfigurationProperty{
//   					Name: jsii.String("name"),
//   					Status: jsii.String("status"),
//   				},
//   			},
//   		},
//   	},
//   	FindingPublishingFrequency: jsii.String("findingPublishingFrequency"),
//   	Tags: []tagItemProperty{
//   		&tagItemProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html
//
type CfnDetectorProps struct {
	// Specifies whether the detector is to be enabled on creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-enable
	//
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Describes which data sources will be enabled for the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-datasources
	//
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// A list of features that will be configured for the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-features
	//
	Features interface{} `field:"optional" json:"features" yaml:"features"`
	// Specifies how frequently updated findings are exported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-findingpublishingfrequency
	//
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// Specifies tags added to a new detector resource.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	//
	// Currently, support is available only for creating and deleting a tag. No support exists for updating the tags.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-tags
	//
	Tags *[]*CfnDetector_TagItemProperty `field:"optional" json:"tags" yaml:"tags"`
}


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
//   cfnDetectorProps := &cfnDetectorProps{
//   	enable: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	dataSources: &cFNDataSourceConfigurationsProperty{
//   		kubernetes: &cFNKubernetesConfigurationProperty{
//   			auditLogs: &cFNKubernetesAuditLogsConfigurationProperty{
//   				enable: jsii.Boolean(false),
//   			},
//   		},
//   		malwareProtection: &cFNMalwareProtectionConfigurationProperty{
//   			scanEc2InstanceWithFindings: &cFNScanEc2InstanceWithFindingsConfigurationProperty{
//   				ebsVolumes: jsii.Boolean(false),
//   			},
//   		},
//   		s3Logs: &cFNS3LogsConfigurationProperty{
//   			enable: jsii.Boolean(false),
//   		},
//   	},
//   	findingPublishingFrequency: jsii.String("findingPublishingFrequency"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnDetectorProps struct {
	// Specifies whether the detector is to be enabled on creation.
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
	// Describes which data sources will be enabled for the detector.
	DataSources interface{} `field:"optional" json:"dataSources" yaml:"dataSources"`
	// Specifies how frequently updated findings are exported.
	FindingPublishingFrequency *string `field:"optional" json:"findingPublishingFrequency" yaml:"findingPublishingFrequency"`
	// `AWS::GuardDuty::Detector.Tags`.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


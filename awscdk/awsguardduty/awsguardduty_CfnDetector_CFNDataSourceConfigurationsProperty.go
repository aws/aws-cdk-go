package awsguardduty


// Describes whether S3 data event logs or Kubernetes audit logs will be enabled as a data source when the detector is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNDataSourceConfigurationsProperty := &cFNDataSourceConfigurationsProperty{
//   	kubernetes: &cFNKubernetesConfigurationProperty{
//   		auditLogs: &cFNKubernetesAuditLogsConfigurationProperty{
//   			enable: jsii.Boolean(false),
//   		},
//   	},
//   	malwareProtection: &cFNMalwareProtectionConfigurationProperty{
//   		scanEc2InstanceWithFindings: &cFNScanEc2InstanceWithFindingsConfigurationProperty{
//   			ebsVolumes: jsii.Boolean(false),
//   		},
//   	},
//   	s3Logs: &cFNS3LogsConfigurationProperty{
//   		enable: jsii.Boolean(false),
//   	},
//   }
//
type CfnDetector_CFNDataSourceConfigurationsProperty struct {
	// Describes which Kuberentes data sources are enabled for a detector.
	Kubernetes interface{} `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// `CfnDetector.CFNDataSourceConfigurationsProperty.MalwareProtection`.
	MalwareProtection interface{} `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// Describes whether S3 data event logs are enabled as a data source.
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}


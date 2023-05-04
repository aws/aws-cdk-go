package awsguardduty


// Describes whether S3 data event logs, Kubernetes audit logs, or Malware Protection will be enabled as a data source when the detector is created.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNDataSourceConfigurationsProperty := &CFNDataSourceConfigurationsProperty{
//   	Kubernetes: &CFNKubernetesConfigurationProperty{
//   		AuditLogs: &CFNKubernetesAuditLogsConfigurationProperty{
//   			Enable: jsii.Boolean(false),
//   		},
//   	},
//   	MalwareProtection: &CFNMalwareProtectionConfigurationProperty{
//   		ScanEc2InstanceWithFindings: &CFNScanEc2InstanceWithFindingsConfigurationProperty{
//   			EbsVolumes: jsii.Boolean(false),
//   		},
//   	},
//   	S3Logs: &CFNS3LogsConfigurationProperty{
//   		Enable: jsii.Boolean(false),
//   	},
//   }
//
type CfnDetector_CFNDataSourceConfigurationsProperty struct {
	// Describes which Kuberentes data sources are enabled for a detector.
	Kubernetes interface{} `field:"optional" json:"kubernetes" yaml:"kubernetes"`
	// Describes whether Malware Protection will be enabled as a data source.
	MalwareProtection interface{} `field:"optional" json:"malwareProtection" yaml:"malwareProtection"`
	// Describes whether S3 data event logs are enabled as a data source.
	S3Logs interface{} `field:"optional" json:"s3Logs" yaml:"s3Logs"`
}


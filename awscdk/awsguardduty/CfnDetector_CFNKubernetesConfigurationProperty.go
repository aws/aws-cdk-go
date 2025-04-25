package awsguardduty


// Describes which Kubernetes protection data sources are enabled for the detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNKubernetesConfigurationProperty := &CFNKubernetesConfigurationProperty{
//   	AuditLogs: &CFNKubernetesAuditLogsConfigurationProperty{
//   		Enable: jsii.Boolean(false),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnkubernetesconfiguration.html
//
type CfnDetector_CFNKubernetesConfigurationProperty struct {
	// Describes whether Kubernetes audit logs are enabled as a data source for the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnkubernetesconfiguration.html#cfn-guardduty-detector-cfnkubernetesconfiguration-auditlogs
	//
	AuditLogs interface{} `field:"required" json:"auditLogs" yaml:"auditLogs"`
}


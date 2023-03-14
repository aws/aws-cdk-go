package awsguardduty


// Describes which Kubernetes protection data sources are enabled for the detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNKubernetesConfigurationProperty := &cFNKubernetesConfigurationProperty{
//   	auditLogs: &cFNKubernetesAuditLogsConfigurationProperty{
//   		enable: jsii.Boolean(false),
//   	},
//   }
//
type CfnDetector_CFNKubernetesConfigurationProperty struct {
	// Describes whether Kubernetes audit logs are enabled as a data source for the detector.
	AuditLogs interface{} `field:"optional" json:"auditLogs" yaml:"auditLogs"`
}


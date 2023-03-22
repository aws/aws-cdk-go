package awsguardduty


// Describes which optional data sources are enabled for a detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNKubernetesAuditLogsConfigurationProperty := &cFNKubernetesAuditLogsConfigurationProperty{
//   	enable: jsii.Boolean(false),
//   }
//
type CfnDetector_CFNKubernetesAuditLogsConfigurationProperty struct {
	// Describes whether Kubernetes audit logs are enabled as a data source for the detector.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
}


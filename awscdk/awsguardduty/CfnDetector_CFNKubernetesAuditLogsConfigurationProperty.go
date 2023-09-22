package awsguardduty


// Describes which optional data sources are enabled for a detector.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cFNKubernetesAuditLogsConfigurationProperty := &CFNKubernetesAuditLogsConfigurationProperty{
//   	Enable: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnkubernetesauditlogsconfiguration.html
//
type CfnDetector_CFNKubernetesAuditLogsConfigurationProperty struct {
	// Describes whether Kubernetes audit logs are enabled as a data source for the detector.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-detector-cfnkubernetesauditlogsconfiguration.html#cfn-guardduty-detector-cfnkubernetesauditlogsconfiguration-enable
	//
	Enable interface{} `field:"required" json:"enable" yaml:"enable"`
}


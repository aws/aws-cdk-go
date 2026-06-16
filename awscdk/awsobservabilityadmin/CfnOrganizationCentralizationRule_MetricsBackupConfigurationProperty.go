package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricsBackupConfigurationProperty := &MetricsBackupConfigurationProperty{
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-metricsbackupconfiguration.html
//
type CfnOrganizationCentralizationRule_MetricsBackupConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-metricsbackupconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-metricsbackupconfiguration-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
}


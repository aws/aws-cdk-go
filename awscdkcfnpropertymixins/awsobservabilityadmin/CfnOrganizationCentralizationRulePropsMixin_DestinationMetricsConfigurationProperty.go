package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   destinationMetricsConfigurationProperty := &DestinationMetricsConfigurationProperty{
//   	BackupConfiguration: &MetricsBackupConfigurationProperty{
//   		Region: jsii.String("region"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationmetricsconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_DestinationMetricsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationmetricsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-destinationmetricsconfiguration-backupconfiguration
	//
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
}


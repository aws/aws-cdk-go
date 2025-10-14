package awsobservabilityadmin


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   destinationLogsConfigurationProperty := &DestinationLogsConfigurationProperty{
//   	BackupConfiguration: &LogsBackupConfigurationProperty{
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   	LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   		EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   		// the properties below are optional
//   		EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html
//
type CfnOrganizationCentralizationRule_DestinationLogsConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration-backupconfiguration
	//
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration-logsencryptionconfiguration
	//
	LogsEncryptionConfiguration interface{} `field:"optional" json:"logsEncryptionConfiguration" yaml:"logsEncryptionConfiguration"`
}


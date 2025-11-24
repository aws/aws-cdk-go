package mixinsawsobservabilityadmin


// Configuration for centralization destination log groups, including encryption and backup settings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   destinationLogsConfigurationProperty := &DestinationLogsConfigurationProperty{
//   	BackupConfiguration: &LogsBackupConfigurationProperty{
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   		Region: jsii.String("region"),
//   	},
//   	LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   		EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   		EncryptionStrategy: jsii.String("encryptionStrategy"),
//   		KmsKeyArn: jsii.String("kmsKeyArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_DestinationLogsConfigurationProperty struct {
	// Configuration defining the backup region and an optional KMS key for the backup destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration-backupconfiguration
	//
	BackupConfiguration interface{} `field:"optional" json:"backupConfiguration" yaml:"backupConfiguration"`
	// The encryption configuration for centralization destination log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-destinationlogsconfiguration-logsencryptionconfiguration
	//
	LogsEncryptionConfiguration interface{} `field:"optional" json:"logsEncryptionConfiguration" yaml:"logsEncryptionConfiguration"`
}


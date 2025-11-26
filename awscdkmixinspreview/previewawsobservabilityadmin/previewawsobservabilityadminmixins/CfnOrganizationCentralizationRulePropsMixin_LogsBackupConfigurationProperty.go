package previewawsobservabilityadminmixins


// Configuration for backing up centralized log data to a secondary region.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logsBackupConfigurationProperty := &LogsBackupConfigurationProperty{
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration.html
//
type CfnOrganizationCentralizationRulePropsMixin_LogsBackupConfigurationProperty struct {
	// KMS Key ARN belonging to the primary destination account and backup region, to encrypt newly created central log groups in the backup destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Logs specific backup destination region within the primary destination account to which log data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration.html#cfn-observabilityadmin-organizationcentralizationrule-logsbackupconfiguration-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}


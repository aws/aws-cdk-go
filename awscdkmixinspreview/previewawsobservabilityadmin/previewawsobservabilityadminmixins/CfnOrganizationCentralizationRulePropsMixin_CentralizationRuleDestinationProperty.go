package previewawsobservabilityadminmixins


// Configuration specifying the primary destination for centralized telemetry data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   centralizationRuleDestinationProperty := &CentralizationRuleDestinationProperty{
//   	Account: jsii.String("account"),
//   	DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   		BackupConfiguration: &LogsBackupConfigurationProperty{
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   			Region: jsii.String("region"),
//   		},
//   		LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   			EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   			EncryptionStrategy: jsii.String("encryptionStrategy"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   	Region: jsii.String("region"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html
//
type CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleDestinationProperty struct {
	// The destination account (within the organization) to which the telemetry data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Log specific configuration for centralization destination log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-destinationlogsconfiguration
	//
	DestinationLogsConfiguration interface{} `field:"optional" json:"destinationLogsConfiguration" yaml:"destinationLogsConfiguration"`
	// The primary destination region to which telemetry data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
}


package awsobservabilityadmin


// Configuration specifying the primary destination for centralized telemetry data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   centralizationRuleDestinationProperty := &CentralizationRuleDestinationProperty{
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	Account: jsii.String("account"),
//   	DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   		BackupConfiguration: &LogsBackupConfigurationProperty{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   		LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   			EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   			// the properties below are optional
//   			EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   			KmsKeyArn: jsii.String("kmsKeyArn"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html
//
type CfnOrganizationCentralizationRule_CentralizationRuleDestinationProperty struct {
	// The primary destination region to which telemetry data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// The destination account (within the organization) to which the telemetry data should be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-account
	//
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Log specific configuration for centralization destination log groups.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationruledestination.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationruledestination-destinationlogsconfiguration
	//
	DestinationLogsConfiguration interface{} `field:"optional" json:"destinationLogsConfiguration" yaml:"destinationLogsConfiguration"`
}


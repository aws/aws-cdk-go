package awsobservabilityadmin


// Defines how telemetry data should be centralized across an AWS Organization, including source and destination configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   centralizationRuleProperty := &CentralizationRuleProperty{
//   	Destination: &CentralizationRuleDestinationProperty{
//   		Region: jsii.String("region"),
//
//   		// the properties below are optional
//   		Account: jsii.String("account"),
//   		DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   			BackupConfiguration: &LogsBackupConfigurationProperty{
//   				Region: jsii.String("region"),
//
//   				// the properties below are optional
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   			LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   				EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   				// the properties below are optional
//   				EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   	},
//   	Source: &CentralizationRuleSourceProperty{
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
//
//   		// the properties below are optional
//   		Scope: jsii.String("scope"),
//   		SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   			EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   			LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html
//
type CfnOrganizationCentralizationRule_CentralizationRuleProperty struct {
	// Configuration determining where the telemetry data should be centralized, backed up, as well as encryption configuration for the primary and backup destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrule-destination
	//
	Destination interface{} `field:"required" json:"destination" yaml:"destination"`
	// Configuration determining the source of the telemetry data to be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrule-source
	//
	Source interface{} `field:"required" json:"source" yaml:"source"`
}


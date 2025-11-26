package previewawsobservabilityadminmixins


// Defines how telemetry data should be centralized across an AWS Organization, including source and destination configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   centralizationRuleProperty := &CentralizationRuleProperty{
//   	Destination: &CentralizationRuleDestinationProperty{
//   		Account: jsii.String("account"),
//   		DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   			BackupConfiguration: &LogsBackupConfigurationProperty{
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   				Region: jsii.String("region"),
//   			},
//   			LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   				EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   				EncryptionStrategy: jsii.String("encryptionStrategy"),
//   				KmsKeyArn: jsii.String("kmsKeyArn"),
//   			},
//   		},
//   		Region: jsii.String("region"),
//   	},
//   	Source: &CentralizationRuleSourceProperty{
//   		Regions: []*string{
//   			jsii.String("regions"),
//   		},
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
type CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleProperty struct {
	// Configuration determining where the telemetry data should be centralized, backed up, as well as encryption configuration for the primary and backup destinations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrule-destination
	//
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Configuration determining the source of the telemetry data to be centralized.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-observabilityadmin-organizationcentralizationrule-centralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-centralizationrule-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
}


package awsobservabilityadmin

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOrganizationCentralizationRule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOrganizationCentralizationRuleProps := &CfnOrganizationCentralizationRuleProps{
//   	Rule: &CentralizationRuleProperty{
//   		Destination: &CentralizationRuleDestinationProperty{
//   			Region: jsii.String("region"),
//
//   			// the properties below are optional
//   			Account: jsii.String("account"),
//   			DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   				BackupConfiguration: &LogsBackupConfigurationProperty{
//   					Region: jsii.String("region"),
//
//   					// the properties below are optional
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   				LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   					EncryptionStrategy: jsii.String("encryptionStrategy"),
//
//   					// the properties below are optional
//   					EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   		},
//   		Source: &CentralizationRuleSourceProperty{
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//
//   			// the properties below are optional
//   			Scope: jsii.String("scope"),
//   			SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   				EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   				LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html
//
type CfnOrganizationCentralizationRuleProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-rule
	//
	Rule interface{} `field:"required" json:"rule" yaml:"rule"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-rulename
	//
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


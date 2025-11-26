package previewawsobservabilityadminmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOrganizationCentralizationRulePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOrganizationCentralizationRuleMixinProps := &CfnOrganizationCentralizationRuleMixinProps{
//   	Rule: &CentralizationRuleProperty{
//   		Destination: &CentralizationRuleDestinationProperty{
//   			Account: jsii.String("account"),
//   			DestinationLogsConfiguration: &DestinationLogsConfigurationProperty{
//   				BackupConfiguration: &LogsBackupConfigurationProperty{
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   					Region: jsii.String("region"),
//   				},
//   				LogsEncryptionConfiguration: &LogsEncryptionConfigurationProperty{
//   					EncryptionConflictResolutionStrategy: jsii.String("encryptionConflictResolutionStrategy"),
//   					EncryptionStrategy: jsii.String("encryptionStrategy"),
//   					KmsKeyArn: jsii.String("kmsKeyArn"),
//   				},
//   			},
//   			Region: jsii.String("region"),
//   		},
//   		Source: &CentralizationRuleSourceProperty{
//   			Regions: []*string{
//   				jsii.String("regions"),
//   			},
//   			Scope: jsii.String("scope"),
//   			SourceLogsConfiguration: &SourceLogsConfigurationProperty{
//   				EncryptedLogGroupStrategy: jsii.String("encryptedLogGroupStrategy"),
//   				LogGroupSelectionCriteria: jsii.String("logGroupSelectionCriteria"),
//   			},
//   		},
//   	},
//   	RuleName: jsii.String("ruleName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html
//
type CfnOrganizationCentralizationRuleMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-rule
	//
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// The name of the organization centralization rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-rulename
	//
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// A key-value pair to filter resources based on tags associated with the resource.
	//
	// For more information about tags, see [What are tags?](https://docs.aws.amazon.com/whitepapers/latest/tagging-best-practices/what-are-tags.html)
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-observabilityadmin-organizationcentralizationrule.html#cfn-observabilityadmin-organizationcentralizationrule-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


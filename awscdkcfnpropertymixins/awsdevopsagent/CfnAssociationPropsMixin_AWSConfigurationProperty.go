package awsdevopsagent


// Configuration for AWS monitor account integration.
//
// Specifies the account ID, assumable role ARN, and resources to be monitored in the primary monitoring account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var resourceMetadata interface{}
//
//   aWSConfigurationProperty := &AWSConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	AccountType: jsii.String("accountType"),
//   	AgentElevatedRoleArn: jsii.String("agentElevatedRoleArn"),
//   	AssumableRoleArn: jsii.String("assumableRoleArn"),
//   	Resources: []interface{}{
//   		&AWSResourceProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//   			ResourceMetadata: resourceMetadata,
//   			ResourceType: jsii.String("resourceType"),
//   		},
//   	},
//   	Tags: []KeyValuePairProperty{
//   		&KeyValuePairProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html
//
type CfnAssociationPropsMixin_AWSConfigurationProperty struct {
	// Account ID corresponding to the provided resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Account Type 'monitor' for AWS DevOps Agent monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accounttype
	//
	AccountType *string `field:"optional" json:"accountType" yaml:"accountType"`
	// The IAM role that AWS DevOps Agent assumes to perform elevated directed actions in this AWS account.
	//
	// Elevated actions also require ElevatedActionsEnabled to be set in the Preferences of the AgentSpace. When not provided, only non-elevated directed actions are available. Removing this property from the template clears the role on update. Setting this role is subject to the same iam:PassRole requirement as AssumableRoleArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-agentelevatedrolearn
	//
	AgentElevatedRoleArn *string `field:"optional" json:"agentElevatedRoleArn" yaml:"agentElevatedRoleArn"`
	// Role ARN used by AWS DevOps Agent to access resources in the primary account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-assumablerolearn
	//
	AssumableRoleArn *string `field:"optional" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// List of resources to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of tags as key-value pairs, used to identify resources for topology crawl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-tags
	//
	Tags *[]*CfnAssociationPropsMixin_KeyValuePairProperty `field:"optional" json:"tags" yaml:"tags"`
}


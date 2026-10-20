package awsdevopsagent


// Configuration for AWS monitor account integration.
//
// Specifies the account ID, assumable role ARN, and resources to be monitored in the primary monitoring account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var resourceMetadata interface{}
//
//   aWSConfigurationProperty := &AWSConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	AccountType: jsii.String("accountType"),
//   	AssumableRoleArn: jsii.String("assumableRoleArn"),
//
//   	// the properties below are optional
//   	AgentElevatedRoleArn: jsii.String("agentElevatedRoleArn"),
//   	Resources: []interface{}{
//   		&AWSResourceProperty{
//   			ResourceArn: jsii.String("resourceArn"),
//
//   			// the properties below are optional
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
type CfnAssociation_AWSConfigurationProperty struct {
	// Account ID corresponding to the provided resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Account Type 'monitor' for AWS DevOps Agent monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accounttype
	//
	AccountType *string `field:"required" json:"accountType" yaml:"accountType"`
	// Role ARN used by AWS DevOps Agent to access resources in the primary account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-assumablerolearn
	//
	AssumableRoleArn *string `field:"required" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// The IAM role that AWS DevOps Agent assumes to perform elevated directed actions in this AWS account.
	//
	// Elevated actions also require ElevatedActionsEnabled to be set in the Preferences of the AgentSpace. When not provided, only non-elevated directed actions are available. Removing this property from the template clears the role on update. Setting this role is subject to the same iam:PassRole requirement as AssumableRoleArn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-agentelevatedrolearn
	//
	AgentElevatedRoleArn *string `field:"optional" json:"agentElevatedRoleArn" yaml:"agentElevatedRoleArn"`
	// List of resources to monitor.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of tags as key-value pairs, used to identify resources for topology crawl.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-tags
	//
	Tags *[]*CfnAssociation_KeyValuePairProperty `field:"optional" json:"tags" yaml:"tags"`
}


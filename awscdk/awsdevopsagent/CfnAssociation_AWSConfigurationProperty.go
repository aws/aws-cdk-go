package awsdevopsagent


// AWS association for 'monitor' account.
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
	// AWS Account Id corresponding to provided resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// Account Type 'monitor' for DevOpsAgent monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-accounttype
	//
	AccountType *string `field:"required" json:"accountType" yaml:"accountType"`
	// Role ARN to be assumed by DevOpsAgent to operate on behalf of customer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-assumablerolearn
	//
	AssumableRoleArn *string `field:"required" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// List of AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of AWS tags as key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-awsconfiguration.html#cfn-devopsagent-association-awsconfiguration-tags
	//
	Tags *[]*CfnAssociation_KeyValuePairProperty `field:"optional" json:"tags" yaml:"tags"`
}


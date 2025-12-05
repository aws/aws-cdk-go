package previewawsdevopsagentmixins


// AWS association for 'source' account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var resourceMetadata interface{}
//
//   sourceAwsConfigurationProperty := &SourceAwsConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	AccountType: jsii.String("accountType"),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html
//
type CfnAssociationPropsMixin_SourceAwsConfigurationProperty struct {
	// AWS Account Id corresponding to provided resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html#cfn-devopsagent-association-sourceawsconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Account Type 'source' for DevOpsAgent monitoring.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html#cfn-devopsagent-association-sourceawsconfiguration-accounttype
	//
	AccountType *string `field:"optional" json:"accountType" yaml:"accountType"`
	// Role ARN to be assumed by DevOpsAgent to operate on behalf of customer.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html#cfn-devopsagent-association-sourceawsconfiguration-assumablerolearn
	//
	AssumableRoleArn *string `field:"optional" json:"assumableRoleArn" yaml:"assumableRoleArn"`
	// List of AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html#cfn-devopsagent-association-sourceawsconfiguration-resources
	//
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// List of AWS tags as key-value pairs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-association-sourceawsconfiguration.html#cfn-devopsagent-association-sourceawsconfiguration-tags
	//
	Tags *[]*CfnAssociationPropsMixin_KeyValuePairProperty `field:"optional" json:"tags" yaml:"tags"`
}


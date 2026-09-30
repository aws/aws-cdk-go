package awswellarchitected


// Defines an account, its in-scope regions, and the access role used to reach resources in that account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   aggregationConfigurationProperty := &AggregationConfigurationProperty{
//   	AccessRoleArn: jsii.String("accessRoleArn"),
//   	AccountId: jsii.String("accountId"),
//   	Regions: []*string{
//   		jsii.String("regions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-agentprofile-aggregationconfiguration.html
//
type CfnAgentProfilePropsMixin_AggregationConfigurationProperty struct {
	// The ARN of the IAM role used to access resources in this account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-agentprofile-aggregationconfiguration.html#cfn-wellarchitected-agentprofile-aggregationconfiguration-accessrolearn
	//
	AccessRoleArn *string `field:"optional" json:"accessRoleArn" yaml:"accessRoleArn"`
	// The target AWS account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-agentprofile-aggregationconfiguration.html#cfn-wellarchitected-agentprofile-aggregationconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// The target regions in the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wellarchitected-agentprofile-aggregationconfiguration.html#cfn-wellarchitected-agentprofile-aggregationconfiguration-regions
	//
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


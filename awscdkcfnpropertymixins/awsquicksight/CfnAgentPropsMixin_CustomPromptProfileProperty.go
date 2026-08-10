package awsquicksight


// Reference to an existing custom prompt profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customPromptProfileProperty := &CustomPromptProfileProperty{
//   	ModelProfileId: jsii.String("modelProfileId"),
//   	QbsAwsAccountId: jsii.String("qbsAwsAccountId"),
//   	SubscriptionId: jsii.String("subscriptionId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptprofile.html
//
type CfnAgentPropsMixin_CustomPromptProfileProperty struct {
	// The identifier of the model profile.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptprofile.html#cfn-quicksight-agent-custompromptprofile-modelprofileid
	//
	ModelProfileId *string `field:"optional" json:"modelProfileId" yaml:"modelProfileId"`
	// The QBS AWS account identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptprofile.html#cfn-quicksight-agent-custompromptprofile-qbsawsaccountid
	//
	QbsAwsAccountId *string `field:"optional" json:"qbsAwsAccountId" yaml:"qbsAwsAccountId"`
	// The subscription identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-agent-custompromptprofile.html#cfn-quicksight-agent-custompromptprofile-subscriptionid
	//
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
}


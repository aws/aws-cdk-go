package awswisdom


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userInteractionConfigurationProperty := &UserInteractionConfigurationProperty{
//   	IsUserConfirmationRequired: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-userinteractionconfiguration.html
//
type CfnAIAgent_UserInteractionConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-aiagent-userinteractionconfiguration.html#cfn-wisdom-aiagent-userinteractionconfiguration-isuserconfirmationrequired
	//
	IsUserConfirmationRequired interface{} `field:"optional" json:"isUserConfirmationRequired" yaml:"isUserConfirmationRequired"`
}


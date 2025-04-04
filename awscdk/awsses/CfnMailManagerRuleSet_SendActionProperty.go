package awsses


// Sends the email to the internet using the ses:SendRawEmail API.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sendActionProperty := &SendActionProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-sendaction.html
//
type CfnMailManagerRuleSet_SendActionProperty struct {
	// The Amazon Resource Name (ARN) of the role to use for this action.
	//
	// This role must have access to the ses:SendRawEmail API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-sendaction.html#cfn-ses-mailmanagerruleset-sendaction-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the caller does not have the permissions to call the sendRawEmail API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-sendaction.html#cfn-ses-mailmanagerruleset-sendaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
}


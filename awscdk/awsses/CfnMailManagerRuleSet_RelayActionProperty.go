package awsses


// The action relays the email via SMTP to another specific SMTP server.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   relayActionProperty := &RelayActionProperty{
//   	Relay: jsii.String("relay"),
//
//   	// the properties below are optional
//   	ActionFailurePolicy: jsii.String("actionFailurePolicy"),
//   	MailFrom: jsii.String("mailFrom"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-relayaction.html
//
type CfnMailManagerRuleSet_RelayActionProperty struct {
	// The identifier of the relay resource to be used when relaying an email.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-relayaction.html#cfn-ses-mailmanagerruleset-relayaction-relay
	//
	Relay *string `field:"required" json:"relay" yaml:"relay"`
	// A policy that states what to do in the case of failure.
	//
	// The action will fail if there are configuration errors. For example, the specified relay has been deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-relayaction.html#cfn-ses-mailmanagerruleset-relayaction-actionfailurepolicy
	//
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// This action specifies whether to preserve or replace original mail from address while relaying received emails to a destination server.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-mailmanagerruleset-relayaction.html#cfn-ses-mailmanagerruleset-relayaction-mailfrom
	//
	MailFrom *string `field:"optional" json:"mailFrom" yaml:"mailFrom"`
}


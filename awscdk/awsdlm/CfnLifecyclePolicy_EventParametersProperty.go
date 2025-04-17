package awsdlm


// *[Event-based policies only]* Specifies an event that activates an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventParametersProperty := &EventParametersProperty{
//   	EventType: jsii.String("eventType"),
//   	SnapshotOwner: []*string{
//   		jsii.String("snapshotOwner"),
//   	},
//
//   	// the properties below are optional
//   	DescriptionRegex: jsii.String("descriptionRegex"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventparameters.html
//
type CfnLifecyclePolicy_EventParametersProperty struct {
	// The type of event.
	//
	// Currently, only snapshot sharing events are supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventparameters.html#cfn-dlm-lifecyclepolicy-eventparameters-eventtype
	//
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
	//
	// The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventparameters.html#cfn-dlm-lifecyclepolicy-eventparameters-snapshotowner
	//
	SnapshotOwner *[]*string `field:"required" json:"snapshotOwner" yaml:"snapshotOwner"`
	// The snapshot description that can trigger the policy.
	//
	// The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	//
	// For example, specifying `^.*Created for policy: policy-1234567890abcdef0.*$` configures the policy to run only if snapshots created by policy `policy-1234567890abcdef0` are shared with your account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dlm-lifecyclepolicy-eventparameters.html#cfn-dlm-lifecyclepolicy-eventparameters-descriptionregex
	//
	DescriptionRegex *string `field:"optional" json:"descriptionRegex" yaml:"descriptionRegex"`
}


package awsdlm


// Specifies an event that triggers an event-based policy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventParametersProperty := &eventParametersProperty{
//   	eventType: jsii.String("eventType"),
//   	snapshotOwner: []*string{
//   		jsii.String("snapshotOwner"),
//   	},
//
//   	// the properties below are optional
//   	descriptionRegex: jsii.String("descriptionRegex"),
//   }
//
type CfnLifecyclePolicy_EventParametersProperty struct {
	// The type of event.
	//
	// Currently, only snapshot sharing events are supported.
	EventType *string `field:"required" json:"eventType" yaml:"eventType"`
	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
	//
	// The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	SnapshotOwner *[]*string `field:"required" json:"snapshotOwner" yaml:"snapshotOwner"`
	// The snapshot description that can trigger the policy.
	//
	// The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	//
	// For example, specifying `^.*Created for policy: policy-1234567890abcdef0.*$` configures the policy to run only if snapshots created by policy `policy-1234567890abcdef0` are shared with your account.
	DescriptionRegex *string `field:"optional" json:"descriptionRegex" yaml:"descriptionRegex"`
}


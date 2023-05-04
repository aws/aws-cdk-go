package awsssmcontacts


// The contact or contact channel that's being engaged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &TargetsProperty{
//   	ChannelTargetInfo: &ChannelTargetInfoProperty{
//   		ChannelId: jsii.String("channelId"),
//   		RetryIntervalInMinutes: jsii.Number(123),
//   	},
//   	ContactTargetInfo: &ContactTargetInfoProperty{
//   		ContactId: jsii.String("contactId"),
//   		IsEssential: jsii.Boolean(false),
//   	},
//   }
//
type CfnContact_TargetsProperty struct {
	// Information about the contact channel that Incident Manager engages.
	ChannelTargetInfo interface{} `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// The contact that Incident Manager is engaging during an incident.
	ContactTargetInfo interface{} `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}


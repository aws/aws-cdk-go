package awsssmcontacts


// The contact or contact channel that's being engaged.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetsProperty := &targetsProperty{
//   	channelTargetInfo: &channelTargetInfoProperty{
//   		channelId: jsii.String("channelId"),
//   		retryIntervalInMinutes: jsii.Number(123),
//   	},
//   	contactTargetInfo: &contactTargetInfoProperty{
//   		contactId: jsii.String("contactId"),
//   		isEssential: jsii.Boolean(false),
//   	},
//   }
//
type CfnContact_TargetsProperty struct {
	// Information about the contact channel Incident Manager is engaging.
	ChannelTargetInfo interface{} `field:"optional" json:"channelTargetInfo" yaml:"channelTargetInfo"`
	// The contact that Incident Manager is engaging during an incident.
	ContactTargetInfo interface{} `field:"optional" json:"contactTargetInfo" yaml:"contactTargetInfo"`
}


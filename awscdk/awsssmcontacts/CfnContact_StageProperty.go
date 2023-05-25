package awsssmcontacts


// The `Stage` property type specifies a set amount of time that an escalation plan or engagement plan engages the specified contacts or contact methods.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stageProperty := &StageProperty{
//   	DurationInMinutes: jsii.Number(123),
//   	RotationIds: []*string{
//   		jsii.String("rotationIds"),
//   	},
//   	Targets: []interface{}{
//   		&TargetsProperty{
//   			ChannelTargetInfo: &ChannelTargetInfoProperty{
//   				ChannelId: jsii.String("channelId"),
//   				RetryIntervalInMinutes: jsii.Number(123),
//   			},
//   			ContactTargetInfo: &ContactTargetInfoProperty{
//   				ContactId: jsii.String("contactId"),
//   				IsEssential: jsii.Boolean(false),
//   			},
//   		},
//   	},
//   }
//
type CfnContact_StageProperty struct {
	// The time to wait until beginning the next stage.
	//
	// The duration can only be set to 0 if a target is specified.
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// `CfnContact.StageProperty.RotationIds`.
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// The contacts or contact methods that the escalation plan or engagement plan is engaging.
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}


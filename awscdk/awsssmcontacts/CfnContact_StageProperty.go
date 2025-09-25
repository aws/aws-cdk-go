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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-stage.html
//
type CfnContact_StageProperty struct {
	// The time to wait until beginning the next stage.
	//
	// The duration can only be set to 0 if a target is specified.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-stage.html#cfn-ssmcontacts-contact-stage-durationinminutes
	//
	DurationInMinutes *float64 `field:"optional" json:"durationInMinutes" yaml:"durationInMinutes"`
	// The Amazon Resource Names (ARNs) of the on-call rotations associated with the plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-stage.html#cfn-ssmcontacts-contact-stage-rotationids
	//
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// The contacts or contact methods that the escalation plan or engagement plan is engaging.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssmcontacts-contact-stage.html#cfn-ssmcontacts-contact-stage-targets
	//
	Targets interface{} `field:"optional" json:"targets" yaml:"targets"`
}


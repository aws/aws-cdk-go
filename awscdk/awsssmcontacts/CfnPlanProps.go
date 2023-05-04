package awsssmcontacts


// Properties for defining a `CfnPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPlanProps := &CfnPlanProps{
//   	ContactId: jsii.String("contactId"),
//
//   	// the properties below are optional
//   	RotationIds: []*string{
//   		jsii.String("rotationIds"),
//   	},
//   	Stages: []interface{}{
//   		&StageProperty{
//   			DurationInMinutes: jsii.Number(123),
//
//   			// the properties below are optional
//   			Targets: []interface{}{
//   				&TargetsProperty{
//   					ChannelTargetInfo: &ChannelTargetInfoProperty{
//   						ChannelId: jsii.String("channelId"),
//   						RetryIntervalInMinutes: jsii.Number(123),
//   					},
//   					ContactTargetInfo: &ContactTargetInfoProperty{
//   						ContactId: jsii.String("contactId"),
//   						IsEssential: jsii.Boolean(false),
//   					},
//   				},
//   			},
//   		},
//   	},
//   }
//
type CfnPlanProps struct {
	// The Amazon Resource Name (ARN) of the contact.
	ContactId *string `field:"required" json:"contactId" yaml:"contactId"`
	// The Amazon Resource Names (ARNs) of the on-call rotations associated with the plan.
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// A list of stages that the escalation plan or engagement plan uses to engage contacts and contact methods.
	Stages interface{} `field:"optional" json:"stages" yaml:"stages"`
}


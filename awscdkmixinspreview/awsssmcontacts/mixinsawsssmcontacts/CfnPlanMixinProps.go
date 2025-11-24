package mixinsawsssmcontacts


// Properties for CfnPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPlanMixinProps := &CfnPlanMixinProps{
//   	ContactId: jsii.String("contactId"),
//   	RotationIds: []*string{
//   		jsii.String("rotationIds"),
//   	},
//   	Stages: []interface{}{
//   		&StageProperty{
//   			DurationInMinutes: jsii.Number(123),
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-plan.html
//
type CfnPlanMixinProps struct {
	// The Amazon Resource Name (ARN) of the contact.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-plan.html#cfn-ssmcontacts-plan-contactid
	//
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
	// The Amazon Resource Names (ARNs) of the on-call rotations associated with the plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-plan.html#cfn-ssmcontacts-plan-rotationids
	//
	RotationIds *[]*string `field:"optional" json:"rotationIds" yaml:"rotationIds"`
	// A list of stages that the escalation plan or engagement plan uses to engage contacts and contact methods.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-plan.html#cfn-ssmcontacts-plan-stages
	//
	Stages interface{} `field:"optional" json:"stages" yaml:"stages"`
}


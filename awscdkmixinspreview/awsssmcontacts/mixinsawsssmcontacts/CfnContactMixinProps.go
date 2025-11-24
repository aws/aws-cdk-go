package mixinsawsssmcontacts

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContactPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnContactMixinProps := &CfnContactMixinProps{
//   	Alias: jsii.String("alias"),
//   	DisplayName: jsii.String("displayName"),
//   	Plan: []interface{}{
//   		&StageProperty{
//   			DurationInMinutes: jsii.Number(123),
//   			RotationIds: []*string{
//   				jsii.String("rotationIds"),
//   			},
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html
//
type CfnContactMixinProps struct {
	// The unique and identifiable alias of the contact or escalation plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html#cfn-ssmcontacts-contact-alias
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The full name of the contact or escalation plan.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html#cfn-ssmcontacts-contact-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A list of stages.
	//
	// A contact has an engagement plan with stages that contact specified contact channels. An escalation plan uses stages that contact specified contacts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html#cfn-ssmcontacts-contact-plan
	//
	Plan interface{} `field:"optional" json:"plan" yaml:"plan"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html#cfn-ssmcontacts-contact-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of contact.
	//
	// - `PERSONAL` : A single, individual contact.
	// - `ESCALATION` : An escalation plan.
	// - `ONCALL_SCHEDULE` : An on-call schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmcontacts-contact.html#cfn-ssmcontacts-contact-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


package previewawsarcregionswitchmixins


// Defines a condition that must be met for a trigger to fire.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggerConditionProperty := &TriggerConditionProperty{
//   	AssociatedAlarmName: jsii.String("associatedAlarmName"),
//   	Condition: jsii.String("condition"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-triggercondition.html
//
type CfnPlanPropsMixin_TriggerConditionProperty struct {
	// The name of the CloudWatch alarm associated with the condition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-triggercondition.html#cfn-arcregionswitch-plan-triggercondition-associatedalarmname
	//
	AssociatedAlarmName *string `field:"optional" json:"associatedAlarmName" yaml:"associatedAlarmName"`
	// The condition that must be met.
	//
	// Valid values include ALARM and OK.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-arcregionswitch-plan-triggercondition.html#cfn-arcregionswitch-plan-triggercondition-condition
	//
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
}


package awsevents


// This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionProperty := &ConditionProperty{
//   	Key: jsii.String("key"),
//   	Type: jsii.String("type"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html
//
type CfnEventBusPolicy_ConditionProperty struct {
	// Specifies the value for the key.
	//
	// Currently, this must be the ID of the organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies the type of condition.
	//
	// Currently the only supported value is StringEquals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the key for the condition.
	//
	// Currently the only supported key is aws:PrincipalOrgID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


package awsevents


// A JSON string which you can use to limit the event bus permissions you are granting to only accounts that fulfill the condition.
//
// Currently, the only supported condition is membership in a certain AWS organization. The string must contain `Type` , `Key` , and `Value` fields. The `Value` field specifies the ID of the AWS organization. Following is an example value for `Condition` :
//
// `'{"Type" : "StringEquals", "Key": "aws:PrincipalOrgID", "Value": "o-1234567890"}'`.
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
	// Specifies the key for the condition.
	//
	// Currently the only supported key is `aws:PrincipalOrgID` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies the type of condition.
	//
	// Currently the only supported value is `StringEquals` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the value for the key.
	//
	// Currently, this must be the ID of the organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


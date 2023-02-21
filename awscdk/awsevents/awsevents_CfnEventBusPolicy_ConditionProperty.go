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
//   conditionProperty := &conditionProperty{
//   	key: jsii.String("key"),
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnEventBusPolicy_ConditionProperty struct {
	// Specifies the key for the condition.
	//
	// Currently the only supported key is `aws:PrincipalOrgID` .
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Specifies the type of condition.
	//
	// Currently the only supported value is `StringEquals` .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the value for the key.
	//
	// Currently, this must be the ID of the organization.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


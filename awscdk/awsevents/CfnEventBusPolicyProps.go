package awsevents


// Properties for defining a `CfnEventBusPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var statement interface{}
//
//   cfnEventBusPolicyProps := &CfnEventBusPolicyProps{
//   	StatementId: jsii.String("statementId"),
//
//   	// the properties below are optional
//   	Action: jsii.String("action"),
//   	Condition: &ConditionProperty{
//   		Key: jsii.String("key"),
//   		Type: jsii.String("type"),
//   		Value: jsii.String("value"),
//   	},
//   	EventBusName: jsii.String("eventBusName"),
//   	Principal: jsii.String("principal"),
//   	Statement: statement,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html
//
type CfnEventBusPolicyProps struct {
	// An identifier string for the external account that you are granting permissions to.
	//
	// If you later want to revoke the permission for this external account, specify this `StatementId` when you run [RemovePermission](https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_RemovePermission.html) .
	//
	// > Each `StatementId` must be unique.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-statementid
	//
	StatementId *string `field:"required" json:"statementId" yaml:"statementId"`
	// The action that you are enabling the other account to perform.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-action
	//
	// Deprecated: this property has been deprecated.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// This parameter enables you to limit the permission to accounts that fulfill a certain condition, such as being a member of a certain AWS organization.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-condition
	//
	// Deprecated: this property has been deprecated.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// The name of the event bus associated with the rule.
	//
	// If you omit this, the default event bus is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-eventbusname
	//
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// The 12-digit AWS account ID that you are permitting to put events to your default event bus.
	//
	// Specify "*" to permit any account to put events to your default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-principal
	//
	// Deprecated: this property has been deprecated.
	Principal *string `field:"optional" json:"principal" yaml:"principal"`
	// A JSON string that describes the permission policy statement.
	//
	// You can include a `Policy` parameter in the request instead of using the `StatementId` , `Action` , `Principal` , or `Condition` parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-statement
	//
	Statement interface{} `field:"optional" json:"statement" yaml:"statement"`
}


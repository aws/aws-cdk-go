package awseventsv2


// Properties for CfnResourcePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var policyDocument interface{}
//
//   cfnResourcePolicyMixinProps := &CfnResourcePolicyMixinProps{
//   	EventBusArn: jsii.String("eventBusArn"),
//   	PolicyDocument: policyDocument,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-resourcepolicy.html
//
type CfnResourcePolicyMixinProps struct {
	// The Amazon Resource Name (ARN) of the event bus whose resource policy this is.
	//
	// The bus must already exist. This resource does not create it.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-resourcepolicy.html#cfn-eventsv2-resourcepolicy-eventbusarn
	//
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
	// The resource policy document, as a JSON object.
	//
	// The document can be up to 20 KB. This quota is adjustable. An empty object is not a valid policy. To remove the policy, delete this resource. The principals in the document must exist and be visible to the service when the policy is written. When you create a new IAM role or user, that principal might not be immediately visible to the service. You might need to enforce a delay before you include it in the document. For more information, see "Changes that I make are not always immediately visible" in the IAM User Guide. Declare Version. Write AWS account and role principals as ARNs rather than as account IDs. Write a single Action, Resource, or principal value as a scalar rather than as a one-element list. The service returns these forms as you wrote them. It normalizes other forms, and a normalized value can appear as drift. A stack update replaces the whole policy with this document, including any change made outside CloudFormation. For more information about event bus resource policies, see the Amazon EventBridge User Guide.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eventsv2-resourcepolicy.html#cfn-eventsv2-resourcepolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
}


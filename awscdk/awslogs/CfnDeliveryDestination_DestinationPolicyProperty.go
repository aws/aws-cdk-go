package awslogs


// An IAM policy that grants permissions to CloudWatch Logs to deliver logs cross-account to a specified destination in this account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryDestinationPolicy interface{}
//
//   destinationPolicyProperty := &DestinationPolicyProperty{
//   	DeliveryDestinationName: jsii.String("deliveryDestinationName"),
//   	DeliveryDestinationPolicy: deliveryDestinationPolicy,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-deliverydestination-destinationpolicy.html
//
type CfnDeliveryDestination_DestinationPolicyProperty struct {
	// A name for an existing destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-deliverydestination-destinationpolicy.html#cfn-logs-deliverydestination-destinationpolicy-deliverydestinationname
	//
	DeliveryDestinationName *string `field:"optional" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
	// Creates or updates an access policy associated with an existing destination.
	//
	// An access policy is an [IAM policy document](https://docs.aws.amazon.com/IAM/latest/UserGuide/policies_overview.html) that is used to authorize claims to register a subscription filter against a given destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-deliverydestination-destinationpolicy.html#cfn-logs-deliverydestination-destinationpolicy-deliverydestinationpolicy
	//
	DeliveryDestinationPolicy interface{} `field:"optional" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
}


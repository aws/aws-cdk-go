package awslogs


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
	// The name of the delivery destination to assign this policy to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-deliverydestination-destinationpolicy.html#cfn-logs-deliverydestination-destinationpolicy-deliverydestinationname
	//
	DeliveryDestinationName *string `field:"optional" json:"deliveryDestinationName" yaml:"deliveryDestinationName"`
	// The contents of the policy attached to the delivery destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-deliverydestination-destinationpolicy.html#cfn-logs-deliverydestination-destinationpolicy-deliverydestinationpolicy
	//
	DeliveryDestinationPolicy interface{} `field:"optional" json:"deliveryDestinationPolicy" yaml:"deliveryDestinationPolicy"`
}


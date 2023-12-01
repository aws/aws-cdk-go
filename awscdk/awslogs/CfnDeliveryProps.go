package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnDelivery`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeliveryProps := &CfnDeliveryProps{
//   	DeliveryDestinationArn: jsii.String("deliveryDestinationArn"),
//   	DeliverySourceName: jsii.String("deliverySourceName"),
//
//   	// the properties below are optional
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html
//
type CfnDeliveryProps struct {
	// Amazon Resource Names (ARNs) uniquely identify AWS resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-deliverydestinationarn
	//
	DeliveryDestinationArn *string `field:"required" json:"deliveryDestinationArn" yaml:"deliveryDestinationArn"`
	// The value of the DeliverySourceName property for this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-deliverysourcename
	//
	DeliverySourceName *string `field:"required" json:"deliverySourceName" yaml:"deliverySourceName"`
	// The value of the Tags property for this object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html#cfn-logs-delivery-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


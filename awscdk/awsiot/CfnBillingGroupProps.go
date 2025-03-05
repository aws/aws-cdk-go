package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBillingGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBillingGroupProps := &CfnBillingGroupProps{
//   	BillingGroupName: jsii.String("billingGroupName"),
//   	BillingGroupProperties: &BillingGroupPropertiesProperty{
//   		BillingGroupDescription: jsii.String("billingGroupDescription"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-billinggroup.html
//
type CfnBillingGroupProps struct {
	// The name of the billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-billinggroup.html#cfn-iot-billinggroup-billinggroupname
	//
	BillingGroupName *string `field:"optional" json:"billingGroupName" yaml:"billingGroupName"`
	// The properties of the billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-billinggroup.html#cfn-iot-billinggroup-billinggroupproperties
	//
	BillingGroupProperties interface{} `field:"optional" json:"billingGroupProperties" yaml:"billingGroupProperties"`
	// Metadata which can be used to manage the billing group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-billinggroup.html#cfn-iot-billinggroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


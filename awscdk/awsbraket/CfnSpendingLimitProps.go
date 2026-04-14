package awsbraket

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnSpendingLimit`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSpendingLimitProps := &CfnSpendingLimitProps{
//   	DeviceArn: jsii.String("deviceArn"),
//   	SpendingLimit: jsii.String("spendingLimit"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TimePeriod: &TimePeriodProperty{
//   		EndAt: jsii.String("endAt"),
//   		StartAt: jsii.String("startAt"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html
//
type CfnSpendingLimitProps struct {
	// The Amazon Resource Name (ARN) of the quantum device to apply the spending limit to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html#cfn-braket-spendinglimit-devicearn
	//
	DeviceArn *string `field:"required" json:"deviceArn" yaml:"deviceArn"`
	// The maximum amount that can be spent on the specified device, in USD.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html#cfn-braket-spendinglimit-spendinglimit
	//
	SpendingLimit *string `field:"required" json:"spendingLimit" yaml:"spendingLimit"`
	// The tags to apply to the spending limit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html#cfn-braket-spendinglimit-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Defines a time range for spending limits, specifying when the limit is active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-braket-spendinglimit.html#cfn-braket-spendinglimit-timeperiod
	//
	TimePeriod interface{} `field:"optional" json:"timePeriod" yaml:"timePeriod"`
}


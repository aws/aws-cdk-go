package awsbcmpricingcalculator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnBillEstimatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnBillEstimateMixinProps := &CfnBillEstimateMixinProps{
//   	BillScenarioId: jsii.String("billScenarioId"),
//   	ExpiresAt: jsii.String("expiresAt"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billestimate.html
//
type CfnBillEstimateMixinProps struct {
	// The unique identifier of the bill scenario this estimate is based on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billestimate.html#cfn-bcmpricingcalculator-billestimate-billscenarioid
	//
	BillScenarioId *string `field:"optional" json:"billScenarioId" yaml:"billScenarioId"`
	// The timestamp when the bill estimate expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billestimate.html#cfn-bcmpricingcalculator-billestimate-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The name of the bill estimate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billestimate.html#cfn-bcmpricingcalculator-billestimate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billestimate.html#cfn-bcmpricingcalculator-billestimate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


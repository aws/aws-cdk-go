package awsbcmpricingcalculator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnWorkloadEstimatePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnWorkloadEstimateMixinProps := &CfnWorkloadEstimateMixinProps{
//   	ExpiresAt: jsii.String("expiresAt"),
//   	Name: jsii.String("name"),
//   	RateType: jsii.String("rateType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html
//
type CfnWorkloadEstimateMixinProps struct {
	// The timestamp when the workload estimate will expire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html#cfn-bcmpricingcalculator-workloadestimate-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// The name of the workload estimate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html#cfn-bcmpricingcalculator-workloadestimate-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html#cfn-bcmpricingcalculator-workloadestimate-ratetype
	//
	RateType *string `field:"optional" json:"rateType" yaml:"rateType"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-workloadestimate.html#cfn-bcmpricingcalculator-workloadestimate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


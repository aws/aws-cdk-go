package awsbcmpricingcalculator

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnBillScenario`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBillScenarioProps := &CfnBillScenarioProps{
//   	CostCategoryGroupSharingPreferenceArn: jsii.String("costCategoryGroupSharingPreferenceArn"),
//   	ExpiresAt: jsii.String("expiresAt"),
//   	GroupSharingPreference: jsii.String("groupSharingPreference"),
//   	Name: jsii.String("name"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html
//
type CfnBillScenarioProps struct {
	// The ARN of the cost category group sharing preference.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html#cfn-bcmpricingcalculator-billscenario-costcategorygroupsharingpreferencearn
	//
	CostCategoryGroupSharingPreferenceArn *string `field:"optional" json:"costCategoryGroupSharingPreferenceArn" yaml:"costCategoryGroupSharingPreferenceArn"`
	// The timestamp when the bill scenario expires.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html#cfn-bcmpricingcalculator-billscenario-expiresat
	//
	ExpiresAt *string `field:"optional" json:"expiresAt" yaml:"expiresAt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html#cfn-bcmpricingcalculator-billscenario-groupsharingpreference
	//
	GroupSharingPreference *string `field:"optional" json:"groupSharingPreference" yaml:"groupSharingPreference"`
	// The name of the bill scenario.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html#cfn-bcmpricingcalculator-billscenario-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bcmpricingcalculator-billscenario.html#cfn-bcmpricingcalculator-billscenario-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


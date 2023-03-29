package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   whatIfRangeScenarioProperty := &WhatIfRangeScenarioProperty{
//   	EndDate: jsii.String("endDate"),
//   	StartDate: jsii.String("startDate"),
//   	Value: jsii.Number(123),
//   }
//
type CfnTemplate_WhatIfRangeScenarioProperty struct {
	// `CfnTemplate.WhatIfRangeScenarioProperty.EndDate`.
	EndDate *string `field:"required" json:"endDate" yaml:"endDate"`
	// `CfnTemplate.WhatIfRangeScenarioProperty.StartDate`.
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// `CfnTemplate.WhatIfRangeScenarioProperty.Value`.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


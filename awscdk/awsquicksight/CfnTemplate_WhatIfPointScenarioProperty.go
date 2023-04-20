package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   whatIfPointScenarioProperty := &WhatIfPointScenarioProperty{
//   	Date: jsii.String("date"),
//   	Value: jsii.Number(123),
//   }
//
type CfnTemplate_WhatIfPointScenarioProperty struct {
	// `CfnTemplate.WhatIfPointScenarioProperty.Date`.
	Date *string `field:"required" json:"date" yaml:"date"`
	// `CfnTemplate.WhatIfPointScenarioProperty.Value`.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


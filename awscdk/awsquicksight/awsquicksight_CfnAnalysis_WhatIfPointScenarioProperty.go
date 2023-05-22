package awsquicksight


// Provides the forecast to meet the target for a particular date.
//
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
type CfnAnalysis_WhatIfPointScenarioProperty struct {
	// The date that you need the forecast results for.
	Date *string `field:"required" json:"date" yaml:"date"`
	// The target value that you want to meet for the provided date.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


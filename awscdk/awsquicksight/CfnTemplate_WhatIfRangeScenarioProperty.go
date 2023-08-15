package awsquicksight


// Provides the forecast to meet the target for a particular date range.
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifrangescenario.html
//
type CfnTemplate_WhatIfRangeScenarioProperty struct {
	// The end date in the date range that you need the forecast results for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifrangescenario.html#cfn-quicksight-template-whatifrangescenario-enddate
	//
	EndDate *string `field:"required" json:"endDate" yaml:"endDate"`
	// The start date in the date range that you need the forecast results for.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifrangescenario.html#cfn-quicksight-template-whatifrangescenario-startdate
	//
	StartDate *string `field:"required" json:"startDate" yaml:"startDate"`
	// The target value that you want to meet for the provided date range.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-whatifrangescenario.html#cfn-quicksight-template-whatifrangescenario-value
	//
	// Default: - 0.
	//
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


package awsquicksight


// The forecast scenario of a forecast in the line chart.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forecastScenarioProperty := &ForecastScenarioProperty{
//   	WhatIfPointScenario: &WhatIfPointScenarioProperty{
//   		Date: jsii.String("date"),
//   		Value: jsii.Number(123),
//   	},
//   	WhatIfRangeScenario: &WhatIfRangeScenarioProperty{
//   		EndDate: jsii.String("endDate"),
//   		StartDate: jsii.String("startDate"),
//   		Value: jsii.Number(123),
//   	},
//   }
//
type CfnTemplate_ForecastScenarioProperty struct {
	// The what-if analysis forecast setup with the target date.
	WhatIfPointScenario interface{} `field:"optional" json:"whatIfPointScenario" yaml:"whatIfPointScenario"`
	// The what-if analysis forecast setup with the date range.
	WhatIfRangeScenario interface{} `field:"optional" json:"whatIfRangeScenario" yaml:"whatIfRangeScenario"`
}


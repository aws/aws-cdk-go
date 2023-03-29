package awsquicksight


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
type CfnDashboard_ForecastScenarioProperty struct {
	// `CfnDashboard.ForecastScenarioProperty.WhatIfPointScenario`.
	WhatIfPointScenario interface{} `field:"optional" json:"whatIfPointScenario" yaml:"whatIfPointScenario"`
	// `CfnDashboard.ForecastScenarioProperty.WhatIfRangeScenario`.
	WhatIfRangeScenario interface{} `field:"optional" json:"whatIfRangeScenario" yaml:"whatIfRangeScenario"`
}


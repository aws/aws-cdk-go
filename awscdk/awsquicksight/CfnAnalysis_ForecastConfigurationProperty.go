package awsquicksight


// The forecast configuration that is used in a line chart's display properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   forecastConfigurationProperty := &ForecastConfigurationProperty{
//   	ForecastProperties: &TimeBasedForecastPropertiesProperty{
//   		LowerBoundary: jsii.Number(123),
//   		PeriodsBackward: jsii.Number(123),
//   		PeriodsForward: jsii.Number(123),
//   		PredictionInterval: jsii.Number(123),
//   		Seasonality: jsii.Number(123),
//   		UpperBoundary: jsii.Number(123),
//   	},
//   	Scenario: &ForecastScenarioProperty{
//   		WhatIfPointScenario: &WhatIfPointScenarioProperty{
//   			Date: jsii.String("date"),
//   			Value: jsii.Number(123),
//   		},
//   		WhatIfRangeScenario: &WhatIfRangeScenarioProperty{
//   			EndDate: jsii.String("endDate"),
//   			StartDate: jsii.String("startDate"),
//   			Value: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnAnalysis_ForecastConfigurationProperty struct {
	// The forecast properties setup of a forecast in the line chart.
	ForecastProperties interface{} `field:"optional" json:"forecastProperties" yaml:"forecastProperties"`
	// The forecast scenario of a forecast in the line chart.
	Scenario interface{} `field:"optional" json:"scenario" yaml:"scenario"`
}


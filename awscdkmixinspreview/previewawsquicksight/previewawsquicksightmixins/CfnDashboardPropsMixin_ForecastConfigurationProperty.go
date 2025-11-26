package previewawsquicksightmixins


// The forecast configuration that is used in a line chart's display properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html
//
type CfnDashboardPropsMixin_ForecastConfigurationProperty struct {
	// The forecast properties setup of a forecast in the line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html#cfn-quicksight-dashboard-forecastconfiguration-forecastproperties
	//
	ForecastProperties interface{} `field:"optional" json:"forecastProperties" yaml:"forecastProperties"`
	// The forecast scenario of a forecast in the line chart.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-forecastconfiguration.html#cfn-quicksight-dashboard-forecastconfiguration-scenario
	//
	Scenario interface{} `field:"optional" json:"scenario" yaml:"scenario"`
}


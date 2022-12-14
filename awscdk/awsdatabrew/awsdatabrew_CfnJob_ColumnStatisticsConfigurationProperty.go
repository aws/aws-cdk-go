package awsdatabrew


// Configuration for column evaluations for a profile job.
//
// ColumnStatisticsConfiguration can be used to select evaluations and override parameters of evaluations for particular columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   columnStatisticsConfigurationProperty := &columnStatisticsConfigurationProperty{
//   	statistics: &statisticsConfigurationProperty{
//   		includedStatistics: []*string{
//   			jsii.String("includedStatistics"),
//   		},
//   		overrides: []interface{}{
//   			&statisticOverrideProperty{
//   				parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				statistic: jsii.String("statistic"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	selectors: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
type CfnJob_ColumnStatisticsConfigurationProperty struct {
	// Configuration for evaluations.
	//
	// Statistics can be used to select evaluations and override parameters of evaluations.
	Statistics interface{} `field:"required" json:"statistics" yaml:"statistics"`
	// List of column selectors.
	//
	// Selectors can be used to select columns from the dataset. When selectors are undefined, configuration will be applied to all supported columns.
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}


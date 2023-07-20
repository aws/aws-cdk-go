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
//   columnStatisticsConfigurationProperty := &ColumnStatisticsConfigurationProperty{
//   	Statistics: &StatisticsConfigurationProperty{
//   		IncludedStatistics: []*string{
//   			jsii.String("includedStatistics"),
//   		},
//   		Overrides: []interface{}{
//   			&StatisticOverrideProperty{
//   				Parameters: map[string]*string{
//   					"parametersKey": jsii.String("parameters"),
//   				},
//   				Statistic: jsii.String("statistic"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	Selectors: []interface{}{
//   		&ColumnSelectorProperty{
//   			Name: jsii.String("name"),
//   			Regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-columnstatisticsconfiguration.html
//
type CfnJob_ColumnStatisticsConfigurationProperty struct {
	// Configuration for evaluations.
	//
	// Statistics can be used to select evaluations and override parameters of evaluations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-columnstatisticsconfiguration.html#cfn-databrew-job-columnstatisticsconfiguration-statistics
	//
	Statistics interface{} `field:"required" json:"statistics" yaml:"statistics"`
	// List of column selectors.
	//
	// Selectors can be used to select columns from the dataset. When selectors are undefined, configuration will be applied to all supported columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-columnstatisticsconfiguration.html#cfn-databrew-job-columnstatisticsconfiguration-selectors
	//
	Selectors interface{} `field:"optional" json:"selectors" yaml:"selectors"`
}


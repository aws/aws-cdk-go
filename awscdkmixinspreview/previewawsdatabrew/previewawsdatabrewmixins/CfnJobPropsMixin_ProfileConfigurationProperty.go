package previewawsdatabrewmixins


// Configuration for profile jobs.
//
// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   profileConfigurationProperty := &ProfileConfigurationProperty{
//   	ColumnStatisticsConfigurations: []interface{}{
//   		&ColumnStatisticsConfigurationProperty{
//   			Selectors: []interface{}{
//   				&ColumnSelectorProperty{
//   					Name: jsii.String("name"),
//   					Regex: jsii.String("regex"),
//   				},
//   			},
//   			Statistics: &StatisticsConfigurationProperty{
//   				IncludedStatistics: []*string{
//   					jsii.String("includedStatistics"),
//   				},
//   				Overrides: []interface{}{
//   					&StatisticOverrideProperty{
//   						Parameters: map[string]*string{
//   							"parametersKey": jsii.String("parameters"),
//   						},
//   						Statistic: jsii.String("statistic"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	DatasetStatisticsConfiguration: &StatisticsConfigurationProperty{
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
//   	EntityDetectorConfiguration: &EntityDetectorConfigurationProperty{
//   		AllowedStatistics: &AllowedStatisticsProperty{
//   			Statistics: []*string{
//   				jsii.String("statistics"),
//   			},
//   		},
//   		EntityTypes: []*string{
//   			jsii.String("entityTypes"),
//   		},
//   	},
//   	ProfileColumns: []interface{}{
//   		&ColumnSelectorProperty{
//   			Name: jsii.String("name"),
//   			Regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html
//
type CfnJobPropsMixin_ProfileConfigurationProperty struct {
	// List of configurations for column evaluations.
	//
	// ColumnStatisticsConfigurations are used to select evaluations and override parameters of evaluations for particular columns. When ColumnStatisticsConfigurations is undefined, the profile job will profile all supported columns and run all supported evaluations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-columnstatisticsconfigurations
	//
	ColumnStatisticsConfigurations interface{} `field:"optional" json:"columnStatisticsConfigurations" yaml:"columnStatisticsConfigurations"`
	// Configuration for inter-column evaluations.
	//
	// Configuration can be used to select evaluations and override parameters of evaluations. When configuration is undefined, the profile job will run all supported inter-column evaluations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-datasetstatisticsconfiguration
	//
	DatasetStatisticsConfiguration interface{} `field:"optional" json:"datasetStatisticsConfiguration" yaml:"datasetStatisticsConfiguration"`
	// Configuration of entity detection for a profile job.
	//
	// When undefined, entity detection is disabled.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-entitydetectorconfiguration
	//
	EntityDetectorConfiguration interface{} `field:"optional" json:"entityDetectorConfiguration" yaml:"entityDetectorConfiguration"`
	// List of column selectors.
	//
	// ProfileColumns can be used to select columns from the dataset. When ProfileColumns is undefined, the profile job will profile all supported columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-profileconfiguration.html#cfn-databrew-job-profileconfiguration-profilecolumns
	//
	ProfileColumns interface{} `field:"optional" json:"profileColumns" yaml:"profileColumns"`
}


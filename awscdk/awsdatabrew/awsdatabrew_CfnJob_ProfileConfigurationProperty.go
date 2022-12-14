package awsdatabrew


// Configuration for profile jobs.
//
// Configuration can be used to select columns, do evaluations, and override default parameters of evaluations. When configuration is undefined, the profile job will apply default settings to all supported columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   profileConfigurationProperty := &profileConfigurationProperty{
//   	columnStatisticsConfigurations: []interface{}{
//   		&columnStatisticsConfigurationProperty{
//   			statistics: &statisticsConfigurationProperty{
//   				includedStatistics: []*string{
//   					jsii.String("includedStatistics"),
//   				},
//   				overrides: []interface{}{
//   					&statisticOverrideProperty{
//   						parameters: map[string]*string{
//   							"parametersKey": jsii.String("parameters"),
//   						},
//   						statistic: jsii.String("statistic"),
//   					},
//   				},
//   			},
//
//   			// the properties below are optional
//   			selectors: []interface{}{
//   				&columnSelectorProperty{
//   					name: jsii.String("name"),
//   					regex: jsii.String("regex"),
//   				},
//   			},
//   		},
//   	},
//   	datasetStatisticsConfiguration: &statisticsConfigurationProperty{
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
//   	entityDetectorConfiguration: &entityDetectorConfigurationProperty{
//   		entityTypes: []*string{
//   			jsii.String("entityTypes"),
//   		},
//
//   		// the properties below are optional
//   		allowedStatistics: &allowedStatisticsProperty{
//   			statistics: []*string{
//   				jsii.String("statistics"),
//   			},
//   		},
//   	},
//   	profileColumns: []interface{}{
//   		&columnSelectorProperty{
//   			name: jsii.String("name"),
//   			regex: jsii.String("regex"),
//   		},
//   	},
//   }
//
type CfnJob_ProfileConfigurationProperty struct {
	// List of configurations for column evaluations.
	//
	// ColumnStatisticsConfigurations are used to select evaluations and override parameters of evaluations for particular columns. When ColumnStatisticsConfigurations is undefined, the profile job will profile all supported columns and run all supported evaluations.
	ColumnStatisticsConfigurations interface{} `field:"optional" json:"columnStatisticsConfigurations" yaml:"columnStatisticsConfigurations"`
	// Configuration for inter-column evaluations.
	//
	// Configuration can be used to select evaluations and override parameters of evaluations. When configuration is undefined, the profile job will run all supported inter-column evaluations.
	DatasetStatisticsConfiguration interface{} `field:"optional" json:"datasetStatisticsConfiguration" yaml:"datasetStatisticsConfiguration"`
	// Configuration of entity detection for a profile job.
	//
	// When undefined, entity detection is disabled.
	EntityDetectorConfiguration interface{} `field:"optional" json:"entityDetectorConfiguration" yaml:"entityDetectorConfiguration"`
	// List of column selectors.
	//
	// ProfileColumns can be used to select columns from the dataset. When ProfileColumns is undefined, the profile job will profile all supported columns.
	ProfileColumns interface{} `field:"optional" json:"profileColumns" yaml:"profileColumns"`
}


package awsdatabrew


// Configuration of evaluations for a profile job.
//
// This configuration can be used to select evaluations and override the parameters of selected evaluations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticsConfigurationProperty := &StatisticsConfigurationProperty{
//   	IncludedStatistics: []*string{
//   		jsii.String("includedStatistics"),
//   	},
//   	Overrides: []interface{}{
//   		&StatisticOverrideProperty{
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			Statistic: jsii.String("statistic"),
//   		},
//   	},
//   }
//
type CfnJob_StatisticsConfigurationProperty struct {
	// List of included evaluations.
	//
	// When the list is undefined, all supported evaluations will be included.
	IncludedStatistics *[]*string `field:"optional" json:"includedStatistics" yaml:"includedStatistics"`
	// List of overrides for evaluations.
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


package previewawsdatabrewmixins


// Configuration of evaluations for a profile job.
//
// This configuration can be used to select evaluations and override the parameters of selected evaluations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html
//
type CfnJobPropsMixin_StatisticsConfigurationProperty struct {
	// List of included evaluations.
	//
	// When the list is undefined, all supported evaluations will be included.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html#cfn-databrew-job-statisticsconfiguration-includedstatistics
	//
	IncludedStatistics *[]*string `field:"optional" json:"includedStatistics" yaml:"includedStatistics"`
	// List of overrides for evaluations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticsconfiguration.html#cfn-databrew-job-statisticsconfiguration-overrides
	//
	Overrides interface{} `field:"optional" json:"overrides" yaml:"overrides"`
}


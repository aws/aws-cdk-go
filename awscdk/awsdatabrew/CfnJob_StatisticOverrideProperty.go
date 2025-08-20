package awsdatabrew


// Override of a particular evaluation for a profile job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticOverrideProperty := &StatisticOverrideProperty{
//   	Parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	Statistic: jsii.String("statistic"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticoverride.html
//
type CfnJob_StatisticOverrideProperty struct {
	// A map that includes overrides of an evaluation’s parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticoverride.html#cfn-databrew-job-statisticoverride-parameters
	//
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The name of an evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-databrew-job-statisticoverride.html#cfn-databrew-job-statisticoverride-statistic
	//
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}


package awsdatabrew


// Override of a particular evaluation for a profile job.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   statisticOverrideProperty := &statisticOverrideProperty{
//   	parameters: map[string]*string{
//   		"parametersKey": jsii.String("parameters"),
//   	},
//   	statistic: jsii.String("statistic"),
//   }
//
type CfnJob_StatisticOverrideProperty struct {
	// A map that includes overrides of an evaluation’s parameters.
	Parameters interface{} `field:"required" json:"parameters" yaml:"parameters"`
	// The name of an evaluation.
	Statistic *string `field:"required" json:"statistic" yaml:"statistic"`
}


package awspersonalize


// The metric to optimize during hyperparameter optimization (HPO).
//
// > Amazon Personalize doesn't support configuring the `hpoObjective` at this time.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoObjectiveProperty := &hpoObjectiveProperty{
//   	metricName: jsii.String("metricName"),
//   	metricRegex: jsii.String("metricRegex"),
//   	type: jsii.String("type"),
//   }
//
type CfnSolution_HpoObjectiveProperty struct {
	// The name of the metric.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A regular expression for finding the metric in the training job logs.
	MetricRegex *string `field:"optional" json:"metricRegex" yaml:"metricRegex"`
	// The type of the metric.
	//
	// Valid values are `Maximize` and `Minimize` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}


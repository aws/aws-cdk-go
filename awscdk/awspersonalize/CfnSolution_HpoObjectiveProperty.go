package awspersonalize


// The metric to optimize during HPO.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hpoObjectiveProperty := &HpoObjectiveProperty{
//   	MetricName: jsii.String("metricName"),
//   	MetricRegex: jsii.String("metricRegex"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoobjective.html
//
type CfnSolution_HpoObjectiveProperty struct {
	// The name of the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoobjective.html#cfn-personalize-solution-hpoobjective-metricname
	//
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// A regular expression for finding the metric in the training job logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoobjective.html#cfn-personalize-solution-hpoobjective-metricregex
	//
	MetricRegex *string `field:"optional" json:"metricRegex" yaml:"metricRegex"`
	// The type of the metric.
	//
	// Valid values are Maximize and Minimize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-personalize-solution-hpoobjective.html#cfn-personalize-solution-hpoobjective-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


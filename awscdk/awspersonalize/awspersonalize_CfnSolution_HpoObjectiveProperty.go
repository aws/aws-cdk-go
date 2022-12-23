package awspersonalize


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
	// `CfnSolution.HpoObjectiveProperty.MetricName`.
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// `CfnSolution.HpoObjectiveProperty.MetricRegex`.
	MetricRegex *string `field:"optional" json:"metricRegex" yaml:"metricRegex"`
	// `CfnSolution.HpoObjectiveProperty.Type`.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


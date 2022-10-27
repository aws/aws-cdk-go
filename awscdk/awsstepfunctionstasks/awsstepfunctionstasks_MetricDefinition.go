package awsstepfunctionstasks


// Specifies the metric name and regular expressions used to parse algorithm logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinition := &metricDefinition{
//   	name: jsii.String("name"),
//   	regex: jsii.String("regex"),
//   }
//
type MetricDefinition struct {
	// Name of the metric.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Regular expression that searches the output of a training job and gets the value of the metric.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}


package awsstepfunctionstasks


// Specifies the metric name and regular expressions used to parse algorithm logs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metricDefinition := &MetricDefinition{
//   	Name: jsii.String("name"),
//   	Regex: jsii.String("regex"),
//   }
//
// Experimental.
type MetricDefinition struct {
	// Name of the metric.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Regular expression that searches the output of a training job and gets the value of the metric.
	// Experimental.
	Regex *string `field:"required" json:"regex" yaml:"regex"`
}


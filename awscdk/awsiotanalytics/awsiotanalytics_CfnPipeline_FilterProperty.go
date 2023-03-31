package awsiotanalytics


// An activity that filters a message based on its attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterProperty := &filterProperty{
//   	filter: jsii.String("filter"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_FilterProperty struct {
	// An expression that looks like an SQL WHERE clause that must return a Boolean value.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// The name of the 'filter' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}


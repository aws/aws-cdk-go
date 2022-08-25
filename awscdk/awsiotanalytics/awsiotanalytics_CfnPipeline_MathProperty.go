package awsiotanalytics


// An activity that computes an arithmetic expression using the message's attributes.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mathProperty := &mathProperty{
//   	attribute: jsii.String("attribute"),
//   	math: jsii.String("math"),
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_MathProperty struct {
	// The name of the attribute that contains the result of the math operation.
	Attribute *string `field:"required" json:"attribute" yaml:"attribute"`
	// An expression that uses one or more existing attributes and must return an integer value.
	Math *string `field:"required" json:"math" yaml:"math"`
	// The name of the 'math' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}


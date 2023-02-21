package awsiotanalytics


// Creates a new message using only the specified attributes from the original message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   selectAttributesProperty := &selectAttributesProperty{
//   	attributes: []*string{
//   		jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_SelectAttributesProperty struct {
	// A list of the attributes to select from the message.
	Attributes *[]*string `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'selectAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}


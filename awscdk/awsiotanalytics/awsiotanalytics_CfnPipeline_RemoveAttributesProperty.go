package awsiotanalytics


// An activity that removes attributes from a message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   removeAttributesProperty := &removeAttributesProperty{
//   	attributes: []*string{
//   		jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_RemoveAttributesProperty struct {
	// A list of 1-50 attributes to remove from the message.
	Attributes *[]*string `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'removeAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}


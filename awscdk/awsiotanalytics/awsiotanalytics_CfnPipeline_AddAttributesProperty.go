package awsiotanalytics


// An activity that adds other attributes based on existing attributes in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   addAttributesProperty := &addAttributesProperty{
//   	attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	next: jsii.String("next"),
//   }
//
type CfnPipeline_AddAttributesProperty struct {
	// A list of 1-50 "AttributeNameMapping" objects that map an existing attribute to a new attribute.
	//
	// > The existing attributes remain in the message, so if you want to remove the originals, use "RemoveAttributeActivity".
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The name of the 'addAttributes' activity.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	Next *string `field:"optional" json:"next" yaml:"next"`
}


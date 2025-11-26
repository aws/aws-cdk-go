package previewawsiotanalyticsmixins


// An activity that adds other attributes based on existing attributes in the message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   addAttributesProperty := &AddAttributesProperty{
//   	Attributes: map[string]*string{
//   		"attributesKey": jsii.String("attributes"),
//   	},
//   	Name: jsii.String("name"),
//   	Next: jsii.String("next"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-addattributes.html
//
type CfnPipelinePropsMixin_AddAttributesProperty struct {
	// A list of 1-50 "AttributeNameMapping" objects that map an existing attribute to a new attribute.
	//
	// > The existing attributes remain in the message, so if you want to remove the originals, use "RemoveAttributeActivity".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-addattributes.html#cfn-iotanalytics-pipeline-addattributes-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The name of the 'addAttributes' activity.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-addattributes.html#cfn-iotanalytics-pipeline-addattributes-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The next activity in the pipeline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-pipeline-addattributes.html#cfn-iotanalytics-pipeline-addattributes-next
	//
	Next *string `field:"optional" json:"next" yaml:"next"`
}


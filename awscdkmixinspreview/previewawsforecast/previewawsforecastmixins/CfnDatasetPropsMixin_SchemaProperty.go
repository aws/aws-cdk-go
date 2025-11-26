package previewawsforecastmixins


// Defines the fields of a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   schemaProperty := &SchemaProperty{
//   	Attributes: []interface{}{
//   		&AttributesItemsProperty{
//   			AttributeName: jsii.String("attributeName"),
//   			AttributeType: jsii.String("attributeType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-forecast-dataset-schema.html
//
type CfnDatasetPropsMixin_SchemaProperty struct {
	// An array of attributes specifying the name and type of each field in a dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-forecast-dataset-schema.html#cfn-forecast-dataset-schema-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
}


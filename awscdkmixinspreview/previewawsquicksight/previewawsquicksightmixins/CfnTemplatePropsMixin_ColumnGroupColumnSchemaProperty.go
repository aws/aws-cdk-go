package previewawsquicksightmixins


// A structure describing the name, data type, and geographic role of the columns.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnGroupColumnSchemaProperty := &ColumnGroupColumnSchemaProperty{
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columngroupcolumnschema.html
//
type CfnTemplatePropsMixin_ColumnGroupColumnSchemaProperty struct {
	// The name of the column group's column schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columngroupcolumnschema.html#cfn-quicksight-template-columngroupcolumnschema-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


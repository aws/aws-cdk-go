package previewawsquicksightmixins


// Dataset schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetSchemaProperty := &DataSetSchemaProperty{
//   	ColumnSchemaList: []interface{}{
//   		&ColumnSchemaProperty{
//   			DataType: jsii.String("dataType"),
//   			GeographicRole: jsii.String("geographicRole"),
//   			Name: jsii.String("name"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetschema.html
//
type CfnTemplatePropsMixin_DataSetSchemaProperty struct {
	// A structure containing the list of column schemas.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datasetschema.html#cfn-quicksight-template-datasetschema-columnschemalist
	//
	ColumnSchemaList interface{} `field:"optional" json:"columnSchemaList" yaml:"columnSchemaList"`
}


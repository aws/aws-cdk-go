package mixinsawsquicksight


// The column group schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnGroupSchemaProperty := &ColumnGroupSchemaProperty{
//   	ColumnGroupColumnSchemaList: []interface{}{
//   		&ColumnGroupColumnSchemaProperty{
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columngroupschema.html
//
type CfnTemplatePropsMixin_ColumnGroupSchemaProperty struct {
	// A structure containing the list of schemas for column group columns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columngroupschema.html#cfn-quicksight-template-columngroupschema-columngroupcolumnschemalist
	//
	ColumnGroupColumnSchemaList interface{} `field:"optional" json:"columnGroupColumnSchemaList" yaml:"columnGroupColumnSchemaList"`
	// The name of the column group schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columngroupschema.html#cfn-quicksight-template-columngroupschema-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


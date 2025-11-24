package mixinsawsquicksight


// The column schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   columnSchemaProperty := &ColumnSchemaProperty{
//   	DataType: jsii.String("dataType"),
//   	GeographicRole: jsii.String("geographicRole"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnschema.html
//
type CfnTemplatePropsMixin_ColumnSchemaProperty struct {
	// The data type of the column schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnschema.html#cfn-quicksight-template-columnschema-datatype
	//
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// The geographic role of the column schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnschema.html#cfn-quicksight-template-columnschema-geographicrole
	//
	GeographicRole *string `field:"optional" json:"geographicRole" yaml:"geographicRole"`
	// The name of the column schema.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-columnschema.html#cfn-quicksight-template-columnschema-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


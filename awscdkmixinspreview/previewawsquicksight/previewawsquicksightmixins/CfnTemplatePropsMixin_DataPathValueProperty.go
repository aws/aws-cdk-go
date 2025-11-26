package previewawsquicksightmixins


// The data path that needs to be sorted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataPathValueProperty := &DataPathValueProperty{
//   	DataPathType: &DataPathTypeProperty{
//   		PivotTableDataPathType: jsii.String("pivotTableDataPathType"),
//   	},
//   	FieldId: jsii.String("fieldId"),
//   	FieldValue: jsii.String("fieldValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathvalue.html
//
type CfnTemplatePropsMixin_DataPathValueProperty struct {
	// The type configuration of the field.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathvalue.html#cfn-quicksight-template-datapathvalue-datapathtype
	//
	DataPathType interface{} `field:"optional" json:"dataPathType" yaml:"dataPathType"`
	// The field ID of the field that needs to be sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathvalue.html#cfn-quicksight-template-datapathvalue-fieldid
	//
	FieldId *string `field:"optional" json:"fieldId" yaml:"fieldId"`
	// The actual value of the field that needs to be sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-datapathvalue.html#cfn-quicksight-template-datapathvalue-fieldvalue
	//
	FieldValue *string `field:"optional" json:"fieldValue" yaml:"fieldValue"`
}


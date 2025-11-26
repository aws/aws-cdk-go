package previewawsquicksightmixins


// Represents a string value used in filter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetStringFilterValueProperty := &DataSetStringFilterValueProperty{
//   	StaticValue: jsii.String("staticValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltervalue.html
//
type CfnDataSetPropsMixin_DataSetStringFilterValueProperty struct {
	// A static string value used for filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetstringfiltervalue.html#cfn-quicksight-dataset-datasetstringfiltervalue-staticvalue
	//
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
}


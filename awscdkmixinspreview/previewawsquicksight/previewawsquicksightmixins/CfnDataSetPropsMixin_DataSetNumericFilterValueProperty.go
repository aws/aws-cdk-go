package previewawsquicksightmixins


// Represents a numeric value used in filter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetNumericFilterValueProperty := &DataSetNumericFilterValueProperty{
//   	StaticValue: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltervalue.html
//
type CfnDataSetPropsMixin_DataSetNumericFilterValueProperty struct {
	// A static numeric value used for filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetnumericfiltervalue.html#cfn-quicksight-dataset-datasetnumericfiltervalue-staticvalue
	//
	StaticValue *float64 `field:"optional" json:"staticValue" yaml:"staticValue"`
}


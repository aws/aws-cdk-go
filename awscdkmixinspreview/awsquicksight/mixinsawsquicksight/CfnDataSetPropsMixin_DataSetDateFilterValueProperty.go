package mixinsawsquicksight


// Represents a date value used in filter conditions.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSetDateFilterValueProperty := &DataSetDateFilterValueProperty{
//   	StaticValue: jsii.String("staticValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltervalue.html
//
type CfnDataSetPropsMixin_DataSetDateFilterValueProperty struct {
	// A static date value used for filtering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datasetdatefiltervalue.html#cfn-quicksight-dataset-datasetdatefiltervalue-staticvalue
	//
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
}


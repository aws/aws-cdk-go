package awsquicksight


// The data path that needs to be sorted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPathValueProperty := &DataPathValueProperty{
//   	FieldId: jsii.String("fieldId"),
//   	FieldValue: jsii.String("fieldValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathvalue.html
//
type CfnAnalysis_DataPathValueProperty struct {
	// The field ID of the field that needs to be sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathvalue.html#cfn-quicksight-analysis-datapathvalue-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// The actual value of the field that needs to be sorted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-datapathvalue.html#cfn-quicksight-analysis-datapathvalue-fieldvalue
	//
	FieldValue *string `field:"required" json:"fieldValue" yaml:"fieldValue"`
}


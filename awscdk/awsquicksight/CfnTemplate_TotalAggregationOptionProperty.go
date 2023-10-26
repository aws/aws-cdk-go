package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   totalAggregationOptionProperty := &TotalAggregationOptionProperty{
//   	FieldId: jsii.String("fieldId"),
//   	TotalAggregationFunction: &TotalAggregationFunctionProperty{
//   		SimpleTotalAggregationFunction: jsii.String("simpleTotalAggregationFunction"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totalaggregationoption.html
//
type CfnTemplate_TotalAggregationOptionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totalaggregationoption.html#cfn-quicksight-template-totalaggregationoption-fieldid
	//
	FieldId *string `field:"required" json:"fieldId" yaml:"fieldId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totalaggregationoption.html#cfn-quicksight-template-totalaggregationoption-totalaggregationfunction
	//
	TotalAggregationFunction interface{} `field:"required" json:"totalAggregationFunction" yaml:"totalAggregationFunction"`
}


package awsquicksight


// An aggregation function that aggregates the total values of a measure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   totalAggregationFunctionProperty := &TotalAggregationFunctionProperty{
//   	SimpleTotalAggregationFunction: jsii.String("simpleTotalAggregationFunction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totalaggregationfunction.html
//
type CfnTemplate_TotalAggregationFunctionProperty struct {
	// A built in aggregation function for total values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-totalaggregationfunction.html#cfn-quicksight-template-totalaggregationfunction-simpletotalaggregationfunction
	//
	SimpleTotalAggregationFunction *string `field:"optional" json:"simpleTotalAggregationFunction" yaml:"simpleTotalAggregationFunction"`
}


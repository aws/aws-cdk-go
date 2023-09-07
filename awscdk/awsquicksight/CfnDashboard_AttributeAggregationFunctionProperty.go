package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributeAggregationFunctionProperty := &AttributeAggregationFunctionProperty{
//   	SimpleAttributeAggregation: jsii.String("simpleAttributeAggregation"),
//   	ValueForMultipleValues: jsii.String("valueForMultipleValues"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-attributeaggregationfunction.html
//
type CfnDashboard_AttributeAggregationFunctionProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-attributeaggregationfunction.html#cfn-quicksight-dashboard-attributeaggregationfunction-simpleattributeaggregation
	//
	SimpleAttributeAggregation *string `field:"optional" json:"simpleAttributeAggregation" yaml:"simpleAttributeAggregation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dashboard-attributeaggregationfunction.html#cfn-quicksight-dashboard-attributeaggregationfunction-valueformultiplevalues
	//
	ValueForMultipleValues *string `field:"optional" json:"valueForMultipleValues" yaml:"valueForMultipleValues"`
}


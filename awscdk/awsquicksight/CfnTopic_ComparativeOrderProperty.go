package awsquicksight


// The order in which data is displayed for the column when it's used in a comparative context.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   comparativeOrderProperty := &ComparativeOrderProperty{
//   	SpecifedOrder: []*string{
//   		jsii.String("specifedOrder"),
//   	},
//   	TreatUndefinedSpecifiedValues: jsii.String("treatUndefinedSpecifiedValues"),
//   	UseOrdering: jsii.String("useOrdering"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-comparativeorder.html
//
type CfnTopic_ComparativeOrderProperty struct {
	// The list of columns to be used in the ordering.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-comparativeorder.html#cfn-quicksight-topic-comparativeorder-specifedorder
	//
	SpecifedOrder *[]*string `field:"optional" json:"specifedOrder" yaml:"specifedOrder"`
	// The treat of undefined specified values.
	//
	// Valid values for this structure are `LEAST` and `MOST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-comparativeorder.html#cfn-quicksight-topic-comparativeorder-treatundefinedspecifiedvalues
	//
	TreatUndefinedSpecifiedValues *string `field:"optional" json:"treatUndefinedSpecifiedValues" yaml:"treatUndefinedSpecifiedValues"`
	// The ordering type for a column.
	//
	// Valid values for this structure are `GREATER_IS_BETTER` , `LESSER_IS_BETTER` and `SPECIFIED` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-topic-comparativeorder.html#cfn-quicksight-topic-comparativeorder-useordering
	//
	UseOrdering *string `field:"optional" json:"useOrdering" yaml:"useOrdering"`
}


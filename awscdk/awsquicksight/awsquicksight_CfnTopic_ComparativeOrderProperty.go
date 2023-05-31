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
type CfnTopic_ComparativeOrderProperty struct {
	// The list of columns to be used in the ordering.
	SpecifedOrder *[]*string `field:"optional" json:"specifedOrder" yaml:"specifedOrder"`
	// The treat of undefined specified values.
	//
	// Valid values for this structure are `LEAST` and `MOST` .
	TreatUndefinedSpecifiedValues *string `field:"optional" json:"treatUndefinedSpecifiedValues" yaml:"treatUndefinedSpecifiedValues"`
	// The ordering type for a column.
	//
	// Valid values for this structure are `GREATER_IS_BETTER` , `LESSER_IS_BETTER` and `SPECIFIED` .
	UseOrdering *string `field:"optional" json:"useOrdering" yaml:"useOrdering"`
}


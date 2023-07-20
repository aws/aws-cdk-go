package awsmacie


// Specifies a condition that defines the property, operator, and one or more values to use in a findings filter.
//
// A *findings filter* , also referred to as a *filter rule* , is a set of custom criteria that specifies which findings to include or exclude from the results of a query for findings. You can also configure a findings filter to suppress (automatically archive) findings that match the filter's criteria. For more information, see [Filtering findings](https://docs.aws.amazon.com/macie/latest/user/findings-filter-overview.html) in the *Amazon Macie User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   criterionAdditionalPropertiesProperty := &CriterionAdditionalPropertiesProperty{
//   	Eq: []*string{
//   		jsii.String("eq"),
//   	},
//   	Gt: jsii.Number(123),
//   	Gte: jsii.Number(123),
//   	Lt: jsii.Number(123),
//   	Lte: jsii.Number(123),
//   	Neq: []*string{
//   		jsii.String("neq"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html
//
type CfnFindingsFilter_CriterionAdditionalPropertiesProperty struct {
	// The value for the specified property matches (equals) the specified value.
	//
	// If you specify multiple values, Amazon Macie uses OR logic to join the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-eq
	//
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// The value for the specified property is greater than the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-gt
	//
	Gt *float64 `field:"optional" json:"gt" yaml:"gt"`
	// The value for the specified property is greater than or equal to the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-gte
	//
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// The value for the specified property is less than the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-lt
	//
	Lt *float64 `field:"optional" json:"lt" yaml:"lt"`
	// The value for the specified property is less than or equal to the specified value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-lte
	//
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
	// The value for the specified property doesn't match (doesn't equal) the specified value.
	//
	// If you specify multiple values, Amazon Macie uses OR logic to join the values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-macie-findingsfilter-criterionadditionalproperties.html#cfn-macie-findingsfilter-criterionadditionalproperties-neq
	//
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}


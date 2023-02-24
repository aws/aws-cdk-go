package awsmacie


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
type CfnFindingsFilter_CriterionAdditionalPropertiesProperty struct {
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.eq`.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.gt`.
	Gt *float64 `field:"optional" json:"gt" yaml:"gt"`
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.gte`.
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.lt`.
	Lt *float64 `field:"optional" json:"lt" yaml:"lt"`
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.lte`.
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
	// `CfnFindingsFilter.CriterionAdditionalPropertiesProperty.neq`.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}


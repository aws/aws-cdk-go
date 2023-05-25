package awsquicksight


// A list of filter configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterListConfigurationProperty := &FilterListConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//
//   	// the properties below are optional
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
type CfnTemplate_FilterListConfigurationProperty struct {
	// The match operator that is used to determine if a filter should be applied.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// The list of category values for the filter.
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


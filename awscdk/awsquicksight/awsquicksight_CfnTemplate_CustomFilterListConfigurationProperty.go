package awsquicksight


// A list of custom filter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFilterListConfigurationProperty := &CustomFilterListConfigurationProperty{
//   	MatchOperator: jsii.String("matchOperator"),
//   	NullOption: jsii.String("nullOption"),
//
//   	// the properties below are optional
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
type CfnTemplate_CustomFilterListConfigurationProperty struct {
	// The match operator that is used to determine if a filter should be applied.
	MatchOperator *string `field:"required" json:"matchOperator" yaml:"matchOperator"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	NullOption *string `field:"required" json:"nullOption" yaml:"nullOption"`
	// The list of category values for the filter.
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


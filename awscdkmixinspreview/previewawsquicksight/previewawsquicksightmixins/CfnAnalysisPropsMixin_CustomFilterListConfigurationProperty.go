package previewawsquicksightmixins


// A list of custom filter values.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customFilterListConfigurationProperty := &CustomFilterListConfigurationProperty{
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	MatchOperator: jsii.String("matchOperator"),
//   	NullOption: jsii.String("nullOption"),
//   	SelectAllOptions: jsii.String("selectAllOptions"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customfilterlistconfiguration.html
//
type CfnAnalysisPropsMixin_CustomFilterListConfigurationProperty struct {
	// The list of category values for the filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customfilterlistconfiguration.html#cfn-quicksight-analysis-customfilterlistconfiguration-categoryvalues
	//
	CategoryValues *[]*string `field:"optional" json:"categoryValues" yaml:"categoryValues"`
	// The match operator that is used to determine if a filter should be applied.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customfilterlistconfiguration.html#cfn-quicksight-analysis-customfilterlistconfiguration-matchoperator
	//
	MatchOperator *string `field:"optional" json:"matchOperator" yaml:"matchOperator"`
	// This option determines how null values should be treated when filtering data.
	//
	// - `ALL_VALUES` : Include null values in filtered results.
	// - `NULLS_ONLY` : Only include null values in filtered results.
	// - `NON_NULLS_ONLY` : Exclude null values from filtered results.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customfilterlistconfiguration.html#cfn-quicksight-analysis-customfilterlistconfiguration-nulloption
	//
	NullOption *string `field:"optional" json:"nullOption" yaml:"nullOption"`
	// Select all of the values. Null is not the assigned value of select all.
	//
	// - `FILTER_ALL_VALUES`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-customfilterlistconfiguration.html#cfn-quicksight-analysis-customfilterlistconfiguration-selectalloptions
	//
	SelectAllOptions *string `field:"optional" json:"selectAllOptions" yaml:"selectAllOptions"`
}


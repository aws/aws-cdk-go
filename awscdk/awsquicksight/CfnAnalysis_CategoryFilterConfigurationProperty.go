package awsquicksight


// The configuration for a `CategoryFilter` .
//
// This is a union type structure. For this structure to be valid, only one of the attributes can be defined.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoryFilterConfigurationProperty := &CategoryFilterConfigurationProperty{
//   	CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   		MatchOperator: jsii.String("matchOperator"),
//   		NullOption: jsii.String("nullOption"),
//
//   		// the properties below are optional
//   		CategoryValue: jsii.String("categoryValue"),
//   		ParameterName: jsii.String("parameterName"),
//   		SelectAllOptions: jsii.String("selectAllOptions"),
//   	},
//   	CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   		MatchOperator: jsii.String("matchOperator"),
//   		NullOption: jsii.String("nullOption"),
//
//   		// the properties below are optional
//   		CategoryValues: []*string{
//   			jsii.String("categoryValues"),
//   		},
//   		SelectAllOptions: jsii.String("selectAllOptions"),
//   	},
//   	FilterListConfiguration: &FilterListConfigurationProperty{
//   		MatchOperator: jsii.String("matchOperator"),
//
//   		// the properties below are optional
//   		CategoryValues: []*string{
//   			jsii.String("categoryValues"),
//   		},
//   		SelectAllOptions: jsii.String("selectAllOptions"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-categoryfilterconfiguration.html
//
type CfnAnalysis_CategoryFilterConfigurationProperty struct {
	// A custom filter that filters based on a single value.
	//
	// This filter can be partially matched.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-categoryfilterconfiguration.html#cfn-quicksight-analysis-categoryfilterconfiguration-customfilterconfiguration
	//
	CustomFilterConfiguration interface{} `field:"optional" json:"customFilterConfiguration" yaml:"customFilterConfiguration"`
	// A list of custom filter values.
	//
	// In the Amazon QuickSight console, this filter type is called a custom filter list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-categoryfilterconfiguration.html#cfn-quicksight-analysis-categoryfilterconfiguration-customfilterlistconfiguration
	//
	CustomFilterListConfiguration interface{} `field:"optional" json:"customFilterListConfiguration" yaml:"customFilterListConfiguration"`
	// A list of filter configurations.
	//
	// In the Amazon QuickSight console, this filter type is called a filter list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-analysis-categoryfilterconfiguration.html#cfn-quicksight-analysis-categoryfilterconfiguration-filterlistconfiguration
	//
	FilterListConfiguration interface{} `field:"optional" json:"filterListConfiguration" yaml:"filterListConfiguration"`
}


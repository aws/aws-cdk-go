package awsquicksight


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
type CfnAnalysis_CategoryFilterConfigurationProperty struct {
	// `CfnAnalysis.CategoryFilterConfigurationProperty.CustomFilterConfiguration`.
	CustomFilterConfiguration interface{} `field:"optional" json:"customFilterConfiguration" yaml:"customFilterConfiguration"`
	// `CfnAnalysis.CategoryFilterConfigurationProperty.CustomFilterListConfiguration`.
	CustomFilterListConfiguration interface{} `field:"optional" json:"customFilterListConfiguration" yaml:"customFilterListConfiguration"`
	// `CfnAnalysis.CategoryFilterConfigurationProperty.FilterListConfiguration`.
	FilterListConfiguration interface{} `field:"optional" json:"filterListConfiguration" yaml:"filterListConfiguration"`
}


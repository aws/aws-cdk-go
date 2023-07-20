package awsquicksight


// A `CategoryFilter` filters text values.
//
// For more information, see [Adding text filters](https://docs.aws.amazon.com/quicksight/latest/user/add-a-text-filter-data-prep.html) in the *Amazon QuickSight User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoryFilterProperty := &CategoryFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	Configuration: &CategoryFilterConfigurationProperty{
//   		CustomFilterConfiguration: &CustomFilterConfigurationProperty{
//   			MatchOperator: jsii.String("matchOperator"),
//   			NullOption: jsii.String("nullOption"),
//
//   			// the properties below are optional
//   			CategoryValue: jsii.String("categoryValue"),
//   			ParameterName: jsii.String("parameterName"),
//   			SelectAllOptions: jsii.String("selectAllOptions"),
//   		},
//   		CustomFilterListConfiguration: &CustomFilterListConfigurationProperty{
//   			MatchOperator: jsii.String("matchOperator"),
//   			NullOption: jsii.String("nullOption"),
//
//   			// the properties below are optional
//   			CategoryValues: []*string{
//   				jsii.String("categoryValues"),
//   			},
//   			SelectAllOptions: jsii.String("selectAllOptions"),
//   		},
//   		FilterListConfiguration: &FilterListConfigurationProperty{
//   			MatchOperator: jsii.String("matchOperator"),
//
//   			// the properties below are optional
//   			CategoryValues: []*string{
//   				jsii.String("categoryValues"),
//   			},
//   			SelectAllOptions: jsii.String("selectAllOptions"),
//   		},
//   	},
//   	FilterId: jsii.String("filterId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html
//
type CfnTemplate_CategoryFilterProperty struct {
	// The column that the filter is applied to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-column
	//
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The configuration for a `CategoryFilter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-configuration
	//
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// An identifier that uniquely identifies a filter within a dashboard, analysis, or template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-categoryfilter.html#cfn-quicksight-template-categoryfilter-filterid
	//
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
}


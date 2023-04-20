package awsquicksight


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
type CfnAnalysis_CategoryFilterProperty struct {
	// `CfnAnalysis.CategoryFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnAnalysis.CategoryFilterProperty.Configuration`.
	Configuration interface{} `field:"required" json:"configuration" yaml:"configuration"`
	// `CfnAnalysis.CategoryFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
}


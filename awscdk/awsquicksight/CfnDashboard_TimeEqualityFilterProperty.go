package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeEqualityFilterProperty := &TimeEqualityFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	FilterId: jsii.String("filterId"),
//
//   	// the properties below are optional
//   	ParameterName: jsii.String("parameterName"),
//   	TimeGranularity: jsii.String("timeGranularity"),
//   	Value: jsii.String("value"),
//   }
//
type CfnDashboard_TimeEqualityFilterProperty struct {
	// `CfnDashboard.TimeEqualityFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.TimeEqualityFilterProperty.FilterId`.
	FilterId *string `field:"required" json:"filterId" yaml:"filterId"`
	// `CfnDashboard.TimeEqualityFilterProperty.ParameterName`.
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// `CfnDashboard.TimeEqualityFilterProperty.TimeGranularity`.
	TimeGranularity *string `field:"optional" json:"timeGranularity" yaml:"timeGranularity"`
	// `CfnDashboard.TimeEqualityFilterProperty.Value`.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


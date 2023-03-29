package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   numericEqualityDrillDownFilterProperty := &NumericEqualityDrillDownFilterProperty{
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   	Value: jsii.Number(123),
//   }
//
type CfnDashboard_NumericEqualityDrillDownFilterProperty struct {
	// `CfnDashboard.NumericEqualityDrillDownFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// `CfnDashboard.NumericEqualityDrillDownFilterProperty.Value`.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


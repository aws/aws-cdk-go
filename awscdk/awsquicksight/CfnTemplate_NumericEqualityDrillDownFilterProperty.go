package awsquicksight


// The category drill down filter.
//
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
type CfnTemplate_NumericEqualityDrillDownFilterProperty struct {
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// The value of the double input numeric drill down filter.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}


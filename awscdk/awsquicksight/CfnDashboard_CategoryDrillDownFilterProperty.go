package awsquicksight


// The numeric equality type drill down filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   categoryDrillDownFilterProperty := &CategoryDrillDownFilterProperty{
//   	CategoryValues: []*string{
//   		jsii.String("categoryValues"),
//   	},
//   	Column: &ColumnIdentifierProperty{
//   		ColumnName: jsii.String("columnName"),
//   		DataSetIdentifier: jsii.String("dataSetIdentifier"),
//   	},
//   }
//
type CfnDashboard_CategoryDrillDownFilterProperty struct {
	// A list of the string inputs that are the values of the category drill down filter.
	CategoryValues *[]*string `field:"required" json:"categoryValues" yaml:"categoryValues"`
	// The column that the filter is applied to.
	Column interface{} `field:"required" json:"column" yaml:"column"`
}


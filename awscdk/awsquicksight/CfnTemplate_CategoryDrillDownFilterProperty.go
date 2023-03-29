package awsquicksight


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
type CfnTemplate_CategoryDrillDownFilterProperty struct {
	// `CfnTemplate.CategoryDrillDownFilterProperty.CategoryValues`.
	CategoryValues *[]*string `field:"required" json:"categoryValues" yaml:"categoryValues"`
	// `CfnTemplate.CategoryDrillDownFilterProperty.Column`.
	Column interface{} `field:"required" json:"column" yaml:"column"`
}


package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lookbackWindowProperty := &LookbackWindowProperty{
//   	ColumnName: jsii.String("columnName"),
//   	Size: jsii.Number(123),
//   	SizeUnit: jsii.String("sizeUnit"),
//   }
//
type CfnDataSet_LookbackWindowProperty struct {
	// `CfnDataSet.LookbackWindowProperty.ColumnName`.
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// `CfnDataSet.LookbackWindowProperty.Size`.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// `CfnDataSet.LookbackWindowProperty.SizeUnit`.
	SizeUnit *string `field:"optional" json:"sizeUnit" yaml:"sizeUnit"`
}


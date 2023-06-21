package awsquicksight


// The lookback window setup of an incremental refresh configuration.
//
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
	// The name of the lookback window column.
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// The lookback window column size.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// The size unit that is used for the lookback window column.
	//
	// Valid values for this structure are `HOUR` , `DAY` , and `WEEK` .
	SizeUnit *string `field:"optional" json:"sizeUnit" yaml:"sizeUnit"`
}


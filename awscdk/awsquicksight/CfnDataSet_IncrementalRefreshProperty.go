package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   incrementalRefreshProperty := &IncrementalRefreshProperty{
//   	LookbackWindow: &LookbackWindowProperty{
//   		ColumnName: jsii.String("columnName"),
//   		Size: jsii.Number(123),
//   		SizeUnit: jsii.String("sizeUnit"),
//   	},
//   }
//
type CfnDataSet_IncrementalRefreshProperty struct {
	// `CfnDataSet.IncrementalRefreshProperty.LookbackWindow`.
	LookbackWindow interface{} `field:"optional" json:"lookbackWindow" yaml:"lookbackWindow"`
}


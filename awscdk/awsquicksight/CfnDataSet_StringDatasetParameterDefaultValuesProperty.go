package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stringDatasetParameterDefaultValuesProperty := &StringDatasetParameterDefaultValuesProperty{
//   	StaticValues: []*string{
//   		jsii.String("staticValues"),
//   	},
//   }
//
type CfnDataSet_StringDatasetParameterDefaultValuesProperty struct {
	// `CfnDataSet.StringDatasetParameterDefaultValuesProperty.StaticValues`.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


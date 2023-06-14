package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dateTimeDatasetParameterDefaultValuesProperty := &DateTimeDatasetParameterDefaultValuesProperty{
//   	StaticValues: []*string{
//   		jsii.String("staticValues"),
//   	},
//   }
//
type CfnDataSet_DateTimeDatasetParameterDefaultValuesProperty struct {
	// `CfnDataSet.DateTimeDatasetParameterDefaultValuesProperty.StaticValues`.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


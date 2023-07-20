package awsquicksight


// <p>List of default values defined for a given string date time parameter type.
//
// Currently only static values are supported.</p>
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameterdefaultvalues.html
//
type CfnDataSet_DateTimeDatasetParameterDefaultValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-datetimedatasetparameterdefaultvalues.html#cfn-quicksight-dataset-datetimedatasetparameterdefaultvalues-staticvalues
	//
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


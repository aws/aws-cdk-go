package awsquicksight


// <p>List of default values defined for a given string dataset parameter type.
//
// Currently only static values are supported.</p>
//
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameterdefaultvalues.html
//
type CfnDataSet_StringDatasetParameterDefaultValuesProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-dataset-stringdatasetparameterdefaultvalues.html#cfn-quicksight-dataset-stringdatasetparameterdefaultvalues-staticvalues
	//
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


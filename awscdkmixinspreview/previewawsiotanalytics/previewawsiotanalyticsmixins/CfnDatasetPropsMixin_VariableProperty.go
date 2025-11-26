package previewawsiotanalyticsmixins


// An instance of a variable to be passed to the `containerAction` execution.
//
// Each variable must have a name and a value given by one of `stringValue` , `datasetContentVersionValue` , or `outputFileUriValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   variableProperty := &VariableProperty{
//   	DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   		DatasetName: jsii.String("datasetName"),
//   	},
//   	DoubleValue: jsii.Number(123),
//   	OutputFileUriValue: &OutputFileUriValueProperty{
//   		FileName: jsii.String("fileName"),
//   	},
//   	StringValue: jsii.String("stringValue"),
//   	VariableName: jsii.String("variableName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html
//
type CfnDatasetPropsMixin_VariableProperty struct {
	// The value of the variable as a structure that specifies a dataset content version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html#cfn-iotanalytics-dataset-variable-datasetcontentversionvalue
	//
	DatasetContentVersionValue interface{} `field:"optional" json:"datasetContentVersionValue" yaml:"datasetContentVersionValue"`
	// The value of the variable as a double (numeric).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html#cfn-iotanalytics-dataset-variable-doublevalue
	//
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value of the variable as a structure that specifies an output file URI.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html#cfn-iotanalytics-dataset-variable-outputfileurivalue
	//
	OutputFileUriValue interface{} `field:"optional" json:"outputFileUriValue" yaml:"outputFileUriValue"`
	// The value of the variable as a string.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html#cfn-iotanalytics-dataset-variable-stringvalue
	//
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
	// The name of the variable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-variable.html#cfn-iotanalytics-dataset-variable-variablename
	//
	VariableName *string `field:"optional" json:"variableName" yaml:"variableName"`
}


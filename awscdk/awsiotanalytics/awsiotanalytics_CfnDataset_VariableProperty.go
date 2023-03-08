package awsiotanalytics


// An instance of a variable to be passed to the `containerAction` execution.
//
// Each variable must have a name and a value given by one of `stringValue` , `datasetContentVersionValue` , or `outputFileUriValue` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   variableProperty := &variableProperty{
//   	variableName: jsii.String("variableName"),
//
//   	// the properties below are optional
//   	datasetContentVersionValue: &datasetContentVersionValueProperty{
//   		datasetName: jsii.String("datasetName"),
//   	},
//   	doubleValue: jsii.Number(123),
//   	outputFileUriValue: &outputFileUriValueProperty{
//   		fileName: jsii.String("fileName"),
//   	},
//   	stringValue: jsii.String("stringValue"),
//   }
//
type CfnDataset_VariableProperty struct {
	// The name of the variable.
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// The value of the variable as a structure that specifies a dataset content version.
	DatasetContentVersionValue interface{} `field:"optional" json:"datasetContentVersionValue" yaml:"datasetContentVersionValue"`
	// The value of the variable as a double (numeric).
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// The value of the variable as a structure that specifies an output file URI.
	OutputFileUriValue interface{} `field:"optional" json:"outputFileUriValue" yaml:"outputFileUriValue"`
	// The value of the variable as a string.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


package previewawsiotanalyticsmixins


// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   containerActionProperty := &ContainerActionProperty{
//   	ExecutionRoleArn: jsii.String("executionRoleArn"),
//   	Image: jsii.String("image"),
//   	ResourceConfiguration: &ResourceConfigurationProperty{
//   		ComputeType: jsii.String("computeType"),
//   		VolumeSizeInGb: jsii.Number(123),
//   	},
//   	Variables: []interface{}{
//   		&VariableProperty{
//   			DatasetContentVersionValue: &DatasetContentVersionValueProperty{
//   				DatasetName: jsii.String("datasetName"),
//   			},
//   			DoubleValue: jsii.Number(123),
//   			OutputFileUriValue: &OutputFileUriValueProperty{
//   				FileName: jsii.String("fileName"),
//   			},
//   			StringValue: jsii.String("stringValue"),
//   			VariableName: jsii.String("variableName"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html
//
type CfnDatasetPropsMixin_ContainerActionProperty struct {
	// The ARN of the role which gives permission to the system to access needed resources in order to run the "containerAction".
	//
	// This includes, at minimum, permission to retrieve the data set contents which are the input to the containerized application.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-executionrolearn
	//
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The ARN of the Docker container stored in your account.
	//
	// The Docker container contains an application and needed support libraries and is used to generate data set contents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-image
	//
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Configuration of the resource which executes the "containerAction".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-resourceconfiguration
	//
	ResourceConfiguration interface{} `field:"optional" json:"resourceConfiguration" yaml:"resourceConfiguration"`
	// The values of variables used within the context of the execution of the containerized application (basically, parameters passed to the application).
	//
	// Each variable must have a name and a value given by one of "stringValue", "datasetContentVersionValue", or "outputFileUriValue".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotanalytics-dataset-containeraction.html#cfn-iotanalytics-dataset-containeraction-variables
	//
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}


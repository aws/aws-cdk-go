package awsiotanalytics


// Information needed to run the "containerAction" to produce data set contents.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   containerActionProperty := &containerActionProperty{
//   	executionRoleArn: jsii.String("executionRoleArn"),
//   	image: jsii.String("image"),
//   	resourceConfiguration: &resourceConfigurationProperty{
//   		computeType: jsii.String("computeType"),
//   		volumeSizeInGb: jsii.Number(123),
//   	},
//
//   	// the properties below are optional
//   	variables: []interface{}{
//   		&variableProperty{
//   			variableName: jsii.String("variableName"),
//
//   			// the properties below are optional
//   			datasetContentVersionValue: &datasetContentVersionValueProperty{
//   				datasetName: jsii.String("datasetName"),
//   			},
//   			doubleValue: jsii.Number(123),
//   			outputFileUriValue: &outputFileUriValueProperty{
//   				fileName: jsii.String("fileName"),
//   			},
//   			stringValue: jsii.String("stringValue"),
//   		},
//   	},
//   }
//
type CfnDataset_ContainerActionProperty struct {
	// The ARN of the role which gives permission to the system to access needed resources in order to run the "containerAction".
	//
	// This includes, at minimum, permission to retrieve the data set contents which are the input to the containerized application.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// The ARN of the Docker container stored in your account.
	//
	// The Docker container contains an application and needed support libraries and is used to generate data set contents.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Configuration of the resource which executes the "containerAction".
	ResourceConfiguration interface{} `field:"required" json:"resourceConfiguration" yaml:"resourceConfiguration"`
	// The values of variables used within the context of the execution of the containerized application (basically, parameters passed to the application).
	//
	// Each variable must have a name and a value given by one of "stringValue", "datasetContentVersionValue", or "outputFileUriValue".
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}


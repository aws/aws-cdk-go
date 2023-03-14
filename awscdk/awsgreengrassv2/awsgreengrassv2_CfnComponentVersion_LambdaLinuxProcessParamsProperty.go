package awsgreengrassv2


// Contains parameters for a Linux process that contains an AWS Lambda function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaLinuxProcessParamsProperty := &lambdaLinuxProcessParamsProperty{
//   	containerParams: &lambdaContainerParamsProperty{
//   		devices: []interface{}{
//   			&lambdaDeviceMountProperty{
//   				addGroupOwner: jsii.Boolean(false),
//   				path: jsii.String("path"),
//   				permission: jsii.String("permission"),
//   			},
//   		},
//   		memorySizeInKb: jsii.Number(123),
//   		mountRoSysfs: jsii.Boolean(false),
//   		volumes: []interface{}{
//   			&lambdaVolumeMountProperty{
//   				addGroupOwner: jsii.Boolean(false),
//   				destinationPath: jsii.String("destinationPath"),
//   				permission: jsii.String("permission"),
//   				sourcePath: jsii.String("sourcePath"),
//   			},
//   		},
//   	},
//   	isolationMode: jsii.String("isolationMode"),
//   }
//
type CfnComponentVersion_LambdaLinuxProcessParamsProperty struct {
	// The parameters for the container in which the Lambda function runs.
	ContainerParams interface{} `field:"optional" json:"containerParams" yaml:"containerParams"`
	// The isolation mode for the process that contains the Lambda function.
	//
	// The process can run in an isolated runtime environment inside the AWS IoT Greengrass container, or as a regular process outside any container.
	//
	// Default: `GreengrassContainer`.
	IsolationMode *string `field:"optional" json:"isolationMode" yaml:"isolationMode"`
}


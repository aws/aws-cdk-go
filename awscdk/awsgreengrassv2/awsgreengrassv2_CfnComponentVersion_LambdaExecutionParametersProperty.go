package awsgreengrassv2


// Contains parameters for a Lambda function that runs on AWS IoT Greengrass .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaExecutionParametersProperty := &lambdaExecutionParametersProperty{
//   	environmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	eventSources: []interface{}{
//   		&lambdaEventSourceProperty{
//   			topic: jsii.String("topic"),
//   			type: jsii.String("type"),
//   		},
//   	},
//   	execArgs: []*string{
//   		jsii.String("execArgs"),
//   	},
//   	inputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   	linuxProcessParams: &lambdaLinuxProcessParamsProperty{
//   		containerParams: &lambdaContainerParamsProperty{
//   			devices: []interface{}{
//   				&lambdaDeviceMountProperty{
//   					addGroupOwner: jsii.Boolean(false),
//   					path: jsii.String("path"),
//   					permission: jsii.String("permission"),
//   				},
//   			},
//   			memorySizeInKb: jsii.Number(123),
//   			mountRoSysfs: jsii.Boolean(false),
//   			volumes: []interface{}{
//   				&lambdaVolumeMountProperty{
//   					addGroupOwner: jsii.Boolean(false),
//   					destinationPath: jsii.String("destinationPath"),
//   					permission: jsii.String("permission"),
//   					sourcePath: jsii.String("sourcePath"),
//   				},
//   			},
//   		},
//   		isolationMode: jsii.String("isolationMode"),
//   	},
//   	maxIdleTimeInSeconds: jsii.Number(123),
//   	maxInstancesCount: jsii.Number(123),
//   	maxQueueSize: jsii.Number(123),
//   	pinned: jsii.Boolean(false),
//   	statusTimeoutInSeconds: jsii.Number(123),
//   	timeoutInSeconds: jsii.Number(123),
//   }
//
type CfnComponentVersion_LambdaExecutionParametersProperty struct {
	// The map of environment variables that are available to the Lambda function when it runs.
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The list of event sources to which to subscribe to receive work messages.
	//
	// The Lambda function runs when it receives a message from an event source. You can subscribe this function to local publish/subscribe messages and AWS IoT Core MQTT messages.
	EventSources interface{} `field:"optional" json:"eventSources" yaml:"eventSources"`
	// The list of arguments to pass to the Lambda function when it runs.
	ExecArgs *[]*string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// The encoding type that the Lambda function supports.
	//
	// Default: `json`.
	InputPayloadEncodingType *string `field:"optional" json:"inputPayloadEncodingType" yaml:"inputPayloadEncodingType"`
	// The parameters for the Linux process that contains the Lambda function.
	LinuxProcessParams interface{} `field:"optional" json:"linuxProcessParams" yaml:"linuxProcessParams"`
	// The maximum amount of time in seconds that a non-pinned Lambda function can idle before the  software stops its process.
	MaxIdleTimeInSeconds *float64 `field:"optional" json:"maxIdleTimeInSeconds" yaml:"maxIdleTimeInSeconds"`
	// The maximum number of instances that a non-pinned Lambda function can run at the same time.
	MaxInstancesCount *float64 `field:"optional" json:"maxInstancesCount" yaml:"maxInstancesCount"`
	// The maximum size of the message queue for the Lambda function component.
	//
	// The Greengrass core device stores messages in a FIFO (first-in-first-out) queue until it can run the Lambda function to consume each message.
	MaxQueueSize *float64 `field:"optional" json:"maxQueueSize" yaml:"maxQueueSize"`
	// Whether or not the Lambda function is pinned, or long-lived.
	//
	// - A pinned Lambda function starts when the  starts and keeps running in its own container.
	// - A non-pinned Lambda function starts only when it receives a work item and exists after it idles for `maxIdleTimeInSeconds` . If the function has multiple work items, the  software creates multiple instances of the function.
	//
	// Default: `true`.
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// The interval in seconds at which a pinned (also known as long-lived) Lambda function component sends status updates to the Lambda manager component.
	StatusTimeoutInSeconds *float64 `field:"optional" json:"statusTimeoutInSeconds" yaml:"statusTimeoutInSeconds"`
	// The maximum amount of time in seconds that the Lambda function can process a work item.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


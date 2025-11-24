package mixinsawsgreengrassv2


// Contains parameters for a Lambda function that runs on AWS IoT Greengrass .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaExecutionParametersProperty := &LambdaExecutionParametersProperty{
//   	EnvironmentVariables: map[string]*string{
//   		"environmentVariablesKey": jsii.String("environmentVariables"),
//   	},
//   	EventSources: []interface{}{
//   		&LambdaEventSourceProperty{
//   			Topic: jsii.String("topic"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	ExecArgs: []*string{
//   		jsii.String("execArgs"),
//   	},
//   	InputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   	LinuxProcessParams: &LambdaLinuxProcessParamsProperty{
//   		ContainerParams: &LambdaContainerParamsProperty{
//   			Devices: []interface{}{
//   				&LambdaDeviceMountProperty{
//   					AddGroupOwner: jsii.Boolean(false),
//   					Path: jsii.String("path"),
//   					Permission: jsii.String("permission"),
//   				},
//   			},
//   			MemorySizeInKb: jsii.Number(123),
//   			MountRoSysfs: jsii.Boolean(false),
//   			Volumes: []interface{}{
//   				&LambdaVolumeMountProperty{
//   					AddGroupOwner: jsii.Boolean(false),
//   					DestinationPath: jsii.String("destinationPath"),
//   					Permission: jsii.String("permission"),
//   					SourcePath: jsii.String("sourcePath"),
//   				},
//   			},
//   		},
//   		IsolationMode: jsii.String("isolationMode"),
//   	},
//   	MaxIdleTimeInSeconds: jsii.Number(123),
//   	MaxInstancesCount: jsii.Number(123),
//   	MaxQueueSize: jsii.Number(123),
//   	Pinned: jsii.Boolean(false),
//   	StatusTimeoutInSeconds: jsii.Number(123),
//   	TimeoutInSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html
//
type CfnComponentVersionPropsMixin_LambdaExecutionParametersProperty struct {
	// The map of environment variables that are available to the Lambda function when it runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-environmentvariables
	//
	EnvironmentVariables interface{} `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// The list of event sources to which to subscribe to receive work messages.
	//
	// The Lambda function runs when it receives a message from an event source. You can subscribe this function to local publish/subscribe messages and AWS IoT Core MQTT messages.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-eventsources
	//
	EventSources interface{} `field:"optional" json:"eventSources" yaml:"eventSources"`
	// The list of arguments to pass to the Lambda function when it runs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-execargs
	//
	ExecArgs *[]*string `field:"optional" json:"execArgs" yaml:"execArgs"`
	// The encoding type that the Lambda function supports.
	//
	// Default: `json`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-inputpayloadencodingtype
	//
	InputPayloadEncodingType *string `field:"optional" json:"inputPayloadEncodingType" yaml:"inputPayloadEncodingType"`
	// The parameters for the Linux process that contains the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-linuxprocessparams
	//
	LinuxProcessParams interface{} `field:"optional" json:"linuxProcessParams" yaml:"linuxProcessParams"`
	// The maximum amount of time in seconds that a non-pinned Lambda function can idle before the AWS IoT Greengrass Core software stops its process.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-maxidletimeinseconds
	//
	MaxIdleTimeInSeconds *float64 `field:"optional" json:"maxIdleTimeInSeconds" yaml:"maxIdleTimeInSeconds"`
	// The maximum number of instances that a non-pinned Lambda function can run at the same time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-maxinstancescount
	//
	MaxInstancesCount *float64 `field:"optional" json:"maxInstancesCount" yaml:"maxInstancesCount"`
	// The maximum size of the message queue for the Lambda function component.
	//
	// The AWS IoT Greengrass core device stores messages in a FIFO (first-in-first-out) queue until it can run the Lambda function to consume each message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-maxqueuesize
	//
	MaxQueueSize *float64 `field:"optional" json:"maxQueueSize" yaml:"maxQueueSize"`
	// Whether or not the Lambda function is pinned, or long-lived.
	//
	// - A pinned Lambda function starts when the AWS IoT Greengrass Core starts and keeps running in its own container.
	// - A non-pinned Lambda function starts only when it receives a work item and exists after it idles for `maxIdleTimeInSeconds` . If the function has multiple work items, the AWS IoT Greengrass Core software creates multiple instances of the function.
	//
	// Default: `true`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-pinned
	//
	Pinned interface{} `field:"optional" json:"pinned" yaml:"pinned"`
	// The interval in seconds at which a pinned (also known as long-lived) Lambda function component sends status updates to the Lambda manager component.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-statustimeoutinseconds
	//
	StatusTimeoutInSeconds *float64 `field:"optional" json:"statusTimeoutInSeconds" yaml:"statusTimeoutInSeconds"`
	// The maximum amount of time in seconds that the Lambda function can process a work item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdaexecutionparameters.html#cfn-greengrassv2-componentversion-lambdaexecutionparameters-timeoutinseconds
	//
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


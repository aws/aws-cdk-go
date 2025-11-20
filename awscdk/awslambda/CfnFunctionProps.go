package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionProps := &CfnFunctionProps{
//   	Code: &CodeProperty{
//   		ImageUri: jsii.String("imageUri"),
//   		S3Bucket: jsii.String("s3Bucket"),
//   		S3Key: jsii.String("s3Key"),
//   		S3ObjectVersion: jsii.String("s3ObjectVersion"),
//   		SourceKmsKeyArn: jsii.String("sourceKmsKeyArn"),
//   		ZipFile: jsii.String("zipFile"),
//   	},
//   	Role: jsii.String("role"),
//
//   	// the properties below are optional
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	CapacityProviderConfig: &CapacityProviderConfigProperty{
//   		Ec2ManagedInstancesCapacityProviderConfig: &EC2ManagedInstancesCapacityProviderConfigProperty{
//   			CapacityProviderArn: jsii.String("capacityProviderArn"),
//
//   			// the properties below are optional
//   			ExecutionEnvironmentMaxConcurrency: jsii.Number(123),
//   			ExecutionEnvironmentMemoryGiBPerVCpu: jsii.Number(123),
//   		},
//   	},
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	DeadLetterConfig: &DeadLetterConfigProperty{
//   		TargetArn: jsii.String("targetArn"),
//   	},
//   	Description: jsii.String("description"),
//   	Environment: &EnvironmentProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	EphemeralStorage: &EphemeralStorageProperty{
//   		Size: jsii.Number(123),
//   	},
//   	FileSystemConfigs: []interface{}{
//   		&FileSystemConfigProperty{
//   			Arn: jsii.String("arn"),
//   			LocalMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
//   	FunctionScalingConfig: &FunctionScalingConfigProperty{
//   		MaxExecutionEnvironments: jsii.Number(123),
//   		MinExecutionEnvironments: jsii.Number(123),
//   	},
//   	Handler: jsii.String("handler"),
//   	ImageConfig: &ImageConfigProperty{
//   		Command: []*string{
//   			jsii.String("command"),
//   		},
//   		EntryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		WorkingDirectory: jsii.String("workingDirectory"),
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	LoggingConfig: &LoggingConfigProperty{
//   		ApplicationLogLevel: jsii.String("applicationLogLevel"),
//   		LogFormat: jsii.String("logFormat"),
//   		LogGroup: jsii.String("logGroup"),
//   		SystemLogLevel: jsii.String("systemLogLevel"),
//   	},
//   	MemorySize: jsii.Number(123),
//   	PackageType: jsii.String("packageType"),
//   	PublicAccessBlockConfig: &PublicAccessBlockConfigProperty{
//   		BlockPublicPolicy: jsii.Boolean(false),
//   		RestrictPublicResource: jsii.Boolean(false),
//   	},
//   	PublishToLatestPublished: jsii.Boolean(false),
//   	RecursiveLoop: jsii.String("recursiveLoop"),
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	Runtime: jsii.String("runtime"),
//   	RuntimeManagementConfig: &RuntimeManagementConfigProperty{
//   		UpdateRuntimeOn: jsii.String("updateRuntimeOn"),
//
//   		// the properties below are optional
//   		RuntimeVersionArn: jsii.String("runtimeVersionArn"),
//   	},
//   	SnapStart: &SnapStartProperty{
//   		ApplyOn: jsii.String("applyOn"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TenancyConfig: &TenancyConfigProperty{
//   		TenantIsolationMode: jsii.String("tenantIsolationMode"),
//   	},
//   	Timeout: jsii.Number(123),
//   	TracingConfig: &TracingConfigProperty{
//   		Mode: jsii.String("mode"),
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		Ipv6AllowedForDualStack: jsii.Boolean(false),
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html
//
type CfnFunctionProps struct {
	// The code for the function. You can define your function code in multiple ways:.
	//
	// - For .zip deployment packages, you can specify the Amazon S3 location of the .zip file in the `S3Bucket` , `S3Key` , and `S3ObjectVersion` properties.
	// - For .zip deployment packages, you can alternatively define the function code inline in the `ZipFile` property. This method works only for Node.js and Python functions.
	// - For container images, specify the URI of your container image in the Amazon ECR registry in the `ImageUri` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-code
	//
	Code interface{} `field:"required" json:"code" yaml:"code"`
	// The Amazon Resource Name (ARN) of the function's execution role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
	// The instruction set architecture that the function supports.
	//
	// Enter a string array with one of the valid values (arm64 or x86_64). The default value is `x86_64` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-capacityproviderconfig
	//
	CapacityProviderConfig interface{} `field:"optional" json:"capacityProviderConfig" yaml:"capacityProviderConfig"`
	// A unique Arn for CodeSigningConfig resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-codesigningconfigarn
	//
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// The dead-letter queue for failed asynchronous invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-deadletterconfig
	//
	DeadLetterConfig interface{} `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// A description of the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A function's environment variable settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// A function's ephemeral storage settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-ephemeralstorage
	//
	EphemeralStorage interface{} `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Connection settings for an Amazon EFS file system.
	//
	// To connect a function to a file system, a mount target must be available in every Availability Zone that your function connects to. If your template contains an AWS::EFS::MountTarget resource, you must also specify a DependsOn attribute to ensure that the mount target is created or updated before the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-filesystemconfigs
	//
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// The name of the Lambda function, up to 64 characters in length.
	//
	// If you don't specify a name, AWS CloudFormation generates one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// The scaling config of a version published into a capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-functionscalingconfig
	//
	FunctionScalingConfig interface{} `field:"optional" json:"functionScalingConfig" yaml:"functionScalingConfig"`
	// The name of the method within your code that Lambda calls to execute your function.
	//
	// The format includes the file name. It can also include namespaces and other qualifiers, depending on the runtime
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-handler
	//
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// Configuration values that override the container image Dockerfile settings.
	//
	// For more information, see [Container image settings](https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-imageconfig
	//
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables.
	//
	// If it's not provided, AWS Lambda uses a default service key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// A list of function layers to add to the function's execution environment.
	//
	// Specify each layer by its ARN, including the version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-layers
	//
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// The function's logging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-loggingconfig
	//
	LoggingConfig interface{} `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// The amount of memory that your function has access to.
	//
	// Increasing the function's memory also increases its CPU allocation. The default value is 128 MB. The value must be a multiple of 64 MB.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-memorysize
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// PackageType.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-packagetype
	//
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-publicaccessblockconfig
	//
	PublicAccessBlockConfig interface{} `field:"optional" json:"publicAccessBlockConfig" yaml:"publicAccessBlockConfig"`
	// A boolean indicating whether to publish $LATEST.PUBLISHED version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-publishtolatestpublished
	//
	PublishToLatestPublished interface{} `field:"optional" json:"publishToLatestPublished" yaml:"publishToLatestPublished"`
	// The function recursion configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-recursiveloop
	//
	RecursiveLoop *string `field:"optional" json:"recursiveLoop" yaml:"recursiveLoop"`
	// The number of simultaneous executions to reserve for the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-reservedconcurrentexecutions
	//
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// The identifier of the function's runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Sets the runtime management configuration for a function's version.
	//
	// For more information, see [Runtime updates](https://docs.aws.amazon.com/lambda/latest/dg/runtimes-update.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-runtimemanagementconfig
	//
	RuntimeManagementConfig interface{} `field:"optional" json:"runtimeManagementConfig" yaml:"runtimeManagementConfig"`
	// The function's SnapStart setting.
	//
	// When set to PublishedVersions, Lambda creates a snapshot of the execution environment when you publish a function version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-snapstart
	//
	SnapStart interface{} `field:"optional" json:"snapStart" yaml:"snapStart"`
	// A list of tags to apply to the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Controls how your Lambda function handles multi-tenant execution environments.
	//
	// Enable this setting to configure specific tenant isolation strategies for your function invocations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tenancyconfig
	//
	TenancyConfig interface{} `field:"optional" json:"tenancyConfig" yaml:"tenancyConfig"`
	// The amount of time that Lambda allows a function to run before stopping it.
	//
	// The default is 3 seconds. The maximum allowed value is 900 seconds.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// The function's AWS X-Ray tracing configuration.
	//
	// To sample and record incoming requests, set Mode to Active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-tracingconfig
	//
	TracingConfig interface{} `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
	// The VPC security groups and subnets that are attached to a Lambda function.
	//
	// When you connect a function to a VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the function's VPC configuration. The function can only access resources and the internet through that VPC.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-function.html#cfn-lambda-function-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


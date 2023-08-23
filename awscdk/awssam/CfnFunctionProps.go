package awssam


// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var assumeRolePolicyDocument interface{}
//
//   cfnFunctionProps := &CfnFunctionProps{
//   	Architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	AssumeRolePolicyDocument: assumeRolePolicyDocument,
//   	AutoPublishAlias: jsii.String("autoPublishAlias"),
//   	AutoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	CodeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	CodeUri: jsii.String("codeUri"),
//   	DeadLetterQueue: &DeadLetterQueueProperty{
//   		TargetArn: jsii.String("targetArn"),
//   		Type: jsii.String("type"),
//   	},
//   	DeploymentPreference: &DeploymentPreferenceProperty{
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		Hooks: &HooksProperty{
//   			PostTraffic: jsii.String("postTraffic"),
//   			PreTraffic: jsii.String("preTraffic"),
//   		},
//   		Role: jsii.String("role"),
//   		Type: jsii.String("type"),
//   	},
//   	Description: jsii.String("description"),
//   	Environment: &FunctionEnvironmentProperty{
//   		Variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	EventInvokeConfig: &EventInvokeConfigProperty{
//   		DestinationConfig: &EventInvokeDestinationConfigProperty{
//   			OnFailure: &DestinationProperty{
//   				Destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   			OnSuccess: &DestinationProperty{
//   				Destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		MaximumEventAgeInSeconds: jsii.Number(123),
//   		MaximumRetryAttempts: jsii.Number(123),
//   	},
//   	Events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &AlexaSkillEventProperty{
//   				"skillId": jsii.String("skillId"),
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	FileSystemConfigs: []interface{}{
//   		&FileSystemConfigProperty{
//   			Arn: jsii.String("arn"),
//   			LocalMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	FunctionName: jsii.String("functionName"),
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
//   	ImageUri: jsii.String("imageUri"),
//   	InlineCode: jsii.String("inlineCode"),
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	Layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	MemorySize: jsii.Number(123),
//   	PackageType: jsii.String("packageType"),
//   	PermissionsBoundary: jsii.String("permissionsBoundary"),
//   	Policies: jsii.String("policies"),
//   	ProvisionedConcurrencyConfig: &ProvisionedConcurrencyConfigProperty{
//   		ProvisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	ReservedConcurrentExecutions: jsii.Number(123),
//   	Role: jsii.String("role"),
//   	Runtime: jsii.String("runtime"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	Timeout: jsii.Number(123),
//   	Tracing: jsii.String("tracing"),
//   	VersionDescription: jsii.String("versionDescription"),
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html
//
type CfnFunctionProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-architectures
	//
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-assumerolepolicydocument
	//
	AssumeRolePolicyDocument interface{} `field:"optional" json:"assumeRolePolicyDocument" yaml:"assumeRolePolicyDocument"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-autopublishalias
	//
	AutoPublishAlias *string `field:"optional" json:"autoPublishAlias" yaml:"autoPublishAlias"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-autopublishcodesha256
	//
	AutoPublishCodeSha256 *string `field:"optional" json:"autoPublishCodeSha256" yaml:"autoPublishCodeSha256"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-codesigningconfigarn
	//
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-codeuri
	//
	CodeUri interface{} `field:"optional" json:"codeUri" yaml:"codeUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-deadletterqueue
	//
	DeadLetterQueue interface{} `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-deploymentpreference
	//
	DeploymentPreference interface{} `field:"optional" json:"deploymentPreference" yaml:"deploymentPreference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-environment
	//
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-eventinvokeconfig
	//
	EventInvokeConfig interface{} `field:"optional" json:"eventInvokeConfig" yaml:"eventInvokeConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-events
	//
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-filesystemconfigs
	//
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-functionname
	//
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-handler
	//
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-imageconfig
	//
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-imageuri
	//
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-inlinecode
	//
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-layers
	//
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-memorysize
	//
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-packagetype
	//
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-permissionsboundary
	//
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-policies
	//
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-provisionedconcurrencyconfig
	//
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-reservedconcurrentexecutions
	//
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-runtime
	//
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-timeout
	//
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-tracing
	//
	Tracing *string `field:"optional" json:"tracing" yaml:"tracing"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-versiondescription
	//
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-serverless-function.html#cfn-serverless-function-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


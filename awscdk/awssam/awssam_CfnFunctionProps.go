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
//   cfnFunctionProps := &cfnFunctionProps{
//   	architectures: []*string{
//   		jsii.String("architectures"),
//   	},
//   	assumeRolePolicyDocument: assumeRolePolicyDocument,
//   	autoPublishAlias: jsii.String("autoPublishAlias"),
//   	autoPublishCodeSha256: jsii.String("autoPublishCodeSha256"),
//   	codeSigningConfigArn: jsii.String("codeSigningConfigArn"),
//   	codeUri: jsii.String("codeUri"),
//   	deadLetterQueue: &deadLetterQueueProperty{
//   		targetArn: jsii.String("targetArn"),
//   		type: jsii.String("type"),
//   	},
//   	deploymentPreference: &deploymentPreferenceProperty{
//   		enabled: jsii.Boolean(false),
//   		type: jsii.String("type"),
//
//   		// the properties below are optional
//   		alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		hooks: &hooksProperty{
//   			postTraffic: jsii.String("postTraffic"),
//   			preTraffic: jsii.String("preTraffic"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	environment: &functionEnvironmentProperty{
//   		variables: map[string]*string{
//   			"variablesKey": jsii.String("variables"),
//   		},
//   	},
//   	eventInvokeConfig: &eventInvokeConfigProperty{
//   		destinationConfig: &eventInvokeDestinationConfigProperty{
//   			onFailure: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   			onSuccess: &destinationProperty{
//   				destination: jsii.String("destination"),
//
//   				// the properties below are optional
//   				type: jsii.String("type"),
//   			},
//   		},
//   		maximumEventAgeInSeconds: jsii.Number(123),
//   		maximumRetryAttempts: jsii.Number(123),
//   	},
//   	events: map[string]interface{}{
//   		"eventsKey": &EventSourceProperty{
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
//   			},
//   			"type": jsii.String("type"),
//   		},
//   	},
//   	fileSystemConfigs: []interface{}{
//   		&fileSystemConfigProperty{
//   			arn: jsii.String("arn"),
//   			localMountPath: jsii.String("localMountPath"),
//   		},
//   	},
//   	functionName: jsii.String("functionName"),
//   	handler: jsii.String("handler"),
//   	imageConfig: &imageConfigProperty{
//   		command: []*string{
//   			jsii.String("command"),
//   		},
//   		entryPoint: []*string{
//   			jsii.String("entryPoint"),
//   		},
//   		workingDirectory: jsii.String("workingDirectory"),
//   	},
//   	imageUri: jsii.String("imageUri"),
//   	inlineCode: jsii.String("inlineCode"),
//   	kmsKeyArn: jsii.String("kmsKeyArn"),
//   	layers: []*string{
//   		jsii.String("layers"),
//   	},
//   	memorySize: jsii.Number(123),
//   	packageType: jsii.String("packageType"),
//   	permissionsBoundary: jsii.String("permissionsBoundary"),
//   	policies: jsii.String("policies"),
//   	provisionedConcurrencyConfig: &provisionedConcurrencyConfigProperty{
//   		provisionedConcurrentExecutions: jsii.String("provisionedConcurrentExecutions"),
//   	},
//   	reservedConcurrentExecutions: jsii.Number(123),
//   	role: jsii.String("role"),
//   	runtime: jsii.String("runtime"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	timeout: jsii.Number(123),
//   	tracing: jsii.String("tracing"),
//   	versionDescription: jsii.String("versionDescription"),
//   	vpcConfig: &vpcConfigProperty{
//   		securityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		subnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
type CfnFunctionProps struct {
	// `AWS::Serverless::Function.Architectures`.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// `AWS::Serverless::Function.AssumeRolePolicyDocument`.
	AssumeRolePolicyDocument interface{} `field:"optional" json:"assumeRolePolicyDocument" yaml:"assumeRolePolicyDocument"`
	// `AWS::Serverless::Function.AutoPublishAlias`.
	AutoPublishAlias *string `field:"optional" json:"autoPublishAlias" yaml:"autoPublishAlias"`
	// `AWS::Serverless::Function.AutoPublishCodeSha256`.
	AutoPublishCodeSha256 *string `field:"optional" json:"autoPublishCodeSha256" yaml:"autoPublishCodeSha256"`
	// `AWS::Serverless::Function.CodeSigningConfigArn`.
	CodeSigningConfigArn *string `field:"optional" json:"codeSigningConfigArn" yaml:"codeSigningConfigArn"`
	// `AWS::Serverless::Function.CodeUri`.
	CodeUri interface{} `field:"optional" json:"codeUri" yaml:"codeUri"`
	// `AWS::Serverless::Function.DeadLetterQueue`.
	DeadLetterQueue interface{} `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// `AWS::Serverless::Function.DeploymentPreference`.
	DeploymentPreference interface{} `field:"optional" json:"deploymentPreference" yaml:"deploymentPreference"`
	// `AWS::Serverless::Function.Description`.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::Serverless::Function.Environment`.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// `AWS::Serverless::Function.EventInvokeConfig`.
	EventInvokeConfig interface{} `field:"optional" json:"eventInvokeConfig" yaml:"eventInvokeConfig"`
	// `AWS::Serverless::Function.Events`.
	Events interface{} `field:"optional" json:"events" yaml:"events"`
	// `AWS::Serverless::Function.FileSystemConfigs`.
	FileSystemConfigs interface{} `field:"optional" json:"fileSystemConfigs" yaml:"fileSystemConfigs"`
	// `AWS::Serverless::Function.FunctionName`.
	FunctionName *string `field:"optional" json:"functionName" yaml:"functionName"`
	// `AWS::Serverless::Function.Handler`.
	Handler *string `field:"optional" json:"handler" yaml:"handler"`
	// `AWS::Serverless::Function.ImageConfig`.
	ImageConfig interface{} `field:"optional" json:"imageConfig" yaml:"imageConfig"`
	// `AWS::Serverless::Function.ImageUri`.
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
	// `AWS::Serverless::Function.InlineCode`.
	InlineCode *string `field:"optional" json:"inlineCode" yaml:"inlineCode"`
	// `AWS::Serverless::Function.KmsKeyArn`.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// `AWS::Serverless::Function.Layers`.
	Layers *[]*string `field:"optional" json:"layers" yaml:"layers"`
	// `AWS::Serverless::Function.MemorySize`.
	MemorySize *float64 `field:"optional" json:"memorySize" yaml:"memorySize"`
	// `AWS::Serverless::Function.PackageType`.
	PackageType *string `field:"optional" json:"packageType" yaml:"packageType"`
	// `AWS::Serverless::Function.PermissionsBoundary`.
	PermissionsBoundary *string `field:"optional" json:"permissionsBoundary" yaml:"permissionsBoundary"`
	// `AWS::Serverless::Function.Policies`.
	Policies interface{} `field:"optional" json:"policies" yaml:"policies"`
	// `AWS::Serverless::Function.ProvisionedConcurrencyConfig`.
	ProvisionedConcurrencyConfig interface{} `field:"optional" json:"provisionedConcurrencyConfig" yaml:"provisionedConcurrencyConfig"`
	// `AWS::Serverless::Function.ReservedConcurrentExecutions`.
	ReservedConcurrentExecutions *float64 `field:"optional" json:"reservedConcurrentExecutions" yaml:"reservedConcurrentExecutions"`
	// `AWS::Serverless::Function.Role`.
	Role *string `field:"optional" json:"role" yaml:"role"`
	// `AWS::Serverless::Function.Runtime`.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// `AWS::Serverless::Function.Tags`.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// `AWS::Serverless::Function.Timeout`.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
	// `AWS::Serverless::Function.Tracing`.
	Tracing *string `field:"optional" json:"tracing" yaml:"tracing"`
	// `AWS::Serverless::Function.VersionDescription`.
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
	// `AWS::Serverless::Function.VpcConfig`.
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


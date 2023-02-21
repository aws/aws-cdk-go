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
//   		Enabled: jsii.Boolean(false),
//   		Type: jsii.String("type"),
//
//   		// the properties below are optional
//   		Alarms: []*string{
//   			jsii.String("alarms"),
//   		},
//   		Hooks: &HooksProperty{
//   			PostTraffic: jsii.String("postTraffic"),
//   			PreTraffic: jsii.String("preTraffic"),
//   		},
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
//   			"properties": &S3EventProperty{
//   				"variables": map[string]*string{
//   					"variablesKey": jsii.String("variables"),
//   				},
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


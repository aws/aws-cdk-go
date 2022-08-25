package awsgreengrassv2


// Contains information about an AWS Lambda function to import to create a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lambdaFunctionRecipeSourceProperty := &lambdaFunctionRecipeSourceProperty{
//   	componentDependencies: map[string]interface{}{
//   		"componentDependenciesKey": &ComponentDependencyRequirementProperty{
//   			"dependencyType": jsii.String("dependencyType"),
//   			"versionRequirement": jsii.String("versionRequirement"),
//   		},
//   	},
//   	componentLambdaParameters: &lambdaExecutionParametersProperty{
//   		environmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		eventSources: []interface{}{
//   			&lambdaEventSourceProperty{
//   				topic: jsii.String("topic"),
//   				type: jsii.String("type"),
//   			},
//   		},
//   		execArgs: []*string{
//   			jsii.String("execArgs"),
//   		},
//   		inputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   		linuxProcessParams: &lambdaLinuxProcessParamsProperty{
//   			containerParams: &lambdaContainerParamsProperty{
//   				devices: []interface{}{
//   					&lambdaDeviceMountProperty{
//   						addGroupOwner: jsii.Boolean(false),
//   						path: jsii.String("path"),
//   						permission: jsii.String("permission"),
//   					},
//   				},
//   				memorySizeInKb: jsii.Number(123),
//   				mountRoSysfs: jsii.Boolean(false),
//   				volumes: []interface{}{
//   					&lambdaVolumeMountProperty{
//   						addGroupOwner: jsii.Boolean(false),
//   						destinationPath: jsii.String("destinationPath"),
//   						permission: jsii.String("permission"),
//   						sourcePath: jsii.String("sourcePath"),
//   					},
//   				},
//   			},
//   			isolationMode: jsii.String("isolationMode"),
//   		},
//   		maxIdleTimeInSeconds: jsii.Number(123),
//   		maxInstancesCount: jsii.Number(123),
//   		maxQueueSize: jsii.Number(123),
//   		pinned: jsii.Boolean(false),
//   		statusTimeoutInSeconds: jsii.Number(123),
//   		timeoutInSeconds: jsii.Number(123),
//   	},
//   	componentName: jsii.String("componentName"),
//   	componentPlatforms: []interface{}{
//   		&componentPlatformProperty{
//   			attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			name: jsii.String("name"),
//   		},
//   	},
//   	componentVersion: jsii.String("componentVersion"),
//   	lambdaArn: jsii.String("lambdaArn"),
//   }
//
type CfnComponentVersion_LambdaFunctionRecipeSourceProperty struct {
	// The component versions on which this Lambda function component depends.
	ComponentDependencies interface{} `field:"optional" json:"componentDependencies" yaml:"componentDependencies"`
	// The system and runtime parameters for the Lambda function as it runs on the Greengrass core device.
	ComponentLambdaParameters interface{} `field:"optional" json:"componentLambdaParameters" yaml:"componentLambdaParameters"`
	// The name of the component.
	//
	// Defaults to the name of the Lambda function.
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The platforms that the component version supports.
	ComponentPlatforms interface{} `field:"optional" json:"componentPlatforms" yaml:"componentPlatforms"`
	// The version of the component.
	//
	// Defaults to the version of the Lambda function as a semantic version. For example, if your function version is `3` , the component version becomes `3.0.0` .
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// The ARN of the Lambda function.
	//
	// The ARN must include the version of the function to import. You can't use version aliases like `$LATEST` .
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}


package awsgreengrassv2


// Properties for defining a `CfnComponentVersion`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnComponentVersionProps := &cfnComponentVersionProps{
//   	inlineRecipe: jsii.String("inlineRecipe"),
//   	lambdaFunction: &lambdaFunctionRecipeSourceProperty{
//   		componentDependencies: map[string]interface{}{
//   			"componentDependenciesKey": &ComponentDependencyRequirementProperty{
//   				"dependencyType": jsii.String("dependencyType"),
//   				"versionRequirement": jsii.String("versionRequirement"),
//   			},
//   		},
//   		componentLambdaParameters: &lambdaExecutionParametersProperty{
//   			environmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			eventSources: []interface{}{
//   				&lambdaEventSourceProperty{
//   					topic: jsii.String("topic"),
//   					type: jsii.String("type"),
//   				},
//   			},
//   			execArgs: []*string{
//   				jsii.String("execArgs"),
//   			},
//   			inputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   			linuxProcessParams: &lambdaLinuxProcessParamsProperty{
//   				containerParams: &lambdaContainerParamsProperty{
//   					devices: []interface{}{
//   						&lambdaDeviceMountProperty{
//   							addGroupOwner: jsii.Boolean(false),
//   							path: jsii.String("path"),
//   							permission: jsii.String("permission"),
//   						},
//   					},
//   					memorySizeInKb: jsii.Number(123),
//   					mountRoSysfs: jsii.Boolean(false),
//   					volumes: []interface{}{
//   						&lambdaVolumeMountProperty{
//   							addGroupOwner: jsii.Boolean(false),
//   							destinationPath: jsii.String("destinationPath"),
//   							permission: jsii.String("permission"),
//   							sourcePath: jsii.String("sourcePath"),
//   						},
//   					},
//   				},
//   				isolationMode: jsii.String("isolationMode"),
//   			},
//   			maxIdleTimeInSeconds: jsii.Number(123),
//   			maxInstancesCount: jsii.Number(123),
//   			maxQueueSize: jsii.Number(123),
//   			pinned: jsii.Boolean(false),
//   			statusTimeoutInSeconds: jsii.Number(123),
//   			timeoutInSeconds: jsii.Number(123),
//   		},
//   		componentName: jsii.String("componentName"),
//   		componentPlatforms: []interface{}{
//   			&componentPlatformProperty{
//   				attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				name: jsii.String("name"),
//   			},
//   		},
//   		componentVersion: jsii.String("componentVersion"),
//   		lambdaArn: jsii.String("lambdaArn"),
//   	},
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnComponentVersionProps struct {
	// The recipe to use to create the component.
	//
	// The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	InlineRecipe *string `field:"optional" json:"inlineRecipe" yaml:"inlineRecipe"`
	// The parameters to create a component from a Lambda function.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Application-specific metadata to attach to the component version.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in AWS CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


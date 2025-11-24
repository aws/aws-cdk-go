package mixinsawsgreengrassv2


// Properties for CfnComponentVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnComponentVersionMixinProps := &CfnComponentVersionMixinProps{
//   	InlineRecipe: jsii.String("inlineRecipe"),
//   	LambdaFunction: &LambdaFunctionRecipeSourceProperty{
//   		ComponentDependencies: map[string]interface{}{
//   			"componentDependenciesKey": &ComponentDependencyRequirementProperty{
//   				"dependencyType": jsii.String("dependencyType"),
//   				"versionRequirement": jsii.String("versionRequirement"),
//   			},
//   		},
//   		ComponentLambdaParameters: &LambdaExecutionParametersProperty{
//   			EnvironmentVariables: map[string]*string{
//   				"environmentVariablesKey": jsii.String("environmentVariables"),
//   			},
//   			EventSources: []interface{}{
//   				&LambdaEventSourceProperty{
//   					Topic: jsii.String("topic"),
//   					Type: jsii.String("type"),
//   				},
//   			},
//   			ExecArgs: []*string{
//   				jsii.String("execArgs"),
//   			},
//   			InputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   			LinuxProcessParams: &LambdaLinuxProcessParamsProperty{
//   				ContainerParams: &LambdaContainerParamsProperty{
//   					Devices: []interface{}{
//   						&LambdaDeviceMountProperty{
//   							AddGroupOwner: jsii.Boolean(false),
//   							Path: jsii.String("path"),
//   							Permission: jsii.String("permission"),
//   						},
//   					},
//   					MemorySizeInKb: jsii.Number(123),
//   					MountRoSysfs: jsii.Boolean(false),
//   					Volumes: []interface{}{
//   						&LambdaVolumeMountProperty{
//   							AddGroupOwner: jsii.Boolean(false),
//   							DestinationPath: jsii.String("destinationPath"),
//   							Permission: jsii.String("permission"),
//   							SourcePath: jsii.String("sourcePath"),
//   						},
//   					},
//   				},
//   				IsolationMode: jsii.String("isolationMode"),
//   			},
//   			MaxIdleTimeInSeconds: jsii.Number(123),
//   			MaxInstancesCount: jsii.Number(123),
//   			MaxQueueSize: jsii.Number(123),
//   			Pinned: jsii.Boolean(false),
//   			StatusTimeoutInSeconds: jsii.Number(123),
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		ComponentName: jsii.String("componentName"),
//   		ComponentPlatforms: []interface{}{
//   			&ComponentPlatformProperty{
//   				Attributes: map[string]*string{
//   					"attributesKey": jsii.String("attributes"),
//   				},
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		ComponentVersion: jsii.String("componentVersion"),
//   		LambdaArn: jsii.String("lambdaArn"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html
//
type CfnComponentVersionMixinProps struct {
	// The recipe to use to create the component.
	//
	// The recipe defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform compatibility.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-inlinerecipe
	//
	InlineRecipe *string `field:"optional" json:"inlineRecipe" yaml:"inlineRecipe"`
	// The parameters to create a component from a Lambda function.
	//
	// You must specify either `InlineRecipe` or `LambdaFunction` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-lambdafunction
	//
	LambdaFunction interface{} `field:"optional" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Application-specific metadata to attach to the component version.
	//
	// You can use tags in IAM policies to control access to AWS IoT Greengrass resources. You can also use tags to categorize your resources. For more information, see [Tag your AWS IoT Greengrass Version 2 resources](https://docs.aws.amazon.com/greengrass/v2/developerguide/tag-resources.html) in the *AWS IoT Greengrass V2 Developer Guide* .
	//
	// This `Json` property type is processed as a map of key-value pairs. It uses the following format, which is different from most `Tags` implementations in CloudFormation templates.
	//
	// ```json
	// "Tags": { "KeyName0": "value", "KeyName1": "value", "KeyName2": "value"
	// }
	// ```.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html#cfn-greengrassv2-componentversion-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


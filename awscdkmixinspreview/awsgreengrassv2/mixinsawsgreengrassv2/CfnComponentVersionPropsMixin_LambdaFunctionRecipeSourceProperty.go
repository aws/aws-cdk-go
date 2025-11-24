package mixinsawsgreengrassv2


// Contains information about an AWS Lambda function to import to create a component.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lambdaFunctionRecipeSourceProperty := &LambdaFunctionRecipeSourceProperty{
//   	ComponentDependencies: map[string]interface{}{
//   		"componentDependenciesKey": &ComponentDependencyRequirementProperty{
//   			"dependencyType": jsii.String("dependencyType"),
//   			"versionRequirement": jsii.String("versionRequirement"),
//   		},
//   	},
//   	ComponentLambdaParameters: &LambdaExecutionParametersProperty{
//   		EnvironmentVariables: map[string]*string{
//   			"environmentVariablesKey": jsii.String("environmentVariables"),
//   		},
//   		EventSources: []interface{}{
//   			&LambdaEventSourceProperty{
//   				Topic: jsii.String("topic"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   		ExecArgs: []*string{
//   			jsii.String("execArgs"),
//   		},
//   		InputPayloadEncodingType: jsii.String("inputPayloadEncodingType"),
//   		LinuxProcessParams: &LambdaLinuxProcessParamsProperty{
//   			ContainerParams: &LambdaContainerParamsProperty{
//   				Devices: []interface{}{
//   					&LambdaDeviceMountProperty{
//   						AddGroupOwner: jsii.Boolean(false),
//   						Path: jsii.String("path"),
//   						Permission: jsii.String("permission"),
//   					},
//   				},
//   				MemorySizeInKb: jsii.Number(123),
//   				MountRoSysfs: jsii.Boolean(false),
//   				Volumes: []interface{}{
//   					&LambdaVolumeMountProperty{
//   						AddGroupOwner: jsii.Boolean(false),
//   						DestinationPath: jsii.String("destinationPath"),
//   						Permission: jsii.String("permission"),
//   						SourcePath: jsii.String("sourcePath"),
//   					},
//   				},
//   			},
//   			IsolationMode: jsii.String("isolationMode"),
//   		},
//   		MaxIdleTimeInSeconds: jsii.Number(123),
//   		MaxInstancesCount: jsii.Number(123),
//   		MaxQueueSize: jsii.Number(123),
//   		Pinned: jsii.Boolean(false),
//   		StatusTimeoutInSeconds: jsii.Number(123),
//   		TimeoutInSeconds: jsii.Number(123),
//   	},
//   	ComponentName: jsii.String("componentName"),
//   	ComponentPlatforms: []interface{}{
//   		&ComponentPlatformProperty{
//   			Attributes: map[string]*string{
//   				"attributesKey": jsii.String("attributes"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	ComponentVersion: jsii.String("componentVersion"),
//   	LambdaArn: jsii.String("lambdaArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html
//
type CfnComponentVersionPropsMixin_LambdaFunctionRecipeSourceProperty struct {
	// The component versions on which this Lambda function component depends.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentdependencies
	//
	ComponentDependencies interface{} `field:"optional" json:"componentDependencies" yaml:"componentDependencies"`
	// The system and runtime parameters for the Lambda function as it runs on the AWS IoT Greengrass core device.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentlambdaparameters
	//
	ComponentLambdaParameters interface{} `field:"optional" json:"componentLambdaParameters" yaml:"componentLambdaParameters"`
	// The name of the component.
	//
	// Defaults to the name of the Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentname
	//
	ComponentName *string `field:"optional" json:"componentName" yaml:"componentName"`
	// The platforms that the component version supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentplatforms
	//
	ComponentPlatforms interface{} `field:"optional" json:"componentPlatforms" yaml:"componentPlatforms"`
	// The version of the component.
	//
	// Defaults to the version of the Lambda function as a semantic version. For example, if your function version is `3` , the component version becomes `3.0.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-componentversion
	//
	ComponentVersion *string `field:"optional" json:"componentVersion" yaml:"componentVersion"`
	// The ARN of the Lambda function.
	//
	// The ARN must include the version of the function to import. You can't use version aliases like `$LATEST` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrassv2-componentversion-lambdafunctionrecipesource.html#cfn-greengrassv2-componentversion-lambdafunctionrecipesource-lambdaarn
	//
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}


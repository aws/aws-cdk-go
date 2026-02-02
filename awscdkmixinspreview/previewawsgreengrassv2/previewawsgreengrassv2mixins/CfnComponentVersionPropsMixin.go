package previewawsgreengrassv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrassv2/previewawsgreengrassv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a component.
//
// Components are software that run on AWS IoT Greengrass core devices. After you develop and test a component on your core device, you can use this operation to upload your component to AWS IoT Greengrass . Then, you can deploy the component to other core devices.
//
// You can use this operation to do the following:
//
// - *Create components from recipes*
//
// Create a component from a recipe, which is a file that defines the component's metadata, parameters, dependencies, lifecycle, artifacts, and platform capability. For more information, see [AWS IoT Greengrass component recipe reference](https://docs.aws.amazon.com/greengrass/v2/developerguide/component-recipe-reference.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// To create a component from a recipe, specify `inlineRecipe` when you call this operation.
// - *Create components from Lambda functions*
//
// Create a component from an AWS Lambda function that runs on AWS IoT Greengrass . This creates a recipe and artifacts from the Lambda function's deployment package. You can use this operation to migrate Lambda functions from AWS IoT Greengrass V1 to AWS IoT Greengrass V2 .
//
// This function accepts Lambda functions in all supported versions of Python, Node.js, and Java runtimes. AWS IoT Greengrass doesn't apply any additional restrictions on deprecated Lambda runtime versions.
//
// To create a component from a Lambda function, specify `lambdaFunction` when you call this operation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnComponentVersionPropsMixin := awscdkmixinspreview.Mixins.NewCfnComponentVersionPropsMixin(&CfnComponentVersionMixinProps{
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
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-componentversion.html
//
type CfnComponentVersionPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnComponentVersionMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct)
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnComponentVersionPropsMixin
type jsiiProxy_CfnComponentVersionPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnComponentVersionPropsMixin) Props() *CfnComponentVersionMixinProps {
	var returns *CfnComponentVersionMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnComponentVersionPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GreengrassV2::ComponentVersion`.
func NewCfnComponentVersionPropsMixin(props *CfnComponentVersionMixinProps, options *mixins.CfnPropertyMixinOptions) CfnComponentVersionPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnComponentVersionPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnComponentVersionPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GreengrassV2::ComponentVersion`.
func NewCfnComponentVersionPropsMixin_Override(c CfnComponentVersionPropsMixin, props *CfnComponentVersionMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnComponentVersionPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnComponentVersionPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnComponentVersionPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnComponentVersionPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnComponentVersionPropsMixin) ApplyTo(construct constructs.IConstruct) {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"applyTo",
		[]interface{}{construct},
	)
}

func (c *jsiiProxy_CfnComponentVersionPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


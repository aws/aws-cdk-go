package mixinsawssagemaker

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/awssagemaker/mixinsawssagemaker/internal"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an inference component, which is a SageMaker AI hosting object that you can use to deploy a model to an endpoint.
//
// In the inference component settings, you specify the model, the endpoint, and how the model utilizes the resources that the endpoint hosts. You can optimize resource utilization by tailoring how the required CPU cores, accelerators, and memory are allocated. You can deploy multiple inference components to an endpoint, where each inference component contains one model and the resource utilization needs for that individual model. After you deploy an inference component, you can directly invoke the associated model when you use the InvokeEndpoint API action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnInferenceComponentPropsMixin := awscdkmixinspreview.Mixins.NewCfnInferenceComponentPropsMixin(&CfnInferenceComponentMixinProps{
//   	DeploymentConfig: &InferenceComponentDeploymentConfigProperty{
//   		AutoRollbackConfiguration: &AutoRollbackConfigurationProperty{
//   			Alarms: []interface{}{
//   				&AlarmProperty{
//   					AlarmName: jsii.String("alarmName"),
//   				},
//   			},
//   		},
//   		RollingUpdatePolicy: &InferenceComponentRollingUpdatePolicyProperty{
//   			MaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			RollbackMaximumBatchSize: &InferenceComponentCapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			WaitIntervalInSeconds: jsii.Number(123),
//   		},
//   	},
//   	EndpointArn: jsii.String("endpointArn"),
//   	EndpointName: jsii.String("endpointName"),
//   	InferenceComponentName: jsii.String("inferenceComponentName"),
//   	RuntimeConfig: &InferenceComponentRuntimeConfigProperty{
//   		CopyCount: jsii.Number(123),
//   		CurrentCopyCount: jsii.Number(123),
//   		DesiredCopyCount: jsii.Number(123),
//   	},
//   	Specification: &InferenceComponentSpecificationProperty{
//   		BaseInferenceComponentName: jsii.String("baseInferenceComponentName"),
//   		ComputeResourceRequirements: &InferenceComponentComputeResourceRequirementsProperty{
//   			MaxMemoryRequiredInMb: jsii.Number(123),
//   			MinMemoryRequiredInMb: jsii.Number(123),
//   			NumberOfAcceleratorDevicesRequired: jsii.Number(123),
//   			NumberOfCpuCoresRequired: jsii.Number(123),
//   		},
//   		Container: &InferenceComponentContainerSpecificationProperty{
//   			ArtifactUrl: jsii.String("artifactUrl"),
//   			DeployedImage: &DeployedImageProperty{
//   				ResolutionTime: jsii.String("resolutionTime"),
//   				ResolvedImage: jsii.String("resolvedImage"),
//   				SpecifiedImage: jsii.String("specifiedImage"),
//   			},
//   			Environment: map[string]*string{
//   				"environmentKey": jsii.String("environment"),
//   			},
//   			Image: jsii.String("image"),
//   		},
//   		ModelName: jsii.String("modelName"),
//   		StartupParameters: &InferenceComponentStartupParametersProperty{
//   			ContainerStartupHealthCheckTimeoutInSeconds: jsii.Number(123),
//   			ModelDataDownloadTimeoutInSeconds: jsii.Number(123),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	VariantName: jsii.String("variantName"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-inferencecomponent.html
//
type CfnInferenceComponentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnInferenceComponentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnInferenceComponentPropsMixin
type jsiiProxy_CfnInferenceComponentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnInferenceComponentPropsMixin) Props() *CfnInferenceComponentMixinProps {
	var returns *CfnInferenceComponentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnInferenceComponentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::InferenceComponent`.
func NewCfnInferenceComponentPropsMixin(props *CfnInferenceComponentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnInferenceComponentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnInferenceComponentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnInferenceComponentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::InferenceComponent`.
func NewCfnInferenceComponentPropsMixin_Override(c CfnInferenceComponentPropsMixin, props *CfnInferenceComponentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnInferenceComponentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnInferenceComponentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnInferenceComponentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnInferenceComponentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnInferenceComponentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnInferenceComponentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


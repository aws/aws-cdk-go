package previewawsgreengrassv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsgreengrassv2/previewawsgreengrassv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a continuous deployment for a target, which is a AWS IoT Greengrass core device or group of core devices.
//
// When you add a new core device to a group of core devices that has a deployment, AWS IoT Greengrass deploys that group's deployment to the new device.
//
// You can define one deployment for each target. When you create a new deployment for a target that has an existing deployment, you replace the previous deployment. AWS IoT Greengrass applies the new deployment to the target devices.
//
// You can only add, update, or delete up to 10 deployments at a time to a single target.
//
// Every deployment has a revision number that indicates how many deployment revisions you define for a target. Use this operation to create a new revision of an existing deployment. This operation returns the revision number of the new deployment when you create it.
//
// For more information, see the [Create deployments](https://docs.aws.amazon.com/greengrass/v2/developerguide/create-deployments.html) in the *AWS IoT Greengrass V2 Developer Guide* .
//
// > Deployment resources are deleted when you delete stacks. To keep the deployments in a stack, you must specify `"DeletionPolicy": "Retain"` on each deployment resource in the stack template that you want to keep. For more information, see [DeletionPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html) .
// >
// > You can only delete up to 10 deployment resources at a time. If you delete more than 10 resources, you receive an error.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var rateIncreaseCriteria interface{}
//
//   cfnDeploymentPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeploymentPropsMixin(&CfnDeploymentMixinProps{
//   	Components: map[string]interface{}{
//   		"componentsKey": &ComponentDeploymentSpecificationProperty{
//   			"componentVersion": jsii.String("componentVersion"),
//   			"configurationUpdate": &ComponentConfigurationUpdateProperty{
//   				"merge": jsii.String("merge"),
//   				"reset": []*string{
//   					jsii.String("reset"),
//   				},
//   			},
//   			"runWith": &ComponentRunWithProperty{
//   				"posixUser": jsii.String("posixUser"),
//   				"systemResourceLimits": &SystemResourceLimitsProperty{
//   					"cpus": jsii.Number(123),
//   					"memory": jsii.Number(123),
//   				},
//   				"windowsUser": jsii.String("windowsUser"),
//   			},
//   		},
//   	},
//   	DeploymentName: jsii.String("deploymentName"),
//   	DeploymentPolicies: &DeploymentPoliciesProperty{
//   		ComponentUpdatePolicy: &DeploymentComponentUpdatePolicyProperty{
//   			Action: jsii.String("action"),
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		ConfigurationValidationPolicy: &DeploymentConfigurationValidationPolicyProperty{
//   			TimeoutInSeconds: jsii.Number(123),
//   		},
//   		FailureHandlingPolicy: jsii.String("failureHandlingPolicy"),
//   	},
//   	IotJobConfiguration: &DeploymentIoTJobConfigurationProperty{
//   		AbortConfig: &IoTJobAbortConfigProperty{
//   			CriteriaList: []interface{}{
//   				&IoTJobAbortCriteriaProperty{
//   					Action: jsii.String("action"),
//   					FailureType: jsii.String("failureType"),
//   					MinNumberOfExecutedThings: jsii.Number(123),
//   					ThresholdPercentage: jsii.Number(123),
//   				},
//   			},
//   		},
//   		JobExecutionsRolloutConfig: &IoTJobExecutionsRolloutConfigProperty{
//   			ExponentialRate: &IoTJobExponentialRolloutRateProperty{
//   				BaseRatePerMinute: jsii.Number(123),
//   				IncrementFactor: jsii.Number(123),
//   				RateIncreaseCriteria: rateIncreaseCriteria,
//   			},
//   			MaximumPerMinute: jsii.Number(123),
//   		},
//   		TimeoutConfig: &IoTJobTimeoutConfigProperty{
//   			InProgressTimeoutInMinutes: jsii.Number(123),
//   		},
//   	},
//   	ParentTargetArn: jsii.String("parentTargetArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	TargetArn: jsii.String("targetArn"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-greengrassv2-deployment.html
//
type CfnDeploymentPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeploymentMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeploymentPropsMixin
type jsiiProxy_CfnDeploymentPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Props() *CfnDeploymentMixinProps {
	var returns *CfnDeploymentMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::GreengrassV2::Deployment`.
func NewCfnDeploymentPropsMixin(props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeploymentPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeploymentPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::GreengrassV2::Deployment`.
func NewCfnDeploymentPropsMixin_Override(c CfnDeploymentPropsMixin, props *CfnDeploymentMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeploymentPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_greengrassv2.mixins.CfnDeploymentPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDeploymentPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawscodedeploymixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawscodedeploy/previewawscodedeploymixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::CodeDeploy::DeploymentGroup` resource creates an AWS CodeDeploy deployment group that specifies which instances your application revisions are deployed to, along with other deployment options.
//
// For more information, see [CreateDeploymentGroup](https://docs.aws.amazon.com/codedeploy/latest/APIReference/API_CreateDeploymentGroup.html) in the *CodeDeploy API Reference* .
//
// > Amazon ECS blue/green deployments through CodeDeploy do not use the `AWS::CodeDeploy::DeploymentGroup` resource. To perform Amazon ECS blue/green deployments, use the `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentGroupPropsMixin := awscdkmixinspreview.Mixins.NewCfnDeploymentGroupPropsMixin(&CfnDeploymentGroupMixinProps{
//   	AlarmConfiguration: &AlarmConfigurationProperty{
//   		Alarms: []interface{}{
//   			&AlarmProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		Enabled: jsii.Boolean(false),
//   		IgnorePollAlarmFailure: jsii.Boolean(false),
//   	},
//   	ApplicationName: jsii.String("applicationName"),
//   	AutoRollbackConfiguration: &AutoRollbackConfigurationProperty{
//   		Enabled: jsii.Boolean(false),
//   		Events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	AutoScalingGroups: []*string{
//   		jsii.String("autoScalingGroups"),
//   	},
//   	BlueGreenDeploymentConfiguration: &BlueGreenDeploymentConfigurationProperty{
//   		DeploymentReadyOption: &DeploymentReadyOptionProperty{
//   			ActionOnTimeout: jsii.String("actionOnTimeout"),
//   			WaitTimeInMinutes: jsii.Number(123),
//   		},
//   		GreenFleetProvisioningOption: &GreenFleetProvisioningOptionProperty{
//   			Action: jsii.String("action"),
//   		},
//   		TerminateBlueInstancesOnDeploymentSuccess: &BlueInstanceTerminationOptionProperty{
//   			Action: jsii.String("action"),
//   			TerminationWaitTimeInMinutes: jsii.Number(123),
//   		},
//   	},
//   	Deployment: &DeploymentProperty{
//   		Description: jsii.String("description"),
//   		IgnoreApplicationStopFailures: jsii.Boolean(false),
//   		Revision: &RevisionLocationProperty{
//   			GitHubLocation: &GitHubLocationProperty{
//   				CommitId: jsii.String("commitId"),
//   				Repository: jsii.String("repository"),
//   			},
//   			RevisionType: jsii.String("revisionType"),
//   			S3Location: &S3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				BundleType: jsii.String("bundleType"),
//   				ETag: jsii.String("eTag"),
//   				Key: jsii.String("key"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	DeploymentConfigName: jsii.String("deploymentConfigName"),
//   	DeploymentGroupName: jsii.String("deploymentGroupName"),
//   	DeploymentStyle: &DeploymentStyleProperty{
//   		DeploymentOption: jsii.String("deploymentOption"),
//   		DeploymentType: jsii.String("deploymentType"),
//   	},
//   	Ec2TagFilters: []interface{}{
//   		&EC2TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Ec2TagSet: &EC2TagSetProperty{
//   		Ec2TagSetList: []interface{}{
//   			&EC2TagSetListObjectProperty{
//   				Ec2TagGroup: []interface{}{
//   					&EC2TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	EcsServices: []interface{}{
//   		&ECSServiceProperty{
//   			ClusterName: jsii.String("clusterName"),
//   			ServiceName: jsii.String("serviceName"),
//   		},
//   	},
//   	LoadBalancerInfo: &LoadBalancerInfoProperty{
//   		ElbInfoList: []interface{}{
//   			&ELBInfoProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TargetGroupInfoList: []interface{}{
//   			&TargetGroupInfoProperty{
//   				Name: jsii.String("name"),
//   			},
//   		},
//   		TargetGroupPairInfoList: []interface{}{
//   			&TargetGroupPairInfoProperty{
//   				ProdTrafficRoute: &TrafficRouteProperty{
//   					ListenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   				TargetGroups: []interface{}{
//   					&TargetGroupInfoProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				TestTrafficRoute: &TrafficRouteProperty{
//   					ListenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	OnPremisesInstanceTagFilters: []interface{}{
//   		&TagFilterProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	OnPremisesTagSet: &OnPremisesTagSetProperty{
//   		OnPremisesTagSetList: []interface{}{
//   			&OnPremisesTagSetListObjectProperty{
//   				OnPremisesTagGroup: []interface{}{
//   					&TagFilterProperty{
//   						Key: jsii.String("key"),
//   						Type: jsii.String("type"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	OutdatedInstancesStrategy: jsii.String("outdatedInstancesStrategy"),
//   	ServiceRoleArn: jsii.String("serviceRoleArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TerminationHookEnabled: jsii.Boolean(false),
//   	TriggerConfigurations: []interface{}{
//   		&TriggerConfigProperty{
//   			TriggerEvents: []*string{
//   				jsii.String("triggerEvents"),
//   			},
//   			TriggerName: jsii.String("triggerName"),
//   			TriggerTargetArn: jsii.String("triggerTargetArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
//
type CfnDeploymentGroupPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnDeploymentGroupMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnDeploymentGroupPropsMixin
type jsiiProxy_CfnDeploymentGroupPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnDeploymentGroupPropsMixin) Props() *CfnDeploymentGroupMixinProps {
	var returns *CfnDeploymentGroupMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnDeploymentGroupPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::CodeDeploy::DeploymentGroup`.
func NewCfnDeploymentGroupPropsMixin(props *CfnDeploymentGroupMixinProps, options *mixins.CfnPropertyMixinOptions) CfnDeploymentGroupPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnDeploymentGroupPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnDeploymentGroupPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::CodeDeploy::DeploymentGroup`.
func NewCfnDeploymentGroupPropsMixin_Override(c CfnDeploymentGroupPropsMixin, props *CfnDeploymentGroupMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnDeploymentGroupPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnDeploymentGroupPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnDeploymentGroupPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_codedeploy.mixins.CfnDeploymentGroupPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnDeploymentGroupPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnDeploymentGroupPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawssagemakermixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawssagemaker/previewawssagemakermixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates a SageMaker HyperPod cluster.
//
// SageMaker HyperPod is a capability of SageMaker for creating and managing persistent clusters for developing large machine learning models, such as large language models (LLMs) and diffusion models. To learn more, see [Amazon SageMaker HyperPod](https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-hyperpod.html) in the *Amazon SageMaker Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var onDemand interface{}
//   var spot interface{}
//
//   cfnClusterPropsMixin := awscdkmixinspreview.Mixins.NewCfnClusterPropsMixin(&CfnClusterMixinProps{
//   	AutoScaling: &ClusterAutoScalingConfigProperty{
//   		AutoScalerType: jsii.String("autoScalerType"),
//   		Mode: jsii.String("mode"),
//   	},
//   	ClusterName: jsii.String("clusterName"),
//   	ClusterRole: jsii.String("clusterRole"),
//   	InstanceGroups: []interface{}{
//   		&ClusterInstanceGroupProperty{
//   			CapacityRequirements: &ClusterCapacityRequirementsProperty{
//   				OnDemand: onDemand,
//   				Spot: spot,
//   			},
//   			CurrentCount: jsii.Number(123),
//   			ExecutionRole: jsii.String("executionRole"),
//   			ImageId: jsii.String("imageId"),
//   			InstanceCount: jsii.Number(123),
//   			InstanceGroupName: jsii.String("instanceGroupName"),
//   			InstanceStorageConfigs: []interface{}{
//   				&ClusterInstanceStorageConfigProperty{
//   					EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   						RootVolume: jsii.Boolean(false),
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   						VolumeSizeInGb: jsii.Number(123),
//   					},
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			KubernetesConfig: &ClusterKubernetesConfigProperty{
//   				Labels: map[string]*string{
//   					"labelsKey": jsii.String("labels"),
//   				},
//   				Taints: []interface{}{
//   					&ClusterKubernetesTaintProperty{
//   						Effect: jsii.String("effect"),
//   						Key: jsii.String("key"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   			},
//   			LifeCycleConfig: &ClusterLifeCycleConfigProperty{
//   				OnCreate: jsii.String("onCreate"),
//   				SourceS3Uri: jsii.String("sourceS3Uri"),
//   			},
//   			OnStartDeepHealthChecks: []*string{
//   				jsii.String("onStartDeepHealthChecks"),
//   			},
//   			OverrideVpcConfig: &VpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			ScheduledUpdateConfig: &ScheduledUpdateConfigProperty{
//   				DeploymentConfig: &DeploymentConfigProperty{
//   					AutoRollbackConfiguration: []interface{}{
//   						&AlarmDetailsProperty{
//   							AlarmName: jsii.String("alarmName"),
//   						},
//   					},
//   					RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   						MaximumBatchSize: &CapacitySizeConfigProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.Number(123),
//   						},
//   						RollbackMaximumBatchSize: &CapacitySizeConfigProperty{
//   							Type: jsii.String("type"),
//   							Value: jsii.Number(123),
//   						},
//   					},
//   					WaitIntervalInSeconds: jsii.Number(123),
//   				},
//   				ScheduleExpression: jsii.String("scheduleExpression"),
//   			},
//   			ThreadsPerCore: jsii.Number(123),
//   			TrainingPlanArn: jsii.String("trainingPlanArn"),
//   		},
//   	},
//   	NodeProvisioningMode: jsii.String("nodeProvisioningMode"),
//   	NodeRecovery: jsii.String("nodeRecovery"),
//   	Orchestrator: &OrchestratorProperty{
//   		Eks: &ClusterOrchestratorEksConfigProperty{
//   			ClusterArn: jsii.String("clusterArn"),
//   		},
//   	},
//   	RestrictedInstanceGroups: []interface{}{
//   		&ClusterRestrictedInstanceGroupProperty{
//   			CurrentCount: jsii.Number(123),
//   			EnvironmentConfig: &EnvironmentConfigProperty{
//   				FSxLustreConfig: &FSxLustreConfigProperty{
//   					PerUnitStorageThroughput: jsii.Number(123),
//   					SizeInGiB: jsii.Number(123),
//   				},
//   			},
//   			ExecutionRole: jsii.String("executionRole"),
//   			InstanceCount: jsii.Number(123),
//   			InstanceGroupName: jsii.String("instanceGroupName"),
//   			InstanceStorageConfigs: []interface{}{
//   				&ClusterInstanceStorageConfigProperty{
//   					EbsVolumeConfig: &ClusterEbsVolumeConfigProperty{
//   						RootVolume: jsii.Boolean(false),
//   						VolumeKmsKeyId: jsii.String("volumeKmsKeyId"),
//   						VolumeSizeInGb: jsii.Number(123),
//   					},
//   				},
//   			},
//   			InstanceType: jsii.String("instanceType"),
//   			OnStartDeepHealthChecks: []*string{
//   				jsii.String("onStartDeepHealthChecks"),
//   			},
//   			OverrideVpcConfig: &VpcConfigProperty{
//   				SecurityGroupIds: []*string{
//   					jsii.String("securityGroupIds"),
//   				},
//   				Subnets: []*string{
//   					jsii.String("subnets"),
//   				},
//   			},
//   			ThreadsPerCore: jsii.Number(123),
//   			TrainingPlanArn: jsii.String("trainingPlanArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TieredStorageConfig: &TieredStorageConfigProperty{
//   		InstanceMemoryAllocationPercentage: jsii.Number(123),
//   		Mode: jsii.String("mode"),
//   	},
//   	VpcConfig: &VpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		Subnets: []*string{
//   			jsii.String("subnets"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-cluster.html
//
type CfnClusterPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnClusterMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnClusterPropsMixin
type jsiiProxy_CfnClusterPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnClusterPropsMixin) Props() *CfnClusterMixinProps {
	var returns *CfnClusterMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnClusterPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::SageMaker::Cluster`.
func NewCfnClusterPropsMixin(props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) CfnClusterPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnClusterPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnClusterPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::SageMaker::Cluster`.
func NewCfnClusterPropsMixin_Override(c CfnClusterPropsMixin, props *CfnClusterMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnClusterPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnClusterPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnClusterPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_sagemaker.mixins.CfnClusterPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnClusterPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnClusterPropsMixin) Supports(construct constructs.IConstruct) *bool {
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


package previewawsecsmixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsecs/previewawsecsmixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// The `AWS::ECS::Service` resource creates an Amazon Elastic Container Service (Amazon ECS) service that runs and maintains the requested number of tasks and associated load balancers.
//
// > The stack update fails if you change any properties that require replacement and at least one Amazon ECS Service Connect `ServiceConnectConfiguration` property is configured. This is because AWS CloudFormation creates the replacement service first, but each `ServiceConnectService` must have a name that is unique in the namespace. > Starting April 15, 2023, AWS ; will not onboard new customers to Amazon Elastic Inference (EI), and will help current customers migrate their workloads to options that offer better price and performance. After April 15, 2023, new customers will not be able to launch instances with Amazon EI accelerators in Amazon SageMaker, Amazon ECS , or Amazon EC2 . However, customers who have used Amazon EI at least once during the past 30-day period are considered current customers and will be able to continue using the service. > On June 12, 2025, Amazon ECS launched support for updating capacity provider configuration for Amazon ECS services. With this launch, Amazon ECS also aligned the CloudFormation update behavior for `CapacityProviderStrategy` parameter with the standard practice. For more information, see [Amazon ECS adds support for updating capacity provider configuration for ECS services](https://docs.aws.amazon.com/about-aws/whats-new/2025/05/amazon-ecs-capacity-provider-configuration-ecs/) . Previously Amazon ECS ignored the `CapacityProviderStrategy` property if it was set to an empty list for example, `[]` in CloudFormation , because updating capacity provider configuration was not supported. Now, with support for capacity provider updates, customers can remove capacity providers from a service by passing an empty list. When you specify an empty list ( `[]` ) for the `CapacityProviderStrategy` property in your CloudFormation template, Amazon ECS will remove any capacity providers associated with the service, as follows:
// >
// > - For services created with a capacity provider strategy after the launch:
// >
// > - If there's a cluster default strategy set, the service will revert to using that default strategy.
// > - If no cluster default strategy exists, you will receive the following error:
// >
// > No launch type to fall back to for empty capacity provider strategy. Your service was not created with a launch type.
// > - For services created with a capacity provider strategy prior to the launch:
// >
// > - If `CapacityProviderStrategy` had `FARGATE_SPOT` or `FARGATE` capacity providers, the launch type will be updated to `FARGATE` and the capacity provider will be removed.
// > - If the strategy included Auto Scaling group capacity providers, the service will revert to EC2 launch type, and the Auto Scaling group capacity providers will not be used.
// >
// > Recommended Actions
// >
// > If you are currently using `CapacityProviderStrategy: []` in your CloudFormation templates, you should take one of the following actions:
// >
// > - If you do not intend to update the Capacity Provider Strategy:
// >
// > - Remove the `CapacityProviderStrategy` property entirely from your CloudFormation template
// > - Alternatively, use `!Ref AWS ::NoValue` for the `CapacityProviderStrategy` property in your template
// > - If you intend to maintain or update the Capacity Provider Strategy, specify the actual Capacity Provider Strategy for the service in your CloudFormation template.
// >
// > If your CloudFormation template had an empty list ([]) for `CapacityProviderStrategy` prior to the aforementioned launch on June 12, and you are using the same template with `CapacityProviderStrategy: []` , you might encounter the following error:
// >
// > Invalid request provided: When switching from launch type to capacity provider strategy on an existing service, or making a change to a capacity provider strategy on a service that is already using one, you must force a new deployment. (Service: Ecs, Status Code: 400, Request ID: xxx) (SDK Attempt Count: 1)" (RequestToken: xxx HandlerErrorCode: InvalidRequest)
// >
// > Note that CloudFormation automatically initiates a new deployment when it detects a parameter change, but customers cannot choose to force a deployment through CloudFormation . This is an invalid input scenario that requires one of the remediation actions listed above.
// >
// > If you are experiencing active production issues related to this change, contact AWS Support or your Technical Account Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var hookDetails interface{}
//
//   cfnServicePropsMixin := awscdkmixinspreview.Mixins.NewCfnServicePropsMixin(&CfnServiceMixinProps{
//   	AvailabilityZoneRebalancing: jsii.String("availabilityZoneRebalancing"),
//   	CapacityProviderStrategy: []interface{}{
//   		&CapacityProviderStrategyItemProperty{
//   			Base: jsii.Number(123),
//   			CapacityProvider: jsii.String("capacityProvider"),
//   			Weight: jsii.Number(123),
//   		},
//   	},
//   	Cluster: jsii.String("cluster"),
//   	DeploymentConfiguration: &DeploymentConfigurationProperty{
//   		Alarms: &DeploymentAlarmsProperty{
//   			AlarmNames: []*string{
//   				jsii.String("alarmNames"),
//   			},
//   			Enable: jsii.Boolean(false),
//   			Rollback: jsii.Boolean(false),
//   		},
//   		BakeTimeInMinutes: jsii.Number(123),
//   		CanaryConfiguration: &CanaryConfigurationProperty{
//   			CanaryBakeTimeInMinutes: jsii.Number(123),
//   			CanaryPercent: jsii.Number(123),
//   		},
//   		DeploymentCircuitBreaker: &DeploymentCircuitBreakerProperty{
//   			Enable: jsii.Boolean(false),
//   			Rollback: jsii.Boolean(false),
//   		},
//   		LifecycleHooks: []interface{}{
//   			&DeploymentLifecycleHookProperty{
//   				HookDetails: hookDetails,
//   				HookTargetArn: jsii.String("hookTargetArn"),
//   				LifecycleStages: []*string{
//   					jsii.String("lifecycleStages"),
//   				},
//   				RoleArn: jsii.String("roleArn"),
//   			},
//   		},
//   		LinearConfiguration: &LinearConfigurationProperty{
//   			StepBakeTimeInMinutes: jsii.Number(123),
//   			StepPercent: jsii.Number(123),
//   		},
//   		MaximumPercent: jsii.Number(123),
//   		MinimumHealthyPercent: jsii.Number(123),
//   		Strategy: jsii.String("strategy"),
//   	},
//   	DeploymentController: &DeploymentControllerProperty{
//   		Type: jsii.String("type"),
//   	},
//   	DesiredCount: jsii.Number(123),
//   	EnableEcsManagedTags: jsii.Boolean(false),
//   	EnableExecuteCommand: jsii.Boolean(false),
//   	ForceNewDeployment: &ForceNewDeploymentProperty{
//   		EnableForceNewDeployment: jsii.Boolean(false),
//   		ForceNewDeploymentNonce: jsii.String("forceNewDeploymentNonce"),
//   	},
//   	HealthCheckGracePeriodSeconds: jsii.Number(123),
//   	LaunchType: jsii.String("launchType"),
//   	LoadBalancers: []interface{}{
//   		&LoadBalancerProperty{
//   			AdvancedConfiguration: &AdvancedConfigurationProperty{
//   				AlternateTargetGroupArn: jsii.String("alternateTargetGroupArn"),
//   				ProductionListenerRule: jsii.String("productionListenerRule"),
//   				RoleArn: jsii.String("roleArn"),
//   				TestListenerRule: jsii.String("testListenerRule"),
//   			},
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			LoadBalancerName: jsii.String("loadBalancerName"),
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   	NetworkConfiguration: &NetworkConfigurationProperty{
//   		AwsvpcConfiguration: &AwsVpcConfigurationProperty{
//   			AssignPublicIp: jsii.String("assignPublicIp"),
//   			SecurityGroups: []*string{
//   				jsii.String("securityGroups"),
//   			},
//   			Subnets: []*string{
//   				jsii.String("subnets"),
//   			},
//   		},
//   	},
//   	PlacementConstraints: []interface{}{
//   		&PlacementConstraintProperty{
//   			Expression: jsii.String("expression"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	PlacementStrategies: []interface{}{
//   		&PlacementStrategyProperty{
//   			Field: jsii.String("field"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	PlatformVersion: jsii.String("platformVersion"),
//   	PropagateTags: jsii.String("propagateTags"),
//   	Role: jsii.String("role"),
//   	SchedulingStrategy: jsii.String("schedulingStrategy"),
//   	ServiceConnectConfiguration: &ServiceConnectConfigurationProperty{
//   		AccessLogConfiguration: &ServiceConnectAccessLogConfigurationProperty{
//   			Format: jsii.String("format"),
//   			IncludeQueryParameters: jsii.String("includeQueryParameters"),
//   		},
//   		Enabled: jsii.Boolean(false),
//   		LogConfiguration: &LogConfigurationProperty{
//   			LogDriver: jsii.String("logDriver"),
//   			Options: map[string]*string{
//   				"optionsKey": jsii.String("options"),
//   			},
//   			SecretOptions: []interface{}{
//   				&SecretProperty{
//   					Name: jsii.String("name"),
//   					ValueFrom: jsii.String("valueFrom"),
//   				},
//   			},
//   		},
//   		Namespace: jsii.String("namespace"),
//   		Services: []interface{}{
//   			&ServiceConnectServiceProperty{
//   				ClientAliases: []interface{}{
//   					&ServiceConnectClientAliasProperty{
//   						DnsName: jsii.String("dnsName"),
//   						Port: jsii.Number(123),
//   						TestTrafficRules: &ServiceConnectTestTrafficRulesProperty{
//   							Header: &ServiceConnectTestTrafficRulesHeaderProperty{
//   								Name: jsii.String("name"),
//   								Value: &ServiceConnectTestTrafficRulesHeaderValueProperty{
//   									Exact: jsii.String("exact"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				DiscoveryName: jsii.String("discoveryName"),
//   				IngressPortOverride: jsii.Number(123),
//   				PortName: jsii.String("portName"),
//   				Timeout: &TimeoutConfigurationProperty{
//   					IdleTimeoutSeconds: jsii.Number(123),
//   					PerRequestTimeoutSeconds: jsii.Number(123),
//   				},
//   				Tls: &ServiceConnectTlsConfigurationProperty{
//   					IssuerCertificateAuthority: &ServiceConnectTlsCertificateAuthorityProperty{
//   						AwsPcaAuthorityArn: jsii.String("awsPcaAuthorityArn"),
//   					},
//   					KmsKey: jsii.String("kmsKey"),
//   					RoleArn: jsii.String("roleArn"),
//   				},
//   			},
//   		},
//   	},
//   	ServiceName: jsii.String("serviceName"),
//   	ServiceRegistries: []interface{}{
//   		&ServiceRegistryProperty{
//   			ContainerName: jsii.String("containerName"),
//   			ContainerPort: jsii.Number(123),
//   			Port: jsii.Number(123),
//   			RegistryArn: jsii.String("registryArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TaskDefinition: jsii.String("taskDefinition"),
//   	VolumeConfigurations: []interface{}{
//   		&ServiceVolumeConfigurationProperty{
//   			ManagedEbsVolume: &ServiceManagedEBSVolumeConfigurationProperty{
//   				Encrypted: jsii.Boolean(false),
//   				FilesystemType: jsii.String("filesystemType"),
//   				Iops: jsii.Number(123),
//   				KmsKeyId: jsii.String("kmsKeyId"),
//   				RoleArn: jsii.String("roleArn"),
//   				SizeInGiB: jsii.Number(123),
//   				SnapshotId: jsii.String("snapshotId"),
//   				TagSpecifications: []interface{}{
//   					&EBSTagSpecificationProperty{
//   						PropagateTags: jsii.String("propagateTags"),
//   						ResourceType: jsii.String("resourceType"),
//   						Tags: []CfnTag{
//   							&CfnTag{
//   								Key: jsii.String("key"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   					},
//   				},
//   				Throughput: jsii.Number(123),
//   				VolumeInitializationRate: jsii.Number(123),
//   				VolumeType: jsii.String("volumeType"),
//   			},
//   			Name: jsii.String("name"),
//   		},
//   	},
//   	VpcLatticeConfigurations: []interface{}{
//   		&VpcLatticeConfigurationProperty{
//   			PortName: jsii.String("portName"),
//   			RoleArn: jsii.String("roleArn"),
//   			TargetGroupArn: jsii.String("targetGroupArn"),
//   		},
//   	},
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecs-service.html
//
type CfnServicePropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnServiceMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnServicePropsMixin
type jsiiProxy_CfnServicePropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnServicePropsMixin) Props() *CfnServiceMixinProps {
	var returns *CfnServiceMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnServicePropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ECS::Service`.
func NewCfnServicePropsMixin(props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) CfnServicePropsMixin {
	_init_.Initialize()

	if err := validateNewCfnServicePropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnServicePropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ECS::Service`.
func NewCfnServicePropsMixin_Override(c CfnServicePropsMixin, props *CfnServiceMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnServicePropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnServicePropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnServicePropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_ecs.mixins.CfnServicePropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnServicePropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
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

func (c *jsiiProxy_CfnServicePropsMixin) Supports(construct constructs.IConstruct) *bool {
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


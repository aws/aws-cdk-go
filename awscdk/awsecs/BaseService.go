package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

// The base class for Ec2Service and FargateService services.
//
// Example:
//   import ecs "github.com/aws/aws-cdk-go/awscdk"
//
//
//   service := ecs.BaseService_FromServiceArnWithCluster(this, jsii.String("EcsService"), jsii.String("arn:aws:ecs:us-east-1:123456789012:service/myClusterName/myServiceName"))
//   pipeline := codepipeline.NewPipeline(this, jsii.String("MyPipeline"))
//   buildOutput := codepipeline.NewArtifact()
//   // add source and build stages to the pipeline as usual...
//   deployStage := pipeline.AddStage(&StageOptions{
//   	StageName: jsii.String("Deploy"),
//   	Actions: []iAction{
//   		codepipeline_actions.NewEcsDeployAction(&EcsDeployActionProps{
//   			ActionName: jsii.String("DeployAction"),
//   			Service: service,
//   			Input: buildOutput,
//   		}),
//   	},
//   })
//
type BaseService interface {
	awscdk.Resource
	IBaseService
	awselasticloadbalancing.ILoadBalancerTarget
	awselasticloadbalancingv2.IApplicationLoadBalancerTarget
	awselasticloadbalancingv2.INetworkLoadBalancerTarget
	// The details of the AWS Cloud Map service.
	CloudmapService() awsservicediscovery.Service
	SetCloudmapService(val awsservicediscovery.Service)
	// The CloudMap service created for this service, if any.
	CloudMapService() awsservicediscovery.IService
	// The cluster that hosts the service.
	Cluster() ICluster
	// The security groups which manage the allowed network traffic for the service.
	Connections() awsec2.Connections
	// The deployment alarms property - this will be rendered directly and lazily as the CfnService.alarms property.
	DeploymentAlarms() *CfnService_DeploymentAlarmsProperty
	SetDeploymentAlarms(val *CfnService_DeploymentAlarmsProperty)
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	LoadBalancers() *[]*CfnService_LoadBalancerProperty
	SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty)
	// A list of Elastic Load Balancing load balancer objects, containing the load balancer name, the container name (as it appears in a container definition), and the container port to access from the load balancer.
	NetworkConfiguration() *CfnService_NetworkConfigurationProperty
	SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty)
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The Amazon Resource Name (ARN) of the service.
	ServiceArn() *string
	// The name of the service.
	ServiceName() *string
	// The details of the service discovery registries to assign to this service.
	//
	// For more information, see Service Discovery.
	ServiceRegistries() *[]*CfnService_ServiceRegistryProperty
	SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty)
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The task definition to use for tasks in the service.
	TaskDefinition() TaskDefinition
	// Adds a volume to the Service.
	AddVolume(volume ServiceManagedVolume)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Associates this service with a CloudMap service.
	AssociateCloudMapService(options *AssociateCloudMapServiceOptions)
	// This method is called to attach this service to an Application Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// Registers the service as a target of a Classic Load Balancer (CLB).
	//
	// Don't call this. Call `loadBalancer.addTarget()` instead.
	AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer)
	// This method is called to attach this service to a Network Load Balancer.
	//
	// Don't call this function directly. Instead, call `listener.addTargets()`
	// to add this service to a load balancer.
	AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps
	// An attribute representing the minimum and maximum task count for an AutoScalingGroup.
	AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount
	// This method is called to create a networkConfiguration.
	ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup)
	// Enable CloudMap service discovery for the service.
	//
	// Returns: The created CloudMap service.
	EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service
	// Enable Deployment Alarms which take advantage of arbitrary alarms and configure them after service initialization.
	//
	// If you have already enabled deployment alarms, this function can be used to tell ECS about additional alarms that
	// should interrupt a deployment.
	//
	// New alarms specified in subsequent calls of this function will be appended to the existing list of alarms.
	//
	// The same Alarm Behavior must be used on all deployment alarms. If you specify different AlarmBehavior values in
	// multiple calls to this function, or the Alarm Behavior used here doesn't match the one used in the service
	// constructor, an error will be thrown.
	//
	// If the alarm's metric references the service, you cannot pass `Alarm.alarmName` here. That will cause a circular
	// dependency between the service and its deployment alarm. See this package's README for options to alarm on service
	// metrics, and avoid this circular dependency.
	EnableDeploymentAlarms(alarmNames *[]*string, options *DeploymentAlarmOptions)
	// Enable Service Connect on this service.
	EnableServiceConnect(config *ServiceConnectProps)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Return a load balancing target for a specific container and port.
	//
	// Use this function to create a load balancer target if you want to load balance to
	// another container than the first essential container or the first mapped port on
	// the container.
	//
	// Use the return value of this function where you would normally use a load balancer
	// target, instead of the `Service` object itself.
	//
	// Example:
	//   var listener applicationListener
	//   var service baseService
	//
	//   listener.AddTargets(jsii.String("ECS"), &AddApplicationTargetsProps{
	//   	Port: jsii.Number(80),
	//   	Targets: []iApplicationLoadBalancerTarget{
	//   		service.LoadBalancerTarget(&LoadBalancerTargetOptions{
	//   			ContainerName: jsii.String("MyContainer"),
	//   			ContainerPort: jsii.Number(1234),
	//   		}),
	//   	},
	//   })
	//
	LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget
	// This method returns the specified CloudWatch metric name for this service.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's CPU utilization.
	// Default: average over 5 minutes.
	//
	MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// This method returns the CloudWatch metric for this service's memory utilization.
	// Default: average over 5 minutes.
	//
	MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Use this function to create all load balancer targets to be registered in this service, add them to target groups, and attach target groups to listeners accordingly.
	//
	// Alternatively, you can use `listener.addTargets()` to create targets and add them to target groups.
	//
	// Example:
	//   var listener applicationListener
	//   var service baseService
	//
	//   service.RegisterLoadBalancerTargets(&EcsTarget{
	//   	ContainerName: jsii.String("web"),
	//   	ContainerPort: jsii.Number(80),
	//   	NewTargetGroupId: jsii.String("ECS"),
	//   	Listener: ecs.ListenerConfig_ApplicationListener(listener, &AddApplicationTargetsProps{
	//   		Protocol: elbv2.ApplicationProtocol_HTTPS,
	//   	}),
	//   })
	//
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for BaseService
type jsiiProxy_BaseService struct {
	internal.Type__awscdkResource
	jsiiProxy_IBaseService
	internal.Type__awselasticloadbalancingILoadBalancerTarget
	internal.Type__awselasticloadbalancingv2IApplicationLoadBalancerTarget
	internal.Type__awselasticloadbalancingv2INetworkLoadBalancerTarget
}

func (j *jsiiProxy_BaseService) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) DeploymentAlarms() *CfnService_DeploymentAlarmsProperty {
	var returns *CfnService_DeploymentAlarmsProperty
	_jsii_.Get(
		j,
		"deploymentAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseService) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the BaseService class.
func NewBaseService_Override(b BaseService, scope constructs.Construct, id *string, props *BaseServiceProps, additionalProps interface{}, taskDefinition TaskDefinition) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.BaseService",
		[]interface{}{scope, id, props, additionalProps, taskDefinition},
		b,
	)
}

func (j *jsiiProxy_BaseService)SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_BaseService)SetDeploymentAlarms(val *CfnService_DeploymentAlarmsProperty) {
	if err := j.validateSetDeploymentAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentAlarms",
		val,
	)
}

func (j *jsiiProxy_BaseService)SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_BaseService)SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_BaseService)SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	if err := j.validateSetServiceRegistriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func BaseService_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	if err := validateBaseService_FromServiceArnWithClusterParameters(scope, id, serviceArn); err != nil {
		panic(err)
	}
	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"fromServiceArnWithCluster",
		[]interface{}{scope, id, serviceArn},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func BaseService_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBaseService_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func BaseService_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBaseService_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func BaseService_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateBaseService_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.BaseService",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) AddVolume(volume ServiceManagedVolume) {
	if err := b.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addVolume",
		[]interface{}{volume},
	)
}

func (b *jsiiProxy_BaseService) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := b.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (b *jsiiProxy_BaseService) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	if err := b.validateAssociateCloudMapServiceParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

func (b *jsiiProxy_BaseService) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := b.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		b,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	if err := b.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

func (b *jsiiProxy_BaseService) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := b.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		b,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	if err := b.validateAutoScaleTaskCountParameters(props); err != nil {
		panic(err)
	}
	var returns ScalableTaskCount

	_jsii_.Invoke(
		b,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	if err := b.validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc, vpcSubnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

func (b *jsiiProxy_BaseService) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	if err := b.validateEnableCloudMapParameters(options); err != nil {
		panic(err)
	}
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		b,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) EnableDeploymentAlarms(alarmNames *[]*string, options *DeploymentAlarmOptions) {
	if err := b.validateEnableDeploymentAlarmsParameters(alarmNames, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"enableDeploymentAlarms",
		[]interface{}{alarmNames, options},
	)
}

func (b *jsiiProxy_BaseService) EnableServiceConnect(config *ServiceConnectProps) {
	if err := b.validateEnableServiceConnectParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"enableServiceConnect",
		[]interface{}{config},
	)
}

func (b *jsiiProxy_BaseService) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := b.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) GetResourceNameAttribute(nameAttr *string) *string {
	if err := b.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	if err := b.validateLoadBalancerTargetParameters(options); err != nil {
		panic(err)
	}
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		b,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := b.validateMetricMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		b,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseService) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	if err := b.validateRegisterLoadBalancerTargetsParameters(&targets); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		b,
		"registerLoadBalancerTargets",
		args,
	)
}

func (b *jsiiProxy_BaseService) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


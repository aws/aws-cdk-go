package awsecs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapplicationautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancing"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicediscovery"
	"github.com/aws/constructs-go/constructs/v10"
)

// This creates a service using the EC2 launch type on an ECS cluster.
//
// Example:
//   var vpc vpc
//
//
//   // Create an ECS cluster
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   })
//
//   // Add capacity to it
//   cluster.AddCapacity(jsii.String("DefaultAutoScalingGroupCapacity"), &AddCapacityOptions{
//   	InstanceType: ec2.NewInstanceType(jsii.String("t2.xlarge")),
//   	DesiredCapacity: jsii.Number(3),
//   })
//
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//
//   taskDefinition.AddContainer(jsii.String("DefaultContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("amazon/amazon-ecs-sample")),
//   	MemoryLimitMiB: jsii.Number(512),
//   })
//
//   // Instantiate an Amazon ECS Service
//   ecsService := ecs.NewEc2Service(this, jsii.String("Service"), &Ec2ServiceProps{
//   	Cluster: Cluster,
//   	TaskDefinition: TaskDefinition,
//   	MinHealthyPercent: jsii.Number(100),
//   })
//
type Ec2Service interface {
	BaseService
	IEc2Service
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
	// Adds one or more placement constraints to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Constraints](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-constraints.html).
	AddPlacementConstraints(constraints ...PlacementConstraint)
	// Adds one or more placement strategies to use for tasks in the service.
	//
	// For more information, see
	// [Amazon ECS Task Placement Strategies](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/task-placement-strategies.html).
	AddPlacementStrategies(strategies ...PlacementStrategy)
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
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   listener.addTargets('ECS', {
	//     port: 80,
	//     targets: [service.loadBalancerTarget({
	//       containerName: 'MyContainer',
	//       containerPort: 1234,
	//     })],
	//   });
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
	//   declare const listener: elbv2.ApplicationListener;
	//   declare const service: ecs.BaseService;
	//   service.registerLoadBalancerTargets(
	//     {
	//       containerName: 'web',
	//       containerPort: 80,
	//       newTargetGroupId: 'ECS',
	//       listener: ecs.ListenerConfig.applicationListener(listener, {
	//         protocol: elbv2.ApplicationProtocol.HTTPS
	//       }),
	//     },
	//   )
	//
	RegisterLoadBalancerTargets(targets ...*EcsTarget)
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for Ec2Service
type jsiiProxy_Ec2Service struct {
	jsiiProxy_BaseService
	jsiiProxy_IEc2Service
}

func (j *jsiiProxy_Ec2Service) CloudmapService() awsservicediscovery.Service {
	var returns awsservicediscovery.Service
	_jsii_.Get(
		j,
		"cloudmapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) CloudMapService() awsservicediscovery.IService {
	var returns awsservicediscovery.IService
	_jsii_.Get(
		j,
		"cloudMapService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Cluster() ICluster {
	var returns ICluster
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) DeploymentAlarms() *CfnService_DeploymentAlarmsProperty {
	var returns *CfnService_DeploymentAlarmsProperty
	_jsii_.Get(
		j,
		"deploymentAlarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) LoadBalancers() *[]*CfnService_LoadBalancerProperty {
	var returns *[]*CfnService_LoadBalancerProperty
	_jsii_.Get(
		j,
		"loadBalancers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) NetworkConfiguration() *CfnService_NetworkConfigurationProperty {
	var returns *CfnService_NetworkConfigurationProperty
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) ServiceRegistries() *[]*CfnService_ServiceRegistryProperty {
	var returns *[]*CfnService_ServiceRegistryProperty
	_jsii_.Get(
		j,
		"serviceRegistries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Service) TaskDefinition() TaskDefinition {
	var returns TaskDefinition
	_jsii_.Get(
		j,
		"taskDefinition",
		&returns,
	)
	return returns
}


// Constructs a new instance of the Ec2Service class.
func NewEc2Service(scope constructs.Construct, id *string, props *Ec2ServiceProps) Ec2Service {
	_init_.Initialize()

	if err := validateNewEc2ServiceParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Service{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Constructs a new instance of the Ec2Service class.
func NewEc2Service_Override(e Ec2Service, scope constructs.Construct, id *string, props *Ec2ServiceProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		[]interface{}{scope, id, props},
		e,
	)
}

func (j *jsiiProxy_Ec2Service)SetCloudmapService(val awsservicediscovery.Service) {
	_jsii_.Set(
		j,
		"cloudmapService",
		val,
	)
}

func (j *jsiiProxy_Ec2Service)SetDeploymentAlarms(val *CfnService_DeploymentAlarmsProperty) {
	if err := j.validateSetDeploymentAlarmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deploymentAlarms",
		val,
	)
}

func (j *jsiiProxy_Ec2Service)SetLoadBalancers(val *[]*CfnService_LoadBalancerProperty) {
	if err := j.validateSetLoadBalancersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancers",
		val,
	)
}

func (j *jsiiProxy_Ec2Service)SetNetworkConfiguration(val *CfnService_NetworkConfigurationProperty) {
	if err := j.validateSetNetworkConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkConfiguration",
		val,
	)
}

func (j *jsiiProxy_Ec2Service)SetServiceRegistries(val *[]*CfnService_ServiceRegistryProperty) {
	if err := j.validateSetServiceRegistriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceRegistries",
		val,
	)
}

// Imports from the specified service ARN.
func Ec2Service_FromEc2ServiceArn(scope constructs.Construct, id *string, ec2ServiceArn *string) IEc2Service {
	_init_.Initialize()

	if err := validateEc2Service_FromEc2ServiceArnParameters(scope, id, ec2ServiceArn); err != nil {
		panic(err)
	}
	var returns IEc2Service

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"fromEc2ServiceArn",
		[]interface{}{scope, id, ec2ServiceArn},
		&returns,
	)

	return returns
}

// Imports from the specified service attributes.
func Ec2Service_FromEc2ServiceAttributes(scope constructs.Construct, id *string, attrs *Ec2ServiceAttributes) IBaseService {
	_init_.Initialize()

	if err := validateEc2Service_FromEc2ServiceAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"fromEc2ServiceAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import an existing ECS/Fargate Service using the service cluster format.
//
// The format is the "new" format "arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name".
// See: https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids
//
func Ec2Service_FromServiceArnWithCluster(scope constructs.Construct, id *string, serviceArn *string) IBaseService {
	_init_.Initialize()

	if err := validateEc2Service_FromServiceArnWithClusterParameters(scope, id, serviceArn); err != nil {
		panic(err)
	}
	var returns IBaseService

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
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
func Ec2Service_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Service_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func Ec2Service_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEc2Service_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func Ec2Service_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateEc2Service_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func Ec2Service_PROPERTY_INJECTION_ID() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_ecs.Ec2Service",
		"PROPERTY_INJECTION_ID",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2Service) AddPlacementConstraints(constraints ...PlacementConstraint) {
	args := []interface{}{}
	for _, a := range constraints {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlacementConstraints",
		args,
	)
}

func (e *jsiiProxy_Ec2Service) AddPlacementStrategies(strategies ...PlacementStrategy) {
	args := []interface{}{}
	for _, a := range strategies {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"addPlacementStrategies",
		args,
	)
}

func (e *jsiiProxy_Ec2Service) AddVolume(volume ServiceManagedVolume) {
	if err := e.validateAddVolumeParameters(volume); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVolume",
		[]interface{}{volume},
	)
}

func (e *jsiiProxy_Ec2Service) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := e.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (e *jsiiProxy_Ec2Service) AssociateCloudMapService(options *AssociateCloudMapServiceOptions) {
	if err := e.validateAssociateCloudMapServiceParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"associateCloudMapService",
		[]interface{}{options},
	)
}

func (e *jsiiProxy_Ec2Service) AttachToApplicationTargetGroup(targetGroup awselasticloadbalancingv2.IApplicationTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := e.validateAttachToApplicationTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToApplicationTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) AttachToClassicLB(loadBalancer awselasticloadbalancing.LoadBalancer) {
	if err := e.validateAttachToClassicLBParameters(loadBalancer); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"attachToClassicLB",
		[]interface{}{loadBalancer},
	)
}

func (e *jsiiProxy_Ec2Service) AttachToNetworkTargetGroup(targetGroup awselasticloadbalancingv2.INetworkTargetGroup) *awselasticloadbalancingv2.LoadBalancerTargetProps {
	if err := e.validateAttachToNetworkTargetGroupParameters(targetGroup); err != nil {
		panic(err)
	}
	var returns *awselasticloadbalancingv2.LoadBalancerTargetProps

	_jsii_.Invoke(
		e,
		"attachToNetworkTargetGroup",
		[]interface{}{targetGroup},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) AutoScaleTaskCount(props *awsapplicationautoscaling.EnableScalingProps) ScalableTaskCount {
	if err := e.validateAutoScaleTaskCountParameters(props); err != nil {
		panic(err)
	}
	var returns ScalableTaskCount

	_jsii_.Invoke(
		e,
		"autoScaleTaskCount",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) ConfigureAwsVpcNetworkingWithSecurityGroups(vpc awsec2.IVpc, assignPublicIp *bool, vpcSubnets *awsec2.SubnetSelection, securityGroups *[]awsec2.ISecurityGroup) {
	if err := e.validateConfigureAwsVpcNetworkingWithSecurityGroupsParameters(vpc, vpcSubnets); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"configureAwsVpcNetworkingWithSecurityGroups",
		[]interface{}{vpc, assignPublicIp, vpcSubnets, securityGroups},
	)
}

func (e *jsiiProxy_Ec2Service) EnableCloudMap(options *CloudMapOptions) awsservicediscovery.Service {
	if err := e.validateEnableCloudMapParameters(options); err != nil {
		panic(err)
	}
	var returns awsservicediscovery.Service

	_jsii_.Invoke(
		e,
		"enableCloudMap",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) EnableDeploymentAlarms(alarmNames *[]*string, options *DeploymentAlarmOptions) {
	if err := e.validateEnableDeploymentAlarmsParameters(alarmNames, options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"enableDeploymentAlarms",
		[]interface{}{alarmNames, options},
	)
}

func (e *jsiiProxy_Ec2Service) EnableServiceConnect(config *ServiceConnectProps) {
	if err := e.validateEnableServiceConnectParameters(config); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"enableServiceConnect",
		[]interface{}{config},
	)
}

func (e *jsiiProxy_Ec2Service) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := e.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) GetResourceNameAttribute(nameAttr *string) *string {
	if err := e.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) LoadBalancerTarget(options *LoadBalancerTargetOptions) IEcsLoadBalancerTarget {
	if err := e.validateLoadBalancerTargetParameters(options); err != nil {
		panic(err)
	}
	var returns IEcsLoadBalancerTarget

	_jsii_.Invoke(
		e,
		"loadBalancerTarget",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) MetricCpuUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricCpuUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricCpuUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) MetricMemoryUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := e.validateMetricMemoryUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		e,
		"metricMemoryUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Service) RegisterLoadBalancerTargets(targets ...*EcsTarget) {
	if err := e.validateRegisterLoadBalancerTargetsParameters(&targets); err != nil {
		panic(err)
	}
	args := []interface{}{}
	for _, a := range targets {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		e,
		"registerLoadBalancerTargets",
		args,
	)
}

func (e *jsiiProxy_Ec2Service) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


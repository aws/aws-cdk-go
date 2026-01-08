package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudwatch"
)

// Construction properties for `ServerDeploymentGroup`.
//
// Example:
//   var alb ApplicationLoadBalancer
//
//   listener := alb.AddListener(jsii.String("Listener"), &BaseApplicationListenerProps{
//   	Port: jsii.Number(80),
//   })
//   targetGroup := listener.AddTargets(jsii.String("Fleet"), &AddApplicationTargetsProps{
//   	Port: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
//   	LoadBalancer: codedeploy.LoadBalancer_Application(targetGroup),
//   })
//
type ServerDeploymentGroupProps struct {
	// The CloudWatch alarms associated with this Deployment Group.
	//
	// CodeDeploy will stop (and optionally roll back)
	// a deployment if during it any of the alarms trigger.
	//
	// Alarms can also be added after the Deployment Group is created using the `#addAlarm` method.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/monitoring-create-alarms.html
	//
	// Default: [].
	//
	Alarms *[]interfacesawscloudwatch.IAlarmRef `field:"optional" json:"alarms" yaml:"alarms"`
	// The CodeDeploy EC2/on-premise Application this Deployment Group belongs to.
	// Default: - A new Application will be created.
	//
	Application IServerApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	// Default: - default AutoRollbackConfig.
	//
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The auto-scaling groups belonging to this Deployment Group.
	//
	// Auto-scaling groups can also be added after the Deployment Group is created
	// using the `#addAutoScalingGroup` method.
	//
	// [disable-awslint:ref-via-interface] is needed because we update userdata
	// for ASGs to install the codedeploy agent.
	// Default: [].
	//
	AutoScalingGroups *[]awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// The EC2/on-premise Deployment Configuration to use for this Deployment Group.
	// Default: ServerDeploymentConfig#OneAtATime.
	//
	DeploymentConfig IServerDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	// Default: - An auto-generated name will be used.
	//
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// All EC2 instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Default: - No additional EC2 instances will be added to the Deployment Group.
	//
	Ec2InstanceTags InstanceTagSet `field:"optional" json:"ec2InstanceTags" yaml:"ec2InstanceTags"`
	// Whether to skip the step of checking CloudWatch alarms during the deployment process.
	// Default: - false.
	//
	IgnoreAlarmConfiguration *bool `field:"optional" json:"ignoreAlarmConfiguration" yaml:"ignoreAlarmConfiguration"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	// Default: false.
	//
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// If you've provided any auto-scaling groups with the `#autoScalingGroups` property, you can set this property to add User Data that installs the CodeDeploy agent on the instances.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install.html
	//
	// Default: true.
	//
	InstallAgent *bool `field:"optional" json:"installAgent" yaml:"installAgent"`
	// The load balancer to place in front of this Deployment Group.
	//
	// Can be created from either a classic Elastic Load Balancer,
	// or an Application Load Balancer / Network Load Balancer Target Group.
	// Default: - Deployment Group will not have a load balancer defined.
	//
	// Deprecated: - Use `loadBalancers` instead.
	LoadBalancer LoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// CodeDeploy supports the deployment to multiple load balancers.
	//
	// Specify either multiple Classic Load Balancers, or
	// Application Load Balancers / Network Load Balancers Target Groups.
	// Default: - Deployment Group will not have load balancers defined.
	//
	LoadBalancers *[]LoadBalancer `field:"optional" json:"loadBalancers" yaml:"loadBalancers"`
	// All on-premise instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	// Default: - No additional on-premise instances will be added to the Deployment Group.
	//
	OnPremiseInstanceTags InstanceTagSet `field:"optional" json:"onPremiseInstanceTags" yaml:"onPremiseInstanceTags"`
	// The service Role of this Deployment Group.
	// Default: - A new Role will be created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
	// Indicates whether the deployment group was configured to have CodeDeploy install a termination hook into an Auto Scaling group.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-auto-scaling.html#integrations-aws-auto-scaling-behaviors
	//
	// Default: - false.
	//
	TerminationHook *bool `field:"optional" json:"terminationHook" yaml:"terminationHook"`
}


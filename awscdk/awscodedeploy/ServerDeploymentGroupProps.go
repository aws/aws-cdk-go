package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// Construction properties for `ServerDeploymentGroup`.
//
// Example:
//   import elb "github.com/aws/aws-cdk-go/awscdk"
//
//   var lb loadBalancer
//
//   lb.AddListener(&LoadBalancerListener{
//   	ExternalPort: jsii.Number(80),
//   })
//
//   deploymentGroup := codedeploy.NewServerDeploymentGroup(this, jsii.String("DeploymentGroup"), &ServerDeploymentGroupProps{
//   	LoadBalancer: codedeploy.LoadBalancer_Classic(lb),
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
	Alarms *[]awscloudwatch.IAlarm `field:"optional" json:"alarms" yaml:"alarms"`
	// The CodeDeploy EC2/on-premise Application this Deployment Group belongs to.
	Application IServerApplication `field:"optional" json:"application" yaml:"application"`
	// The auto-rollback configuration for this Deployment Group.
	AutoRollback *AutoRollbackConfig `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// The auto-scaling groups belonging to this Deployment Group.
	//
	// Auto-scaling groups can also be added after the Deployment Group is created
	// using the `#addAutoScalingGroup` method.
	//
	// [disable-awslint:ref-via-interface] is needed because we update userdata
	// for ASGs to install the codedeploy agent.
	AutoScalingGroups *[]awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// The EC2/on-premise Deployment Configuration to use for this Deployment Group.
	DeploymentConfig IServerDeploymentConfig `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The physical, human-readable name of the CodeDeploy Deployment Group.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// All EC2 instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	Ec2InstanceTags InstanceTagSet `field:"optional" json:"ec2InstanceTags" yaml:"ec2InstanceTags"`
	// Whether to continue a deployment even if fetching the alarm status from CloudWatch failed.
	IgnorePollAlarmsFailure *bool `field:"optional" json:"ignorePollAlarmsFailure" yaml:"ignorePollAlarmsFailure"`
	// If you've provided any auto-scaling groups with the `#autoScalingGroups` property, you can set this property to add User Data that installs the CodeDeploy agent on the instances.
	// See: https://docs.aws.amazon.com/codedeploy/latest/userguide/codedeploy-agent-operations-install.html
	//
	InstallAgent *bool `field:"optional" json:"installAgent" yaml:"installAgent"`
	// The load balancer to place in front of this Deployment Group.
	//
	// Can be created from either a classic Elastic Load Balancer,
	// or an Application Load Balancer / Network Load Balancer Target Group.
	LoadBalancer LoadBalancer `field:"optional" json:"loadBalancer" yaml:"loadBalancer"`
	// All on-premise instances matching the given set of tags when a deployment occurs will be added to this Deployment Group.
	OnPremiseInstanceTags InstanceTagSet `field:"optional" json:"onPremiseInstanceTags" yaml:"onPremiseInstanceTags"`
	// The service Role of this Deployment Group.
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


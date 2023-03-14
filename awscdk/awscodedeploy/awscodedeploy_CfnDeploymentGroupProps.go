package awscodedeploy

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnDeploymentGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnDeploymentGroupProps := &cfnDeploymentGroupProps{
//   	applicationName: jsii.String("applicationName"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	alarmConfiguration: &alarmConfigurationProperty{
//   		alarms: []interface{}{
//   			&alarmProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		enabled: jsii.Boolean(false),
//   		ignorePollAlarmFailure: jsii.Boolean(false),
//   	},
//   	autoRollbackConfiguration: &autoRollbackConfigurationProperty{
//   		enabled: jsii.Boolean(false),
//   		events: []*string{
//   			jsii.String("events"),
//   		},
//   	},
//   	autoScalingGroups: []*string{
//   		jsii.String("autoScalingGroups"),
//   	},
//   	blueGreenDeploymentConfiguration: &blueGreenDeploymentConfigurationProperty{
//   		deploymentReadyOption: &deploymentReadyOptionProperty{
//   			actionOnTimeout: jsii.String("actionOnTimeout"),
//   			waitTimeInMinutes: jsii.Number(123),
//   		},
//   		greenFleetProvisioningOption: &greenFleetProvisioningOptionProperty{
//   			action: jsii.String("action"),
//   		},
//   		terminateBlueInstancesOnDeploymentSuccess: &blueInstanceTerminationOptionProperty{
//   			action: jsii.String("action"),
//   			terminationWaitTimeInMinutes: jsii.Number(123),
//   		},
//   	},
//   	deployment: &deploymentProperty{
//   		revision: &revisionLocationProperty{
//   			gitHubLocation: &gitHubLocationProperty{
//   				commitId: jsii.String("commitId"),
//   				repository: jsii.String("repository"),
//   			},
//   			revisionType: jsii.String("revisionType"),
//   			s3Location: &s3LocationProperty{
//   				bucket: jsii.String("bucket"),
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				bundleType: jsii.String("bundleType"),
//   				eTag: jsii.String("eTag"),
//   				version: jsii.String("version"),
//   			},
//   		},
//
//   		// the properties below are optional
//   		description: jsii.String("description"),
//   		ignoreApplicationStopFailures: jsii.Boolean(false),
//   	},
//   	deploymentConfigName: jsii.String("deploymentConfigName"),
//   	deploymentGroupName: jsii.String("deploymentGroupName"),
//   	deploymentStyle: &deploymentStyleProperty{
//   		deploymentOption: jsii.String("deploymentOption"),
//   		deploymentType: jsii.String("deploymentType"),
//   	},
//   	ec2TagFilters: []interface{}{
//   		&eC2TagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	ec2TagSet: &eC2TagSetProperty{
//   		ec2TagSetList: []interface{}{
//   			&eC2TagSetListObjectProperty{
//   				ec2TagGroup: []interface{}{
//   					&eC2TagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ecsServices: []interface{}{
//   		&eCSServiceProperty{
//   			clusterName: jsii.String("clusterName"),
//   			serviceName: jsii.String("serviceName"),
//   		},
//   	},
//   	loadBalancerInfo: &loadBalancerInfoProperty{
//   		elbInfoList: []interface{}{
//   			&eLBInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupInfoList: []interface{}{
//   			&targetGroupInfoProperty{
//   				name: jsii.String("name"),
//   			},
//   		},
//   		targetGroupPairInfoList: []interface{}{
//   			&targetGroupPairInfoProperty{
//   				prodTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   				targetGroups: []interface{}{
//   					&targetGroupInfoProperty{
//   						name: jsii.String("name"),
//   					},
//   				},
//   				testTrafficRoute: &trafficRouteProperty{
//   					listenerArns: []*string{
//   						jsii.String("listenerArns"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	onPremisesInstanceTagFilters: []interface{}{
//   		&tagFilterProperty{
//   			key: jsii.String("key"),
//   			type: jsii.String("type"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	onPremisesTagSet: &onPremisesTagSetProperty{
//   		onPremisesTagSetList: []interface{}{
//   			&onPremisesTagSetListObjectProperty{
//   				onPremisesTagGroup: []interface{}{
//   					&tagFilterProperty{
//   						key: jsii.String("key"),
//   						type: jsii.String("type"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	outdatedInstancesStrategy: jsii.String("outdatedInstancesStrategy"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	triggerConfigurations: []interface{}{
//   		&triggerConfigProperty{
//   			triggerEvents: []*string{
//   				jsii.String("triggerEvents"),
//   			},
//   			triggerName: jsii.String("triggerName"),
//   			triggerTargetArn: jsii.String("triggerTargetArn"),
//   		},
//   	},
//   }
//
type CfnDeploymentGroupProps struct {
	// The name of an existing CodeDeploy application to associate this deployment group with.
	ApplicationName *string `field:"required" json:"applicationName" yaml:"applicationName"`
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	//
	// For more information, see [Create a Service Role for AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/getting-started-create-service-role.html) in the *AWS CodeDeploy User Guide* .
	//
	// > In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	AlarmConfiguration interface{} `field:"optional" json:"alarmConfiguration" yaml:"alarmConfiguration"`
	// Information about the automatic rollback configuration that is associated with the deployment group.
	//
	// If you specify this property, don't specify the `Deployment` property.
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	//
	// Duplicates are not allowed.
	AutoScalingGroups *[]*string `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration interface{} `field:"optional" json:"blueGreenDeploymentConfiguration" yaml:"blueGreenDeploymentConfiguration"`
	// The application revision to deploy to this deployment group.
	//
	// If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the `AutoRollbackConfiguration` property.
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// A deployment configuration name or a predefined configuration name.
	//
	// With predefined configurations, you can deploy application revisions to one instance at a time ( `CodeDeployDefault.OneAtATime` ), half of the instances at a time ( `CodeDeployDefault.HalfAtATime` ), or all the instances at once ( `CodeDeployDefault.AllAtOnce` ). For more information and valid values, see [Working with Deployment Configurations](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html) in the *AWS CodeDeploy User Guide* .
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// A name for the deployment group.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the deployment group name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	//
	// If you specify this property with a blue/green deployment type, don't specify the `AutoScalingGroups` , `LoadBalancerInfo` , or `Deployment` properties.
	//
	// > For blue/green deployments, AWS CloudFormation supports deployments on Lambda compute platforms only. You can perform Amazon ECS blue/green deployments using `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using AWS CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
	DeploymentStyle interface{} `field:"optional" json:"deploymentStyle" yaml:"deploymentStyle"`
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed.
	//
	// You can specify `EC2TagFilters` or `Ec2TagSet` , but not both.
	Ec2TagFilters interface{} `field:"optional" json:"ec2TagFilters" yaml:"ec2TagFilters"`
	// Information about groups of tags applied to Amazon EC2 instances.
	//
	// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as `ec2TagFilter` .
	Ec2TagSet interface{} `field:"optional" json:"ec2TagSet" yaml:"ec2TagSet"`
	// The target Amazon ECS services in the deployment group.
	//
	// This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format `<clustername>:<servicename>` .
	EcsServices interface{} `field:"optional" json:"ecsServices" yaml:"ecsServices"`
	// Information about the load balancer to use in a deployment.
	//
	// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
	LoadBalancerInfo interface{} `field:"optional" json:"loadBalancerInfo" yaml:"loadBalancerInfo"`
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. To register on-premises instances with CodeDeploy , see [Working with On-Premises Instances for CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* . Duplicates are not allowed.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesInstanceTagFilters interface{} `field:"optional" json:"onPremisesInstanceTagFilters" yaml:"onPremisesInstanceTagFilters"`
	// Information about groups of tags applied to on-premises instances.
	//
	// The deployment group includes only on-premises instances identified by all the tag groups.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	OnPremisesTagSet interface{} `field:"optional" json:"onPremisesTagSet" yaml:"onPremisesTagSet"`
	// Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.
	//
	// If this option is set to `UPDATE` or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances.
	//
	// If this option is set to `IGNORE` , CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions.
	OutdatedInstancesStrategy *string `field:"optional" json:"outdatedInstancesStrategy" yaml:"outdatedInstancesStrategy"`
	// The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Information about triggers associated with the deployment group.
	//
	// Duplicates are not allowed.
	TriggerConfigurations interface{} `field:"optional" json:"triggerConfigurations" yaml:"triggerConfigurations"`
}


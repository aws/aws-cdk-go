package previewawscodedeploymixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDeploymentGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDeploymentGroupMixinProps := &CfnDeploymentGroupMixinProps{
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
//
type CfnDeploymentGroupMixinProps struct {
	// Information about the Amazon CloudWatch alarms that are associated with the deployment group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-alarmconfiguration
	//
	AlarmConfiguration interface{} `field:"optional" json:"alarmConfiguration" yaml:"alarmConfiguration"`
	// The name of an existing CodeDeploy application to associate this deployment group with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-applicationname
	//
	ApplicationName *string `field:"optional" json:"applicationName" yaml:"applicationName"`
	// Information about the automatic rollback configuration that is associated with the deployment group.
	//
	// If you specify this property, don't specify the `Deployment` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autorollbackconfiguration
	//
	AutoRollbackConfiguration interface{} `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// A list of associated Auto Scaling groups that CodeDeploy automatically deploys revisions to when new instances are created.
	//
	// Duplicates are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autoscalinggroups
	//
	AutoScalingGroups *[]*string `field:"optional" json:"autoScalingGroups" yaml:"autoScalingGroups"`
	// Information about blue/green deployment options for a deployment group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-bluegreendeploymentconfiguration
	//
	BlueGreenDeploymentConfiguration interface{} `field:"optional" json:"blueGreenDeploymentConfiguration" yaml:"blueGreenDeploymentConfiguration"`
	// The application revision to deploy to this deployment group.
	//
	// If you specify this property, your target application revision is deployed as soon as the provisioning process is complete. If you specify this property, don't specify the `AutoRollbackConfiguration` property.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deployment
	//
	Deployment interface{} `field:"optional" json:"deployment" yaml:"deployment"`
	// A deployment configuration name or a predefined configuration name.
	//
	// With predefined configurations, you can deploy application revisions to one instance at a time ( `CodeDeployDefault.OneAtATime` ), half of the instances at a time ( `CodeDeployDefault.HalfAtATime` ), or all the instances at once ( `CodeDeployDefault.AllAtOnce` ). For more information and valid values, see [Working with Deployment Configurations](https://docs.aws.amazon.com/codedeploy/latest/userguide/deployment-configurations.html) in the *AWS CodeDeploy User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentconfigname
	//
	DeploymentConfigName *string `field:"optional" json:"deploymentConfigName" yaml:"deploymentConfigName"`
	// A name for the deployment group.
	//
	// If you don't specify a name, CloudFormation generates a unique physical ID and uses that ID for the deployment group name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentgroupname
	//
	DeploymentGroupName *string `field:"optional" json:"deploymentGroupName" yaml:"deploymentGroupName"`
	// Attributes that determine the type of deployment to run and whether to route deployment traffic behind a load balancer.
	//
	// If you specify this property with a blue/green deployment type, don't specify the `AutoScalingGroups` , `LoadBalancerInfo` , or `Deployment` properties.
	//
	// > For blue/green deployments, CloudFormation supports deployments on Lambda compute platforms only. You can perform Amazon ECS blue/green deployments using `AWS::CodeDeploy::BlueGreen` hook. See [Perform Amazon ECS blue/green deployments through CodeDeploy using CloudFormation](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/blue-green.html) for more information.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentstyle
	//
	DeploymentStyle interface{} `field:"optional" json:"deploymentStyle" yaml:"deploymentStyle"`
	// The Amazon EC2 tags that are already applied to Amazon EC2 instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all Amazon EC2 instances identified by any of the tags you specify in this deployment group. Duplicates are not allowed.
	//
	// You can specify `EC2TagFilters` or `Ec2TagSet` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagfilters
	//
	Ec2TagFilters interface{} `field:"optional" json:"ec2TagFilters" yaml:"ec2TagFilters"`
	// Information about groups of tags applied to Amazon EC2 instances.
	//
	// The deployment group includes only Amazon EC2 instances identified by all the tag groups. Cannot be used in the same call as `ec2TagFilter` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagset
	//
	Ec2TagSet interface{} `field:"optional" json:"ec2TagSet" yaml:"ec2TagSet"`
	// The target Amazon ECS services in the deployment group.
	//
	// This applies only to deployment groups that use the Amazon ECS compute platform. A target Amazon ECS service is specified as an Amazon ECS cluster and service name pair using the format `<clustername>:<servicename>` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ecsservices
	//
	EcsServices interface{} `field:"optional" json:"ecsServices" yaml:"ecsServices"`
	// Information about the load balancer to use in a deployment.
	//
	// For more information, see [Integrating CodeDeploy with Elastic Load Balancing](https://docs.aws.amazon.com/codedeploy/latest/userguide/integrations-aws-elastic-load-balancing.html) in the *AWS CodeDeploy User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-loadbalancerinfo
	//
	LoadBalancerInfo interface{} `field:"optional" json:"loadBalancerInfo" yaml:"loadBalancerInfo"`
	// The on-premises instance tags already applied to on-premises instances that you want to include in the deployment group.
	//
	// CodeDeploy includes all on-premises instances identified by any of the tags you specify in this deployment group. To register on-premises instances with CodeDeploy , see [Working with On-Premises Instances for CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/instances-on-premises.html) in the *AWS CodeDeploy User Guide* . Duplicates are not allowed.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisesinstancetagfilters
	//
	OnPremisesInstanceTagFilters interface{} `field:"optional" json:"onPremisesInstanceTagFilters" yaml:"onPremisesInstanceTagFilters"`
	// Information about groups of tags applied to on-premises instances.
	//
	// The deployment group includes only on-premises instances identified by all the tag groups.
	//
	// You can specify `OnPremisesInstanceTagFilters` or `OnPremisesInstanceTagSet` , but not both.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisestagset
	//
	OnPremisesTagSet interface{} `field:"optional" json:"onPremisesTagSet" yaml:"onPremisesTagSet"`
	// Indicates what happens when new Amazon EC2 instances are launched mid-deployment and do not receive the deployed application revision.
	//
	// If this option is set to `UPDATE` or is unspecified, CodeDeploy initiates one or more 'auto-update outdated instances' deployments to apply the deployed application revision to the new Amazon EC2 instances.
	//
	// If this option is set to `IGNORE` , CodeDeploy does not initiate a deployment to update the new Amazon EC2 instances. This may result in instances having different revisions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-outdatedinstancesstrategy
	//
	OutdatedInstancesStrategy *string `field:"optional" json:"outdatedInstancesStrategy" yaml:"outdatedInstancesStrategy"`
	// A service role Amazon Resource Name (ARN) that grants CodeDeploy permission to make calls to AWS services on your behalf.
	//
	// For more information, see [Create a Service Role for AWS CodeDeploy](https://docs.aws.amazon.com/codedeploy/latest/userguide/getting-started-create-service-role.html) in the *AWS CodeDeploy User Guide* .
	//
	// > In some cases, you might need to add a dependency on the service role's policy. For more information, see IAM role policy in [DependsOn Attribute](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-servicerolearn
	//
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// The metadata that you apply to CodeDeploy deployment groups to help you organize and categorize them.
	//
	// Each tag consists of a key and an optional value, both of which you define.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// Indicates whether the deployment group was configured to have CodeDeploy install a termination hook into an Auto Scaling group.
	//
	// For more information about the termination hook, see [How Amazon EC2 Auto Scaling works with CodeDeploy](https://docs.aws.amazon.com//codedeploy/latest/userguide/integrations-aws-auto-scaling.html#integrations-aws-auto-scaling-behaviors) in the *AWS CodeDeploy User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-terminationhookenabled
	//
	TerminationHookEnabled interface{} `field:"optional" json:"terminationHookEnabled" yaml:"terminationHookEnabled"`
	// Information about triggers associated with the deployment group.
	//
	// Duplicates are not allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-triggerconfigurations
	//
	TriggerConfigurations interface{} `field:"optional" json:"triggerConfigurations" yaml:"triggerConfigurations"`
}


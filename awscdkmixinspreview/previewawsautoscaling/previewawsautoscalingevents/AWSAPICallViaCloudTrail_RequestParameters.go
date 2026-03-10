package previewawsautoscalingevents


// Type definition for RequestParameters.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var lifecycleHookSpecificationList interface{}
//   var overrides interface{}
//   var scheduledUpdateGroupActions interface{}
//
//   requestParameters := &RequestParameters{
//   	AdjustmentType: []*string{
//   		jsii.String("adjustmentType"),
//   	},
//   	AutoScalingGroupName: []*string{
//   		jsii.String("autoScalingGroupName"),
//   	},
//   	AvailabilityZones: []*string{
//   		jsii.String("availabilityZones"),
//   	},
//   	BreachThreshold: []*string{
//   		jsii.String("breachThreshold"),
//   	},
//   	DefaultCooldown: []*string{
//   		jsii.String("defaultCooldown"),
//   	},
//   	DesiredCapacity: []*string{
//   		jsii.String("desiredCapacity"),
//   	},
//   	ForceDelete: []*string{
//   		jsii.String("forceDelete"),
//   	},
//   	Granularity: []*string{
//   		jsii.String("granularity"),
//   	},
//   	HealthCheckGracePeriod: []*string{
//   		jsii.String("healthCheckGracePeriod"),
//   	},
//   	HealthCheckType: []*string{
//   		jsii.String("healthCheckType"),
//   	},
//   	HonorCooldown: []*string{
//   		jsii.String("honorCooldown"),
//   	},
//   	ImageId: []*string{
//   		jsii.String("imageId"),
//   	},
//   	InstanceIds: []*string{
//   		jsii.String("instanceIds"),
//   	},
//   	InstanceType: []*string{
//   		jsii.String("instanceType"),
//   	},
//   	LaunchConfigurationName: []*string{
//   		jsii.String("launchConfigurationName"),
//   	},
//   	LaunchTemplate: &LaunchTemplate{
//   		LaunchTemplateName: []*string{
//   			jsii.String("launchTemplateName"),
//   		},
//   	},
//   	LifecycleHookSpecificationList: []interface{}{
//   		lifecycleHookSpecificationList,
//   	},
//   	LoadBalancerNames: []*string{
//   		jsii.String("loadBalancerNames"),
//   	},
//   	MaxSize: []*string{
//   		jsii.String("maxSize"),
//   	},
//   	Metrics: []*string{
//   		jsii.String("metrics"),
//   	},
//   	MetricValue: []*string{
//   		jsii.String("metricValue"),
//   	},
//   	MinSize: []*string{
//   		jsii.String("minSize"),
//   	},
//   	MixedInstancesPolicy: &MixedInstancesPolicy{
//   		InstancesDistribution: &InstancesDistribution{
//   			OnDemandAllocationStrategy: []*string{
//   				jsii.String("onDemandAllocationStrategy"),
//   			},
//   			OnDemandBaseCapacity: []*string{
//   				jsii.String("onDemandBaseCapacity"),
//   			},
//   			OnDemandPercentageAboveBaseCapacity: []*string{
//   				jsii.String("onDemandPercentageAboveBaseCapacity"),
//   			},
//   			SpotAllocationStrategy: []*string{
//   				jsii.String("spotAllocationStrategy"),
//   			},
//   			SpotInstancePools: []*string{
//   				jsii.String("spotInstancePools"),
//   			},
//   		},
//   		LaunchTemplate: &LaunchTemplate1{
//   			LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   				LaunchTemplateName: []*string{
//   					jsii.String("launchTemplateName"),
//   				},
//   				Version: []*string{
//   					jsii.String("version"),
//   				},
//   			},
//   			Overrides: []interface{}{
//   				overrides,
//   			},
//   		},
//   	},
//   	NewInstancesProtectedFromScaleIn: []*string{
//   		jsii.String("newInstancesProtectedFromScaleIn"),
//   	},
//   	NotificationTypes: []*string{
//   		jsii.String("notificationTypes"),
//   	},
//   	PolicyName: []*string{
//   		jsii.String("policyName"),
//   	},
//   	PolicyType: []*string{
//   		jsii.String("policyType"),
//   	},
//   	ProtectedFromScaleIn: []*string{
//   		jsii.String("protectedFromScaleIn"),
//   	},
//   	ScalingAdjustment: []*string{
//   		jsii.String("scalingAdjustment"),
//   	},
//   	ScheduledActionName: []*string{
//   		jsii.String("scheduledActionName"),
//   	},
//   	ScheduledActionNames: []*string{
//   		jsii.String("scheduledActionNames"),
//   	},
//   	ScheduledUpdateGroupActions: []interface{}{
//   		scheduledUpdateGroupActions,
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	ServiceLinkedRoleArn: []*string{
//   		jsii.String("serviceLinkedRoleArn"),
//   	},
//   	SpotPrice: []*string{
//   		jsii.String("spotPrice"),
//   	},
//   	StartTime: []*string{
//   		jsii.String("startTime"),
//   	},
//   	StepAdjustments: []RequestParametersItem1{
//   		&RequestParametersItem1{
//   			MetricIntervalLowerBound: []*string{
//   				jsii.String("metricIntervalLowerBound"),
//   			},
//   			ScalingAdjustment: []*string{
//   				jsii.String("scalingAdjustment"),
//   			},
//   		},
//   	},
//   	Tags: []RequestParametersItem{
//   		&RequestParametersItem{
//   			Key: []*string{
//   				jsii.String("key"),
//   			},
//   			PropagateAtLaunch: []*string{
//   				jsii.String("propagateAtLaunch"),
//   			},
//   			ResourceId: []*string{
//   				jsii.String("resourceId"),
//   			},
//   			ResourceType: []*string{
//   				jsii.String("resourceType"),
//   			},
//   			Value: []*string{
//   				jsii.String("value"),
//   			},
//   		},
//   	},
//   	TargetGroupArns: []*string{
//   		jsii.String("targetGroupArns"),
//   	},
//   	TargetTrackingConfiguration: &TargetTrackingConfiguration{
//   		CustomizedMetricSpecification: &CustomizedMetricSpecification{
//   			Dimensions: []CustomizedMetricSpecificationItem{
//   				&CustomizedMetricSpecificationItem{
//   					Name: []*string{
//   						jsii.String("name"),
//   					},
//   					Value: []*string{
//   						jsii.String("value"),
//   					},
//   				},
//   			},
//   			MetricName: []*string{
//   				jsii.String("metricName"),
//   			},
//   			Namespace: []*string{
//   				jsii.String("namespace"),
//   			},
//   			Statistic: []*string{
//   				jsii.String("statistic"),
//   			},
//   			Unit: []*string{
//   				jsii.String("unit"),
//   			},
//   		},
//   		PredefinedMetricSpecification: &PredefinedMetricSpecification{
//   			PredefinedMetricType: []*string{
//   				jsii.String("predefinedMetricType"),
//   			},
//   		},
//   		TargetValue: []*string{
//   			jsii.String("targetValue"),
//   		},
//   	},
//   	Time: []*string{
//   		jsii.String("time"),
//   	},
//   	TopicArn: []*string{
//   		jsii.String("topicArn"),
//   	},
//   	UserData: []*string{
//   		jsii.String("userData"),
//   	},
//   	VPcZoneIdentifier: []*string{
//   		jsii.String("vPcZoneIdentifier"),
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_RequestParameters struct {
	// adjustmentType property.
	//
	// Specify an array of string values to match this event if the actual value of adjustmentType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AdjustmentType *[]*string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// autoScalingGroupName property.
	//
	// Specify an array of string values to match this event if the actual value of autoScalingGroupName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AutoScalingGroupName *[]*string `field:"optional" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// availabilityZones property.
	//
	// Specify an array of string values to match this event if the actual value of availabilityZones is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// breachThreshold property.
	//
	// Specify an array of string values to match this event if the actual value of breachThreshold is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	BreachThreshold *[]*string `field:"optional" json:"breachThreshold" yaml:"breachThreshold"`
	// defaultCooldown property.
	//
	// Specify an array of string values to match this event if the actual value of defaultCooldown is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DefaultCooldown *[]*string `field:"optional" json:"defaultCooldown" yaml:"defaultCooldown"`
	// desiredCapacity property.
	//
	// Specify an array of string values to match this event if the actual value of desiredCapacity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	DesiredCapacity *[]*string `field:"optional" json:"desiredCapacity" yaml:"desiredCapacity"`
	// forceDelete property.
	//
	// Specify an array of string values to match this event if the actual value of forceDelete is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ForceDelete *[]*string `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// granularity property.
	//
	// Specify an array of string values to match this event if the actual value of granularity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Granularity *[]*string `field:"optional" json:"granularity" yaml:"granularity"`
	// healthCheckGracePeriod property.
	//
	// Specify an array of string values to match this event if the actual value of healthCheckGracePeriod is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HealthCheckGracePeriod *[]*string `field:"optional" json:"healthCheckGracePeriod" yaml:"healthCheckGracePeriod"`
	// healthCheckType property.
	//
	// Specify an array of string values to match this event if the actual value of healthCheckType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HealthCheckType *[]*string `field:"optional" json:"healthCheckType" yaml:"healthCheckType"`
	// honorCooldown property.
	//
	// Specify an array of string values to match this event if the actual value of honorCooldown is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	HonorCooldown *[]*string `field:"optional" json:"honorCooldown" yaml:"honorCooldown"`
	// imageId property.
	//
	// Specify an array of string values to match this event if the actual value of imageId is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ImageId *[]*string `field:"optional" json:"imageId" yaml:"imageId"`
	// instanceIds property.
	//
	// Specify an array of string values to match this event if the actual value of instanceIds is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// instanceType property.
	//
	// Specify an array of string values to match this event if the actual value of instanceType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launchConfigurationName property.
	//
	// Specify an array of string values to match this event if the actual value of launchConfigurationName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchConfigurationName *[]*string `field:"optional" json:"launchConfigurationName" yaml:"launchConfigurationName"`
	// launchTemplate property.
	//
	// Specify an array of string values to match this event if the actual value of launchTemplate is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LaunchTemplate *AWSAPICallViaCloudTrail_LaunchTemplate `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// lifecycleHookSpecificationList property.
	//
	// Specify an array of string values to match this event if the actual value of lifecycleHookSpecificationList is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LifecycleHookSpecificationList *[]interface{} `field:"optional" json:"lifecycleHookSpecificationList" yaml:"lifecycleHookSpecificationList"`
	// loadBalancerNames property.
	//
	// Specify an array of string values to match this event if the actual value of loadBalancerNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	LoadBalancerNames *[]*string `field:"optional" json:"loadBalancerNames" yaml:"loadBalancerNames"`
	// maxSize property.
	//
	// Specify an array of string values to match this event if the actual value of maxSize is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MaxSize *[]*string `field:"optional" json:"maxSize" yaml:"maxSize"`
	// metrics property.
	//
	// Specify an array of string values to match this event if the actual value of metrics is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Metrics *[]*string `field:"optional" json:"metrics" yaml:"metrics"`
	// metricValue property.
	//
	// Specify an array of string values to match this event if the actual value of metricValue is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MetricValue *[]*string `field:"optional" json:"metricValue" yaml:"metricValue"`
	// minSize property.
	//
	// Specify an array of string values to match this event if the actual value of minSize is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MinSize *[]*string `field:"optional" json:"minSize" yaml:"minSize"`
	// mixedInstancesPolicy property.
	//
	// Specify an array of string values to match this event if the actual value of mixedInstancesPolicy is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	MixedInstancesPolicy *AWSAPICallViaCloudTrail_MixedInstancesPolicy `field:"optional" json:"mixedInstancesPolicy" yaml:"mixedInstancesPolicy"`
	// newInstancesProtectedFromScaleIn property.
	//
	// Specify an array of string values to match this event if the actual value of newInstancesProtectedFromScaleIn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NewInstancesProtectedFromScaleIn *[]*string `field:"optional" json:"newInstancesProtectedFromScaleIn" yaml:"newInstancesProtectedFromScaleIn"`
	// notificationTypes property.
	//
	// Specify an array of string values to match this event if the actual value of notificationTypes is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	NotificationTypes *[]*string `field:"optional" json:"notificationTypes" yaml:"notificationTypes"`
	// policyName property.
	//
	// Specify an array of string values to match this event if the actual value of policyName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyName *[]*string `field:"optional" json:"policyName" yaml:"policyName"`
	// policyType property.
	//
	// Specify an array of string values to match this event if the actual value of policyType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	PolicyType *[]*string `field:"optional" json:"policyType" yaml:"policyType"`
	// protectedFromScaleIn property.
	//
	// Specify an array of string values to match this event if the actual value of protectedFromScaleIn is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ProtectedFromScaleIn *[]*string `field:"optional" json:"protectedFromScaleIn" yaml:"protectedFromScaleIn"`
	// scalingAdjustment property.
	//
	// Specify an array of string values to match this event if the actual value of scalingAdjustment is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScalingAdjustment *[]*string `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// scheduledActionName property.
	//
	// Specify an array of string values to match this event if the actual value of scheduledActionName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScheduledActionName *[]*string `field:"optional" json:"scheduledActionName" yaml:"scheduledActionName"`
	// scheduledActionNames property.
	//
	// Specify an array of string values to match this event if the actual value of scheduledActionNames is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScheduledActionNames *[]*string `field:"optional" json:"scheduledActionNames" yaml:"scheduledActionNames"`
	// scheduledUpdateGroupActions property.
	//
	// Specify an array of string values to match this event if the actual value of scheduledUpdateGroupActions is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ScheduledUpdateGroupActions *[]interface{} `field:"optional" json:"scheduledUpdateGroupActions" yaml:"scheduledUpdateGroupActions"`
	// securityGroups property.
	//
	// Specify an array of string values to match this event if the actual value of securityGroups is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// serviceLinkedRoleARN property.
	//
	// Specify an array of string values to match this event if the actual value of serviceLinkedRoleARN is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ServiceLinkedRoleArn *[]*string `field:"optional" json:"serviceLinkedRoleArn" yaml:"serviceLinkedRoleArn"`
	// spotPrice property.
	//
	// Specify an array of string values to match this event if the actual value of spotPrice is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SpotPrice *[]*string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// startTime property.
	//
	// Specify an array of string values to match this event if the actual value of startTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StartTime *[]*string `field:"optional" json:"startTime" yaml:"startTime"`
	// stepAdjustments property.
	//
	// Specify an array of string values to match this event if the actual value of stepAdjustments is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	StepAdjustments *[]*AWSAPICallViaCloudTrail_RequestParametersItem1 `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
	// tags property.
	//
	// Specify an array of string values to match this event if the actual value of tags is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Tags *[]*AWSAPICallViaCloudTrail_RequestParametersItem `field:"optional" json:"tags" yaml:"tags"`
	// targetGroupARNs property.
	//
	// Specify an array of string values to match this event if the actual value of targetGroupARNs is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetGroupArns *[]*string `field:"optional" json:"targetGroupArns" yaml:"targetGroupArns"`
	// targetTrackingConfiguration property.
	//
	// Specify an array of string values to match this event if the actual value of targetTrackingConfiguration is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TargetTrackingConfiguration *AWSAPICallViaCloudTrail_TargetTrackingConfiguration `field:"optional" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// time property.
	//
	// Specify an array of string values to match this event if the actual value of time is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	Time *[]*string `field:"optional" json:"time" yaml:"time"`
	// topicARN property.
	//
	// Specify an array of string values to match this event if the actual value of topicARN is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	TopicArn *[]*string `field:"optional" json:"topicArn" yaml:"topicArn"`
	// userData property.
	//
	// Specify an array of string values to match this event if the actual value of userData is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserData *[]*string `field:"optional" json:"userData" yaml:"userData"`
	// vPCZoneIdentifier property.
	//
	// Specify an array of string values to match this event if the actual value of vPCZoneIdentifier is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	VPcZoneIdentifier *[]*string `field:"optional" json:"vPcZoneIdentifier" yaml:"vPcZoneIdentifier"`
}


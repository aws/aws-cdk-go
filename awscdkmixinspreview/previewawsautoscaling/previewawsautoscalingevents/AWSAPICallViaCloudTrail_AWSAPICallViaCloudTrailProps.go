package previewawsautoscalingevents

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Props type for aws.autoscaling@AWSAPICallViaCloudTrail event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var alarms interface{}
//   var failedScheduledActions interface{}
//   var failedScheduledUpdateGroupActions interface{}
//   var lifecycleHookSpecificationList interface{}
//   var overrides interface{}
//   var scheduledUpdateGroupActions interface{}
//
//   aWSAPICallViaCloudTrailProps := &AWSAPICallViaCloudTrailProps{
//   	AwsRegion: []*string{
//   		jsii.String("awsRegion"),
//   	},
//   	ErrorCode: []*string{
//   		jsii.String("errorCode"),
//   	},
//   	ErrorMessage: []*string{
//   		jsii.String("errorMessage"),
//   	},
//   	EventId: []*string{
//   		jsii.String("eventId"),
//   	},
//   	EventMetadata: &AWSEventMetadataProps{
//   		Region: []*string{
//   			jsii.String("region"),
//   		},
//   		Resources: []*string{
//   			jsii.String("resources"),
//   		},
//   		Version: []*string{
//   			jsii.String("version"),
//   		},
//   	},
//   	EventName: []*string{
//   		jsii.String("eventName"),
//   	},
//   	EventSource: []*string{
//   		jsii.String("eventSource"),
//   	},
//   	EventTime: []*string{
//   		jsii.String("eventTime"),
//   	},
//   	EventType: []*string{
//   		jsii.String("eventType"),
//   	},
//   	EventVersion: []*string{
//   		jsii.String("eventVersion"),
//   	},
//   	RequestId: []*string{
//   		jsii.String("requestId"),
//   	},
//   	RequestParameters: &RequestParameters{
//   		AdjustmentType: []*string{
//   			jsii.String("adjustmentType"),
//   		},
//   		AutoScalingGroupName: []*string{
//   			jsii.String("autoScalingGroupName"),
//   		},
//   		AvailabilityZones: []*string{
//   			jsii.String("availabilityZones"),
//   		},
//   		BreachThreshold: []*string{
//   			jsii.String("breachThreshold"),
//   		},
//   		DefaultCooldown: []*string{
//   			jsii.String("defaultCooldown"),
//   		},
//   		DesiredCapacity: []*string{
//   			jsii.String("desiredCapacity"),
//   		},
//   		ForceDelete: []*string{
//   			jsii.String("forceDelete"),
//   		},
//   		Granularity: []*string{
//   			jsii.String("granularity"),
//   		},
//   		HealthCheckGracePeriod: []*string{
//   			jsii.String("healthCheckGracePeriod"),
//   		},
//   		HealthCheckType: []*string{
//   			jsii.String("healthCheckType"),
//   		},
//   		HonorCooldown: []*string{
//   			jsii.String("honorCooldown"),
//   		},
//   		ImageId: []*string{
//   			jsii.String("imageId"),
//   		},
//   		InstanceIds: []*string{
//   			jsii.String("instanceIds"),
//   		},
//   		InstanceType: []*string{
//   			jsii.String("instanceType"),
//   		},
//   		LaunchConfigurationName: []*string{
//   			jsii.String("launchConfigurationName"),
//   		},
//   		LaunchTemplate: &LaunchTemplate{
//   			LaunchTemplateName: []*string{
//   				jsii.String("launchTemplateName"),
//   			},
//   		},
//   		LifecycleHookSpecificationList: []interface{}{
//   			lifecycleHookSpecificationList,
//   		},
//   		LoadBalancerNames: []*string{
//   			jsii.String("loadBalancerNames"),
//   		},
//   		MaxSize: []*string{
//   			jsii.String("maxSize"),
//   		},
//   		Metrics: []*string{
//   			jsii.String("metrics"),
//   		},
//   		MetricValue: []*string{
//   			jsii.String("metricValue"),
//   		},
//   		MinSize: []*string{
//   			jsii.String("minSize"),
//   		},
//   		MixedInstancesPolicy: &MixedInstancesPolicy{
//   			InstancesDistribution: &InstancesDistribution{
//   				OnDemandAllocationStrategy: []*string{
//   					jsii.String("onDemandAllocationStrategy"),
//   				},
//   				OnDemandBaseCapacity: []*string{
//   					jsii.String("onDemandBaseCapacity"),
//   				},
//   				OnDemandPercentageAboveBaseCapacity: []*string{
//   					jsii.String("onDemandPercentageAboveBaseCapacity"),
//   				},
//   				SpotAllocationStrategy: []*string{
//   					jsii.String("spotAllocationStrategy"),
//   				},
//   				SpotInstancePools: []*string{
//   					jsii.String("spotInstancePools"),
//   				},
//   			},
//   			LaunchTemplate: &LaunchTemplate1{
//   				LaunchTemplateSpecification: &LaunchTemplateSpecification{
//   					LaunchTemplateName: []*string{
//   						jsii.String("launchTemplateName"),
//   					},
//   					Version: []*string{
//   						jsii.String("version"),
//   					},
//   				},
//   				Overrides: []interface{}{
//   					overrides,
//   				},
//   			},
//   		},
//   		NewInstancesProtectedFromScaleIn: []*string{
//   			jsii.String("newInstancesProtectedFromScaleIn"),
//   		},
//   		NotificationTypes: []*string{
//   			jsii.String("notificationTypes"),
//   		},
//   		PolicyName: []*string{
//   			jsii.String("policyName"),
//   		},
//   		PolicyType: []*string{
//   			jsii.String("policyType"),
//   		},
//   		ProtectedFromScaleIn: []*string{
//   			jsii.String("protectedFromScaleIn"),
//   		},
//   		ScalingAdjustment: []*string{
//   			jsii.String("scalingAdjustment"),
//   		},
//   		ScheduledActionName: []*string{
//   			jsii.String("scheduledActionName"),
//   		},
//   		ScheduledActionNames: []*string{
//   			jsii.String("scheduledActionNames"),
//   		},
//   		ScheduledUpdateGroupActions: []interface{}{
//   			scheduledUpdateGroupActions,
//   		},
//   		SecurityGroups: []*string{
//   			jsii.String("securityGroups"),
//   		},
//   		ServiceLinkedRoleArn: []*string{
//   			jsii.String("serviceLinkedRoleArn"),
//   		},
//   		SpotPrice: []*string{
//   			jsii.String("spotPrice"),
//   		},
//   		StartTime: []*string{
//   			jsii.String("startTime"),
//   		},
//   		StepAdjustments: []RequestParametersItem1{
//   			&RequestParametersItem1{
//   				MetricIntervalLowerBound: []*string{
//   					jsii.String("metricIntervalLowerBound"),
//   				},
//   				ScalingAdjustment: []*string{
//   					jsii.String("scalingAdjustment"),
//   				},
//   			},
//   		},
//   		Tags: []RequestParametersItem{
//   			&RequestParametersItem{
//   				Key: []*string{
//   					jsii.String("key"),
//   				},
//   				PropagateAtLaunch: []*string{
//   					jsii.String("propagateAtLaunch"),
//   				},
//   				ResourceId: []*string{
//   					jsii.String("resourceId"),
//   				},
//   				ResourceType: []*string{
//   					jsii.String("resourceType"),
//   				},
//   				Value: []*string{
//   					jsii.String("value"),
//   				},
//   			},
//   		},
//   		TargetGroupArns: []*string{
//   			jsii.String("targetGroupArns"),
//   		},
//   		TargetTrackingConfiguration: &TargetTrackingConfiguration{
//   			CustomizedMetricSpecification: &CustomizedMetricSpecification{
//   				Dimensions: []CustomizedMetricSpecificationItem{
//   					&CustomizedMetricSpecificationItem{
//   						Name: []*string{
//   							jsii.String("name"),
//   						},
//   						Value: []*string{
//   							jsii.String("value"),
//   						},
//   					},
//   				},
//   				MetricName: []*string{
//   					jsii.String("metricName"),
//   				},
//   				Namespace: []*string{
//   					jsii.String("namespace"),
//   				},
//   				Statistic: []*string{
//   					jsii.String("statistic"),
//   				},
//   				Unit: []*string{
//   					jsii.String("unit"),
//   				},
//   			},
//   			PredefinedMetricSpecification: &PredefinedMetricSpecification{
//   				PredefinedMetricType: []*string{
//   					jsii.String("predefinedMetricType"),
//   				},
//   			},
//   			TargetValue: []*string{
//   				jsii.String("targetValue"),
//   			},
//   		},
//   		Time: []*string{
//   			jsii.String("time"),
//   		},
//   		TopicArn: []*string{
//   			jsii.String("topicArn"),
//   		},
//   		UserData: []*string{
//   			jsii.String("userData"),
//   		},
//   		VPcZoneIdentifier: []*string{
//   			jsii.String("vPcZoneIdentifier"),
//   		},
//   	},
//   	ResponseElements: &ResponseElements{
//   		Alarms: []interface{}{
//   			alarms,
//   		},
//   		FailedScheduledActions: []interface{}{
//   			failedScheduledActions,
//   		},
//   		FailedScheduledUpdateGroupActions: []interface{}{
//   			failedScheduledUpdateGroupActions,
//   		},
//   		PolicyArn: []*string{
//   			jsii.String("policyArn"),
//   		},
//   	},
//   	SourceIpAddress: []*string{
//   		jsii.String("sourceIpAddress"),
//   	},
//   	UserAgent: []*string{
//   		jsii.String("userAgent"),
//   	},
//   	UserIdentity: &UserIdentity{
//   		AccessKeyId: []*string{
//   			jsii.String("accessKeyId"),
//   		},
//   		AccountId: []*string{
//   			jsii.String("accountId"),
//   		},
//   		Arn: []*string{
//   			jsii.String("arn"),
//   		},
//   		PrincipalId: []*string{
//   			jsii.String("principalId"),
//   		},
//   		Type: []*string{
//   			jsii.String("type"),
//   		},
//   	},
//   }
//
// Experimental.
type AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps struct {
	// awsRegion property.
	//
	// Specify an array of string values to match this event if the actual value of awsRegion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	AwsRegion *[]*string `field:"optional" json:"awsRegion" yaml:"awsRegion"`
	// errorCode property.
	//
	// Specify an array of string values to match this event if the actual value of errorCode is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorCode *[]*string `field:"optional" json:"errorCode" yaml:"errorCode"`
	// errorMessage property.
	//
	// Specify an array of string values to match this event if the actual value of errorMessage is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ErrorMessage *[]*string `field:"optional" json:"errorMessage" yaml:"errorMessage"`
	// eventID property.
	//
	// Specify an array of string values to match this event if the actual value of eventID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventId *[]*string `field:"optional" json:"eventId" yaml:"eventId"`
	// EventBridge event metadata.
	// Default: - -.
	//
	// Experimental.
	EventMetadata *awscdk.AWSEventMetadataProps `field:"optional" json:"eventMetadata" yaml:"eventMetadata"`
	// eventName property.
	//
	// Specify an array of string values to match this event if the actual value of eventName is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventName *[]*string `field:"optional" json:"eventName" yaml:"eventName"`
	// eventSource property.
	//
	// Specify an array of string values to match this event if the actual value of eventSource is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventSource *[]*string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// eventTime property.
	//
	// Specify an array of string values to match this event if the actual value of eventTime is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventTime *[]*string `field:"optional" json:"eventTime" yaml:"eventTime"`
	// eventType property.
	//
	// Specify an array of string values to match this event if the actual value of eventType is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventType *[]*string `field:"optional" json:"eventType" yaml:"eventType"`
	// eventVersion property.
	//
	// Specify an array of string values to match this event if the actual value of eventVersion is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	EventVersion *[]*string `field:"optional" json:"eventVersion" yaml:"eventVersion"`
	// requestID property.
	//
	// Specify an array of string values to match this event if the actual value of requestID is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestId *[]*string `field:"optional" json:"requestId" yaml:"requestId"`
	// requestParameters property.
	//
	// Specify an array of string values to match this event if the actual value of requestParameters is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	RequestParameters *AWSAPICallViaCloudTrail_RequestParameters `field:"optional" json:"requestParameters" yaml:"requestParameters"`
	// responseElements property.
	//
	// Specify an array of string values to match this event if the actual value of responseElements is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	ResponseElements *AWSAPICallViaCloudTrail_ResponseElements `field:"optional" json:"responseElements" yaml:"responseElements"`
	// sourceIPAddress property.
	//
	// Specify an array of string values to match this event if the actual value of sourceIPAddress is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	SourceIpAddress *[]*string `field:"optional" json:"sourceIpAddress" yaml:"sourceIpAddress"`
	// userAgent property.
	//
	// Specify an array of string values to match this event if the actual value of userAgent is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserAgent *[]*string `field:"optional" json:"userAgent" yaml:"userAgent"`
	// userIdentity property.
	//
	// Specify an array of string values to match this event if the actual value of userIdentity is one of the values in the array. Use one of the constructors on the `aws_events.Match`  for more advanced matching options.
	// Default: - Do not filter on this field.
	//
	// Experimental.
	UserIdentity *AWSAPICallViaCloudTrail_UserIdentity `field:"optional" json:"userIdentity" yaml:"userIdentity"`
}


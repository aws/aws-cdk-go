package previewawsautoscalingplansmixins


// `ScalingInstruction` is a property of [ScalingPlan](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html) that specifies the scaling instruction for a scalable resource in a scaling plan. Each scaling instruction applies to one resource.
//
// AWS Auto Scaling creates target tracking scaling policies based on the scaling instructions. Target tracking scaling policies adjust the capacity of your scalable resource as required to maintain resource utilization at the target value that you specified.
//
// AWS Auto Scaling also configures predictive scaling for your Amazon EC2 Auto Scaling groups using a subset of properties, including the load metric, the scaling metric, the target value for the scaling metric, the predictive scaling mode (forecast and scale or forecast only), and the desired behavior when the forecast capacity exceeds the maximum capacity of the resource. With predictive scaling, AWS Auto Scaling generates forecasts with traffic predictions for the two days ahead and schedules scaling actions that proactively add and remove resource capacity to match the forecast.
//
// > We recommend waiting a minimum of 24 hours after creating an Auto Scaling group to configure predictive scaling. At minimum, there must be 24 hours of historical data to generate a forecast. For more information, see [Best practices for scaling plans](https://docs.aws.amazon.com/autoscaling/plans/userguide/gs-best-practices.html) in the *Scaling Plans User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scalingInstructionProperty := &ScalingInstructionProperty{
//   	CustomizedLoadMetricSpecification: &CustomizedLoadMetricSpecificationProperty{
//   		Dimensions: []interface{}{
//   			&MetricDimensionProperty{
//   				Name: jsii.String("name"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		MetricName: jsii.String("metricName"),
//   		Namespace: jsii.String("namespace"),
//   		Statistic: jsii.String("statistic"),
//   		Unit: jsii.String("unit"),
//   	},
//   	DisableDynamicScaling: jsii.Boolean(false),
//   	MaxCapacity: jsii.Number(123),
//   	MinCapacity: jsii.Number(123),
//   	PredefinedLoadMetricSpecification: &PredefinedLoadMetricSpecificationProperty{
//   		PredefinedLoadMetricType: jsii.String("predefinedLoadMetricType"),
//   		ResourceLabel: jsii.String("resourceLabel"),
//   	},
//   	PredictiveScalingMaxCapacityBehavior: jsii.String("predictiveScalingMaxCapacityBehavior"),
//   	PredictiveScalingMaxCapacityBuffer: jsii.Number(123),
//   	PredictiveScalingMode: jsii.String("predictiveScalingMode"),
//   	ResourceId: jsii.String("resourceId"),
//   	ScalableDimension: jsii.String("scalableDimension"),
//   	ScalingPolicyUpdateBehavior: jsii.String("scalingPolicyUpdateBehavior"),
//   	ScheduledActionBufferTime: jsii.Number(123),
//   	ServiceNamespace: jsii.String("serviceNamespace"),
//   	TargetTrackingConfigurations: []interface{}{
//   		&TargetTrackingConfigurationProperty{
//   			CustomizedScalingMetricSpecification: &CustomizedScalingMetricSpecificationProperty{
//   				Dimensions: []interface{}{
//   					&MetricDimensionProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				MetricName: jsii.String("metricName"),
//   				Namespace: jsii.String("namespace"),
//   				Statistic: jsii.String("statistic"),
//   				Unit: jsii.String("unit"),
//   			},
//   			DisableScaleIn: jsii.Boolean(false),
//   			EstimatedInstanceWarmup: jsii.Number(123),
//   			PredefinedScalingMetricSpecification: &PredefinedScalingMetricSpecificationProperty{
//   				PredefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			ScaleInCooldown: jsii.Number(123),
//   			ScaleOutCooldown: jsii.Number(123),
//   			TargetValue: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html
//
type CfnScalingPlanPropsMixin_ScalingInstructionProperty struct {
	// The customized load metric to use for predictive scaling.
	//
	// This property or a *PredefinedLoadMetricSpecification* is required when configuring predictive scaling, and cannot be used otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-customizedloadmetricspecification
	//
	CustomizedLoadMetricSpecification interface{} `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// Controls whether dynamic scaling is disabled.
	//
	// When dynamic scaling is enabled, AWS Auto Scaling creates target tracking scaling policies based on the specified target tracking configurations.
	//
	// The default is enabled ( `false` ).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-disabledynamicscaling
	//
	DisableDynamicScaling interface{} `field:"optional" json:"disableDynamicScaling" yaml:"disableDynamicScaling"`
	// The maximum capacity of the resource.
	//
	// The exception to this upper limit is if you specify a non-default setting for *PredictiveScalingMaxCapacityBehavior* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-maxcapacity
	//
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity of the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-mincapacity
	//
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// The predefined load metric to use for predictive scaling.
	//
	// This property or a *CustomizedLoadMetricSpecification* is required when configuring predictive scaling, and cannot be used otherwise.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predefinedloadmetricspecification
	//
	PredefinedLoadMetricSpecification interface{} `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	//
	// The default value is `SetForecastCapacityToMaxCapacity` .
	//
	// The following are possible values:
	//
	// - `SetForecastCapacityToMaxCapacity` - AWS Auto Scaling cannot scale resource capacity higher than the maximum capacity. The maximum capacity is enforced as a hard limit.
	// - `SetMaxCapacityToForecastCapacity` - AWS Auto Scaling can scale resource capacity higher than the maximum capacity to equal but not exceed forecast capacity.
	// - `SetMaxCapacityAboveForecastCapacity` - AWS Auto Scaling can scale resource capacity higher than the maximum capacity by a specified buffer value. The intention is to give the target tracking scaling policy extra capacity if unexpected traffic occurs.
	//
	// Valid only when configuring predictive scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmaxcapacitybehavior
	//
	PredictiveScalingMaxCapacityBehavior *string `field:"optional" json:"predictiveScalingMaxCapacityBehavior" yaml:"predictiveScalingMaxCapacityBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer. With a 10 percent buffer, if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//
	// Valid only when configuring predictive scaling. Required if *PredictiveScalingMaxCapacityBehavior* is set to `SetMaxCapacityAboveForecastCapacity` , and cannot be used otherwise.
	//
	// The range is 1-100.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmaxcapacitybuffer
	//
	PredictiveScalingMaxCapacityBuffer *float64 `field:"optional" json:"predictiveScalingMaxCapacityBuffer" yaml:"predictiveScalingMaxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// The default value is `ForecastAndScale` . Otherwise, AWS Auto Scaling forecasts capacity but does not apply any scheduled scaling actions based on the capacity forecast.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-predictivescalingmode
	//
	PredictiveScalingMode *string `field:"optional" json:"predictiveScalingMode" yaml:"predictiveScalingMode"`
	// The ID of the resource. This string consists of the resource type and unique identifier.
	//
	// - Auto Scaling group - The resource type is `autoScalingGroup` and the unique identifier is the name of the Auto Scaling group. Example: `autoScalingGroup/my-asg` .
	// - ECS service - The resource type is `service` and the unique identifier is the cluster name and service name. Example: `service/default/sample-webapp` .
	// - Spot Fleet request - The resource type is `spot-fleet-request` and the unique identifier is the Spot Fleet request ID. Example: `spot-fleet-request/sfr-73fbd2ce-aa30-494c-8788-1cee4EXAMPLE` .
	// - DynamoDB table - The resource type is `table` and the unique identifier is the resource ID. Example: `table/my-table` .
	// - DynamoDB global secondary index - The resource type is `index` and the unique identifier is the resource ID. Example: `table/my-table/index/my-table-index` .
	// - Aurora DB cluster - The resource type is `cluster` and the unique identifier is the cluster name. Example: `cluster:my-db-cluster` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-resourceid
	//
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension associated with the resource.
	//
	// - `autoscaling:autoScalingGroup:DesiredCapacity` - The desired capacity of an Auto Scaling group.
	// - `ecs:service:DesiredCount` - The desired task count of an ECS service.
	// - `ec2:spot-fleet-request:TargetCapacity` - The target capacity of a Spot Fleet request.
	// - `dynamodb:table:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB table.
	// - `dynamodb:table:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB table.
	// - `dynamodb:index:ReadCapacityUnits` - The provisioned read capacity for a DynamoDB global secondary index.
	// - `dynamodb:index:WriteCapacityUnits` - The provisioned write capacity for a DynamoDB global secondary index.
	// - `rds:cluster:ReadReplicaCount` - The count of Aurora Replicas in an Aurora DB cluster. Available for Aurora MySQL-compatible edition and Aurora PostgreSQL-compatible edition.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scalabledimension
	//
	ScalableDimension *string `field:"optional" json:"scalableDimension" yaml:"scalableDimension"`
	// Controls whether a resource's externally created scaling policies are deleted and new target tracking scaling policies created.
	//
	// The default value is `KeepExternalPolicies` .
	//
	// Valid only when configuring dynamic scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scalingpolicyupdatebehavior
	//
	ScalingPolicyUpdateBehavior *string `field:"optional" json:"scalingPolicyUpdateBehavior" yaml:"scalingPolicyUpdateBehavior"`
	// The amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	//
	// For example, if the forecast says to add capacity at 10:00 AM, and the buffer time is 5 minutes, then the run time of the corresponding scheduled scaling action will be 9:55 AM. The intention is to give resources time to be provisioned. For example, it can take a few minutes to launch an EC2 instance. The actual amount of time required depends on several factors, such as the size of the instance and whether there are startup scripts to complete.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). The default is 300 seconds.
	//
	// Valid only when configuring predictive scaling.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-scheduledactionbuffertime
	//
	ScheduledActionBufferTime *float64 `field:"optional" json:"scheduledActionBufferTime" yaml:"scheduledActionBufferTime"`
	// The namespace of the AWS service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-servicenamespace
	//
	ServiceNamespace *string `field:"optional" json:"serviceNamespace" yaml:"serviceNamespace"`
	// The target tracking configurations (up to 10).
	//
	// Each of these structures must specify a unique scaling metric and a target value for the metric.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscalingplans-scalingplan-scalinginstruction.html#cfn-autoscalingplans-scalingplan-scalinginstruction-targettrackingconfigurations
	//
	TargetTrackingConfigurations interface{} `field:"optional" json:"targetTrackingConfigurations" yaml:"targetTrackingConfigurations"`
}


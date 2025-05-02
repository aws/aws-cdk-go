package awsautoscalingplans


// Properties for defining a `CfnScalingPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPlanProps := &CfnScalingPlanProps{
//   	ApplicationSource: &ApplicationSourceProperty{
//   		CloudFormationStackArn: jsii.String("cloudFormationStackArn"),
//   		TagFilters: []interface{}{
//   			&TagFilterProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	ScalingInstructions: []interface{}{
//   		&ScalingInstructionProperty{
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			ResourceId: jsii.String("resourceId"),
//   			ScalableDimension: jsii.String("scalableDimension"),
//   			ServiceNamespace: jsii.String("serviceNamespace"),
//   			TargetTrackingConfigurations: []interface{}{
//   				&TargetTrackingConfigurationProperty{
//   					TargetValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					CustomizedScalingMetricSpecification: &CustomizedScalingMetricSpecificationProperty{
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: jsii.String("statistic"),
//
//   						// the properties below are optional
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						Unit: jsii.String("unit"),
//   					},
//   					DisableScaleIn: jsii.Boolean(false),
//   					EstimatedInstanceWarmup: jsii.Number(123),
//   					PredefinedScalingMetricSpecification: &PredefinedScalingMetricSpecificationProperty{
//   						PredefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//
//   						// the properties below are optional
//   						ResourceLabel: jsii.String("resourceLabel"),
//   					},
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			CustomizedLoadMetricSpecification: &CustomizedLoadMetricSpecificationProperty{
//   				MetricName: jsii.String("metricName"),
//   				Namespace: jsii.String("namespace"),
//   				Statistic: jsii.String("statistic"),
//
//   				// the properties below are optional
//   				Dimensions: []interface{}{
//   					&MetricDimensionProperty{
//   						Name: jsii.String("name"),
//   						Value: jsii.String("value"),
//   					},
//   				},
//   				Unit: jsii.String("unit"),
//   			},
//   			DisableDynamicScaling: jsii.Boolean(false),
//   			PredefinedLoadMetricSpecification: &PredefinedLoadMetricSpecificationProperty{
//   				PredefinedLoadMetricType: jsii.String("predefinedLoadMetricType"),
//
//   				// the properties below are optional
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			PredictiveScalingMaxCapacityBehavior: jsii.String("predictiveScalingMaxCapacityBehavior"),
//   			PredictiveScalingMaxCapacityBuffer: jsii.Number(123),
//   			PredictiveScalingMode: jsii.String("predictiveScalingMode"),
//   			ScalingPolicyUpdateBehavior: jsii.String("scalingPolicyUpdateBehavior"),
//   			ScheduledActionBufferTime: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
//
type CfnScalingPlanProps struct {
	// A CloudFormation stack or a set of tags.
	//
	// You can create one scaling plan per application source. The `ApplicationSource` property must be present to ensure interoperability with the AWS Auto Scaling console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-applicationsource
	//
	ApplicationSource interface{} `field:"required" json:"applicationSource" yaml:"applicationSource"`
	// The scaling instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-scalinginstructions
	//
	ScalingInstructions interface{} `field:"required" json:"scalingInstructions" yaml:"scalingInstructions"`
}


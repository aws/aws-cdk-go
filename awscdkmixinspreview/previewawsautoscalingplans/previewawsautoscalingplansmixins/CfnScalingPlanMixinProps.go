package previewawsautoscalingplansmixins


// Properties for CfnScalingPlanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnScalingPlanMixinProps := &CfnScalingPlanMixinProps{
//   	ApplicationSource: &ApplicationSourceProperty{
//   		CloudFormationStackArn: jsii.String("cloudFormationStackArn"),
//   		TagFilters: []interface{}{
//   			&TagFilterProperty{
//   				Key: jsii.String("key"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	ScalingInstructions: []interface{}{
//   		&ScalingInstructionProperty{
//   			CustomizedLoadMetricSpecification: &CustomizedLoadMetricSpecificationProperty{
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
//   			DisableDynamicScaling: jsii.Boolean(false),
//   			MaxCapacity: jsii.Number(123),
//   			MinCapacity: jsii.Number(123),
//   			PredefinedLoadMetricSpecification: &PredefinedLoadMetricSpecificationProperty{
//   				PredefinedLoadMetricType: jsii.String("predefinedLoadMetricType"),
//   				ResourceLabel: jsii.String("resourceLabel"),
//   			},
//   			PredictiveScalingMaxCapacityBehavior: jsii.String("predictiveScalingMaxCapacityBehavior"),
//   			PredictiveScalingMaxCapacityBuffer: jsii.Number(123),
//   			PredictiveScalingMode: jsii.String("predictiveScalingMode"),
//   			ResourceId: jsii.String("resourceId"),
//   			ScalableDimension: jsii.String("scalableDimension"),
//   			ScalingPolicyUpdateBehavior: jsii.String("scalingPolicyUpdateBehavior"),
//   			ScheduledActionBufferTime: jsii.Number(123),
//   			ServiceNamespace: jsii.String("serviceNamespace"),
//   			TargetTrackingConfigurations: []interface{}{
//   				&TargetTrackingConfigurationProperty{
//   					CustomizedScalingMetricSpecification: &CustomizedScalingMetricSpecificationProperty{
//   						Dimensions: []interface{}{
//   							&MetricDimensionProperty{
//   								Name: jsii.String("name"),
//   								Value: jsii.String("value"),
//   							},
//   						},
//   						MetricName: jsii.String("metricName"),
//   						Namespace: jsii.String("namespace"),
//   						Statistic: jsii.String("statistic"),
//   						Unit: jsii.String("unit"),
//   					},
//   					DisableScaleIn: jsii.Boolean(false),
//   					EstimatedInstanceWarmup: jsii.Number(123),
//   					PredefinedScalingMetricSpecification: &PredefinedScalingMetricSpecificationProperty{
//   						PredefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//   						ResourceLabel: jsii.String("resourceLabel"),
//   					},
//   					ScaleInCooldown: jsii.Number(123),
//   					ScaleOutCooldown: jsii.Number(123),
//   					TargetValue: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html
//
type CfnScalingPlanMixinProps struct {
	// A CloudFormation stack or a set of tags.
	//
	// You can create one scaling plan per application source. The `ApplicationSource` property must be present to ensure interoperability with the AWS Auto Scaling console.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-applicationsource
	//
	ApplicationSource interface{} `field:"optional" json:"applicationSource" yaml:"applicationSource"`
	// The scaling instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscalingplans-scalingplan.html#cfn-autoscalingplans-scalingplan-scalinginstructions
	//
	ScalingInstructions interface{} `field:"optional" json:"scalingInstructions" yaml:"scalingInstructions"`
}


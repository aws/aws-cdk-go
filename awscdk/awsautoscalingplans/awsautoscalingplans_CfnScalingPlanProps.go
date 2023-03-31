package awsautoscalingplans


// Properties for defining a `CfnScalingPlan`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScalingPlanProps := &cfnScalingPlanProps{
//   	applicationSource: &applicationSourceProperty{
//   		cloudFormationStackArn: jsii.String("cloudFormationStackArn"),
//   		tagFilters: []interface{}{
//   			&tagFilterProperty{
//   				key: jsii.String("key"),
//
//   				// the properties below are optional
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   		},
//   	},
//   	scalingInstructions: []interface{}{
//   		&scalingInstructionProperty{
//   			maxCapacity: jsii.Number(123),
//   			minCapacity: jsii.Number(123),
//   			resourceId: jsii.String("resourceId"),
//   			scalableDimension: jsii.String("scalableDimension"),
//   			serviceNamespace: jsii.String("serviceNamespace"),
//   			targetTrackingConfigurations: []interface{}{
//   				&targetTrackingConfigurationProperty{
//   					targetValue: jsii.Number(123),
//
//   					// the properties below are optional
//   					customizedScalingMetricSpecification: &customizedScalingMetricSpecificationProperty{
//   						metricName: jsii.String("metricName"),
//   						namespace: jsii.String("namespace"),
//   						statistic: jsii.String("statistic"),
//
//   						// the properties below are optional
//   						dimensions: []interface{}{
//   							&metricDimensionProperty{
//   								name: jsii.String("name"),
//   								value: jsii.String("value"),
//   							},
//   						},
//   						unit: jsii.String("unit"),
//   					},
//   					disableScaleIn: jsii.Boolean(false),
//   					estimatedInstanceWarmup: jsii.Number(123),
//   					predefinedScalingMetricSpecification: &predefinedScalingMetricSpecificationProperty{
//   						predefinedScalingMetricType: jsii.String("predefinedScalingMetricType"),
//
//   						// the properties below are optional
//   						resourceLabel: jsii.String("resourceLabel"),
//   					},
//   					scaleInCooldown: jsii.Number(123),
//   					scaleOutCooldown: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			customizedLoadMetricSpecification: &customizedLoadMetricSpecificationProperty{
//   				metricName: jsii.String("metricName"),
//   				namespace: jsii.String("namespace"),
//   				statistic: jsii.String("statistic"),
//
//   				// the properties below are optional
//   				dimensions: []interface{}{
//   					&metricDimensionProperty{
//   						name: jsii.String("name"),
//   						value: jsii.String("value"),
//   					},
//   				},
//   				unit: jsii.String("unit"),
//   			},
//   			disableDynamicScaling: jsii.Boolean(false),
//   			predefinedLoadMetricSpecification: &predefinedLoadMetricSpecificationProperty{
//   				predefinedLoadMetricType: jsii.String("predefinedLoadMetricType"),
//
//   				// the properties below are optional
//   				resourceLabel: jsii.String("resourceLabel"),
//   			},
//   			predictiveScalingMaxCapacityBehavior: jsii.String("predictiveScalingMaxCapacityBehavior"),
//   			predictiveScalingMaxCapacityBuffer: jsii.Number(123),
//   			predictiveScalingMode: jsii.String("predictiveScalingMode"),
//   			scalingPolicyUpdateBehavior: jsii.String("scalingPolicyUpdateBehavior"),
//   			scheduledActionBufferTime: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnScalingPlanProps struct {
	// A CloudFormation stack or a set of tags.
	//
	// You can create one scaling plan per application source. The `ApplicationSource` property must be present to ensure interoperability with the AWS Auto Scaling console.
	ApplicationSource interface{} `field:"required" json:"applicationSource" yaml:"applicationSource"`
	// The scaling instructions.
	ScalingInstructions interface{} `field:"required" json:"scalingInstructions" yaml:"scalingInstructions"`
}


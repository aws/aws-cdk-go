package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &CfnEndpointProps{
//   	EndpointConfigName: jsii.String("endpointConfigName"),
//
//   	// the properties below are optional
//   	DeploymentConfig: &DeploymentConfigProperty{
//   		AutoRollbackConfiguration: &AutoRollbackConfigProperty{
//   			Alarms: []interface{}{
//   				&AlarmProperty{
//   					AlarmName: jsii.String("alarmName"),
//   				},
//   			},
//   		},
//   		BlueGreenUpdatePolicy: &BlueGreenUpdatePolicyProperty{
//   			TrafficRoutingConfiguration: &TrafficRoutingConfigProperty{
//   				Type: jsii.String("type"),
//
//   				// the properties below are optional
//   				CanarySize: &CapacitySizeProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   				LinearStepSize: &CapacitySizeProperty{
//   					Type: jsii.String("type"),
//   					Value: jsii.Number(123),
//   				},
//   				WaitIntervalInSeconds: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			TerminationWaitInSeconds: jsii.Number(123),
//   		},
//   		RollingUpdatePolicy: &RollingUpdatePolicyProperty{
//   			MaximumBatchSize: &CapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   			WaitIntervalInSeconds: jsii.Number(123),
//
//   			// the properties below are optional
//   			MaximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			RollbackMaximumBatchSize: &CapacitySizeProperty{
//   				Type: jsii.String("type"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   	EndpointName: jsii.String("endpointName"),
//   	ExcludeRetainedVariantProperties: []interface{}{
//   		&VariantPropertyProperty{
//   			VariantPropertyType: jsii.String("variantPropertyType"),
//   		},
//   	},
//   	RetainAllVariantProperties: jsii.Boolean(false),
//   	RetainDeploymentConfig: jsii.Boolean(false),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html
//
type CfnEndpointProps struct {
	// The name of the [AWS::SageMaker::EndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html) resource that specifies the configuration for the endpoint. For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointconfigname
	//
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-deploymentconfig
	//
	DeploymentConfig interface{} `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The name of the endpoint.The name must be unique within an AWS Region in your AWS account. The name is case-insensitive in `CreateEndpoint` , but the case is preserved and must be matched in [](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointname
	//
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// When you are updating endpoint resources with [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) whose value is set to `true` , `ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html) to override with the values provided by `EndpointConfig` . If you don't specify a value for `ExcludeAllVariantProperties` , no variant properties are overridden. Don't use this property when creating new endpoint resources or when `RetainAllVariantProperties` is set to `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-excluderetainedvariantproperties
	//
	ExcludeRetainedVariantProperties interface{} `field:"optional" json:"excludeRetainedVariantProperties" yaml:"excludeRetainedVariantProperties"`
	// When updating endpoint resources, enables or disables the retention of variant properties, such as the instance count or the variant weight.
	//
	// To retain the variant properties of an endpoint when updating it, set `RetainAllVariantProperties` to `true` . To use the variant properties specified in a new `EndpointConfig` call when updating an endpoint, set `RetainAllVariantProperties` to `false` . Use this property only when updating endpoint resources, not when creating new endpoint resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-retainallvariantproperties
	//
	RetainAllVariantProperties interface{} `field:"optional" json:"retainAllVariantProperties" yaml:"retainAllVariantProperties"`
	// Specifies whether to reuse the last deployment configuration.
	//
	// The default value is false (the configuration is not reused).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-retaindeploymentconfig
	//
	RetainDeploymentConfig interface{} `field:"optional" json:"retainDeploymentConfig" yaml:"retainDeploymentConfig"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


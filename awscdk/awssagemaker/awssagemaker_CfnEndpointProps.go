package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnEndpoint`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEndpointProps := &cfnEndpointProps{
//   	endpointConfigName: jsii.String("endpointConfigName"),
//
//   	// the properties below are optional
//   	deploymentConfig: &deploymentConfigProperty{
//   		blueGreenUpdatePolicy: &blueGreenUpdatePolicyProperty{
//   			trafficRoutingConfiguration: &trafficRoutingConfigProperty{
//   				type: jsii.String("type"),
//
//   				// the properties below are optional
//   				canarySize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				linearStepSize: &capacitySizeProperty{
//   					type: jsii.String("type"),
//   					value: jsii.Number(123),
//   				},
//   				waitIntervalInSeconds: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			maximumExecutionTimeoutInSeconds: jsii.Number(123),
//   			terminationWaitInSeconds: jsii.Number(123),
//   		},
//
//   		// the properties below are optional
//   		autoRollbackConfiguration: &autoRollbackConfigProperty{
//   			alarms: []interface{}{
//   				&alarmProperty{
//   					alarmName: jsii.String("alarmName"),
//   				},
//   			},
//   		},
//   	},
//   	endpointName: jsii.String("endpointName"),
//   	excludeRetainedVariantProperties: []interface{}{
//   		&variantPropertyProperty{
//   			variantPropertyType: jsii.String("variantPropertyType"),
//   		},
//   	},
//   	retainAllVariantProperties: jsii.Boolean(false),
//   	retainDeploymentConfig: jsii.Boolean(false),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEndpointProps struct {
	// The name of the [AWS::SageMaker::EndpointConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html) resource that specifies the configuration for the endpoint. For more information, see [CreateEndpointConfig](https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html) .
	EndpointConfigName *string `field:"required" json:"endpointConfigName" yaml:"endpointConfigName"`
	// The deployment configuration for an endpoint, which contains the desired deployment strategy and rollback configurations.
	DeploymentConfig interface{} `field:"optional" json:"deploymentConfig" yaml:"deploymentConfig"`
	// The name of the endpoint.The name must be unique within an AWS Region in your AWS account. The name is case-insensitive in `CreateEndpoint` , but the case is preserved and must be matched in [](https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html) .
	EndpointName *string `field:"optional" json:"endpointName" yaml:"endpointName"`
	// When you are updating endpoint resources with [RetainAllVariantProperties](https://docs.aws.amazon.com/sagemaker/latest/dg/API_UpdateEndpoint.html#SageMaker-UpdateEndpoint-request-RetainAllVariantProperties) whose value is set to `true` , `ExcludeRetainedVariantProperties` specifies the list of type [VariantProperty](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-endpoint-variantproperty.html) to override with the values provided by `EndpointConfig` . If you don't specify a value for `ExcludeAllVariantProperties` , no variant properties are overridden. Don't use this property when creating new endpoint resources or when `RetainAllVariantProperties` is set to `false` .
	ExcludeRetainedVariantProperties interface{} `field:"optional" json:"excludeRetainedVariantProperties" yaml:"excludeRetainedVariantProperties"`
	// When updating endpoint resources, enables or disables the retention of variant properties, such as the instance count or the variant weight.
	//
	// To retain the variant properties of an endpoint when updating it, set `RetainAllVariantProperties` to `true` . To use the variant properties specified in a new `EndpointConfig` call when updating an endpoint, set `RetainAllVariantProperties` to `false` . Use this property only when updating endpoint resources, not when creating new endpoint resources.
	RetainAllVariantProperties interface{} `field:"optional" json:"retainAllVariantProperties" yaml:"retainAllVariantProperties"`
	// Specifies whether to reuse the last deployment configuration.
	//
	// The default value is false (the configuration is not reused).
	RetainDeploymentConfig interface{} `field:"optional" json:"retainDeploymentConfig" yaml:"retainDeploymentConfig"`
	// A list of key-value pairs to apply to this resource.
	//
	// For more information, see [Resource Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html) and [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what) in the *AWS Billing and Cost Management User Guide* .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapacityProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnCapacityProviderMixinProps := &CfnCapacityProviderMixinProps{
//   	CapacityProviderName: jsii.String("capacityProviderName"),
//   	CapacityProviderScalingConfig: &CapacityProviderScalingConfigProperty{
//   		MaxVCpuCount: jsii.Number(123),
//   		ScalingMode: jsii.String("scalingMode"),
//   		ScalingPolicies: []interface{}{
//   			&TargetTrackingScalingPolicyProperty{
//   				PredefinedMetricType: jsii.String("predefinedMetricType"),
//   				TargetValue: jsii.Number(123),
//   			},
//   		},
//   	},
//   	InstanceRequirements: &InstanceRequirementsProperty{
//   		AllowedInstanceTypes: []*string{
//   			jsii.String("allowedInstanceTypes"),
//   		},
//   		Architectures: []*string{
//   			jsii.String("architectures"),
//   		},
//   		ExcludedInstanceTypes: []*string{
//   			jsii.String("excludedInstanceTypes"),
//   		},
//   	},
//   	KmsKeyArn: jsii.String("kmsKeyArn"),
//   	PermissionsConfig: &CapacityProviderPermissionsConfigProperty{
//   		CapacityProviderOperatorRoleArn: jsii.String("capacityProviderOperatorRoleArn"),
//   	},
//   	PropagateTags: &PropagateTagsConfigProperty{
//   		ExplicitTags: []interface{}{
//   			&CfnTag{
//   				Key: jsii.String("key"),
//   				Value: jsii.String("value"),
//   			},
//   		},
//   		Mode: jsii.String("mode"),
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TelemetryConfig: &CapacityProviderTelemetryConfigProperty{
//   		LoggingConfig: &CapacityProviderLoggingConfigProperty{
//   			LogGroup: jsii.String("logGroup"),
//   			SystemLogLevel: jsii.String("systemLogLevel"),
//   		},
//   	},
//   	VpcConfig: &CapacityProviderVpcConfigProperty{
//   		SecurityGroupIds: []*string{
//   			jsii.String("securityGroupIds"),
//   		},
//   		SubnetIds: []*string{
//   			jsii.String("subnetIds"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html
//
type CfnCapacityProviderMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-capacityprovidername
	//
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// The scaling configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-capacityproviderscalingconfig
	//
	CapacityProviderScalingConfig interface{} `field:"optional" json:"capacityProviderScalingConfig" yaml:"capacityProviderScalingConfig"`
	// The instance requirements for compute resources managed by the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The ARN of the KMS key used to encrypt the capacity provider's resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The permissions configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-permissionsconfig
	//
	PermissionsConfig interface{} `field:"optional" json:"permissionsConfig" yaml:"permissionsConfig"`
	// Configuration that defines how tags are propagated to managed resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-propagatetags
	//
	PropagateTags interface{} `field:"optional" json:"propagateTags" yaml:"propagateTags"`
	// A key-value pair that provides metadata for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-telemetryconfig
	//
	TelemetryConfig interface{} `field:"optional" json:"telemetryConfig" yaml:"telemetryConfig"`
	// The VPC configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


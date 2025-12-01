package previewawslambdamixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCapacityProviderPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// The name of the capacity provider.
	//
	// The name must be unique within your AWS account and region. If you don't specify a name, CloudFormation generates one.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-capacityprovidername
	//
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// The scaling configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-capacityproviderscalingconfig
	//
	CapacityProviderScalingConfig interface{} `field:"optional" json:"capacityProviderScalingConfig" yaml:"capacityProviderScalingConfig"`
	// Specifications for the types of EC2 instances that the capacity provider can use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-instancerequirements
	//
	InstanceRequirements interface{} `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The ARN of the AWS Key Management Service (KMS) key used by the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-kmskeyarn
	//
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// IAM permissions configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-permissionsconfig
	//
	PermissionsConfig interface{} `field:"optional" json:"permissionsConfig" yaml:"permissionsConfig"`
	// A list of tags to apply to the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// VPC configuration for the capacity provider.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lambda-capacityprovider.html#cfn-lambda-capacityprovider-vpcconfig
	//
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}


package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for creating a Lambda capacity provider.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("MyVpc"))
//   securityGroup := ec2.NewSecurityGroup(this, jsii.String("SecurityGroup"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   capacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProvider"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	ScalingOptions: lambda.ScalingOptions_Manual([]TargetTrackingScalingPolicy{
//   		lambda.TargetTrackingScalingPolicy_CpuUtilization(jsii.Number(70)),
//   	}),
//   })
//
type CapacityProviderProps struct {
	// A list of security group IDs to associate with EC2 instances launched by the capacity provider.
	//
	// Up to 5 security groups can be specified.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// A list of subnets where the capacity provider can launch EC2 instances.
	//
	// At least one subnet must be specified, and up to 16 subnets are supported.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// The instruction set architecture required for compute instances.
	//
	// Only one architecture can be specified per capacity provider.
	// Default: - No architecture constraints specified.
	//
	Architectures *[]Architecture `field:"optional" json:"architectures" yaml:"architectures"`
	// The name of the capacity provider.
	//
	// The name must be unique within the AWS account and region.
	// Default: - AWS CloudFormation generates a unique physical ID and uses that
	// ID for the capacity provider's name.
	//
	CapacityProviderName *string `field:"optional" json:"capacityProviderName" yaml:"capacityProviderName"`
	// Configuration for filtering instance types that the capacity provider can use.
	// Default: - No instance type filtering applied.
	//
	InstanceTypeFilter InstanceTypeFilter `field:"optional" json:"instanceTypeFilter" yaml:"instanceTypeFilter"`
	// The AWS Key Management Service (KMS) key used to encrypt data associated with the capacity provider.
	// Default: - No KMS key specified, uses an AWS-managed key instead.
	//
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// The maximum number of vCPUs that the capacity provider can scale up to.
	// Default: - No maximum limit specified, service default is 400.
	//
	MaxVCpuCount *float64 `field:"optional" json:"maxVCpuCount" yaml:"maxVCpuCount"`
	// The IAM role that the Lambda service assumes to manage the capacity provider.
	// Default: - A role will be generated containing the AWSLambdaManagedEC2ResourceOperator managed policy.
	//
	OperatorRole awsiam.IRole `field:"optional" json:"operatorRole" yaml:"operatorRole"`
	// The options for scaling a capacity provider, including scaling policies.
	// Default: - The `Auto` option is applied by default.
	//
	ScalingOptions ScalingOptions `field:"optional" json:"scalingOptions" yaml:"scalingOptions"`
}


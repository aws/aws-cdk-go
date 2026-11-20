package awslambda


// Accelerator configuration for Lambda Managed Instance (LMI) functions.
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
//   gpuCapacityProvider := lambda.NewCapacityProvider(this, jsii.String("GpuCapacityProvider"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	Architectures: []Architecture{
//   		lambda.Architecture_X86_64(),
//   	},
//   	InstanceTypeFilter: lambda.InstanceTypeFilter_Allow([]InstanceType{
//   		ec2.InstanceType_Of(ec2.InstanceClass_G6, ec2.InstanceSize_XLARGE),
//   	}),
//   	ScalingOptions: lambda.ScalingOptions_Manual([]TargetTrackingScalingPolicy{
//   		lambda.TargetTrackingScalingPolicy_GpuUtilization(jsii.Number(60)),
//   	}),
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_FROM_IMAGE(),
//   	Handler: lambda.Handler_FROM_IMAGE(),
//   	Code: lambda.Code_FromAssetImage(path.join(__dirname, jsii.String("lambda-handler"))),
//   	MemorySize: jsii.Number(2048),
//   	AcceleratorConfig: &AcceleratorConfig{
//   		AcceleratorMemorySize: lambda.AcceleratorMemorySize_GIB_24(),
//   	},
//   })
//
//   gpuCapacityProvider.AddFunction(fn)
//
type AcceleratorConfig struct {
	// The minimum amount of memory of each accelerator device exposed to the execution environment.
	//
	// Lambda might allocate more, depending on the accelerator instance types available.
	AcceleratorMemorySize AcceleratorMemorySize `field:"required" json:"acceleratorMemorySize" yaml:"acceleratorMemorySize"`
}


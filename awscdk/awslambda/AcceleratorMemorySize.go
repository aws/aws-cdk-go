package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The minimum amount of memory, in gibibytes (GiB), of each accelerator device exposed to a Lambda Managed Instance (LMI) function's execution environment.
//
// Lambda might allocate more, depending on the accelerator instance types available.
//
// Use one of the predefined values (`AcceleratorMemorySize.GIB_24`, etc.). If the
// service adds a size that is not yet listed here, use `AcceleratorMemorySize.of()`
// to supply it directly.
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
type AcceleratorMemorySize interface {
	// The amount of accelerator memory in gibibytes (GiB).
	//
	// May be an unresolved token when supplied via `of()`.
	Gibibytes() *float64
}

// The jsii proxy struct for AcceleratorMemorySize
type jsiiProxy_AcceleratorMemorySize struct {
	_ byte // padding
}

func (j *jsiiProxy_AcceleratorMemorySize) Gibibytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gibibytes",
		&returns,
	)
	return returns
}


// An amount of accelerator memory expressed as a number of gibibytes (GiB).
//
// Prefer the predefined constants; use this escape hatch only for a value the
// service supports that is not yet represented as a constant. The specific set
// of valid sizes is a service-side constraint enforced at deploy time, so this
// method does not restrict the value to the predefined constants. It does reject
// clearly-invalid input (a non-positive or non-integer number of GiB). The value
// may be an unresolved token, in which case validation is skipped.
func AcceleratorMemorySize_Of(gibibytes *float64) AcceleratorMemorySize {
	_init_.Initialize()

	if err := validateAcceleratorMemorySize_OfParameters(gibibytes); err != nil {
		panic(err)
	}
	var returns AcceleratorMemorySize

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"of",
		[]interface{}{gibibytes},
		&returns,
	)

	return returns
}

func AcceleratorMemorySize_GIB_12() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_12",
		&returns,
	)
	return returns
}

func AcceleratorMemorySize_GIB_16() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_16",
		&returns,
	)
	return returns
}

func AcceleratorMemorySize_GIB_24() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_24",
		&returns,
	)
	return returns
}

func AcceleratorMemorySize_GIB_3() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_3",
		&returns,
	)
	return returns
}

func AcceleratorMemorySize_GIB_48() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_48",
		&returns,
	)
	return returns
}

func AcceleratorMemorySize_GIB_6() AcceleratorMemorySize {
	_init_.Initialize()
	var returns AcceleratorMemorySize
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.AcceleratorMemorySize",
		"GIB_6",
		&returns,
	)
	return returns
}


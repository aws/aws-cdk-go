package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration for filtering instance types that a capacity provider can use.
//
// Instances types can either be allowed or excluded, not both.
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
//   // Allow only specific instance families
//   allowCapacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProviderAllowed"), &CapacityProviderProps{
//   	Subnets: vpc.PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	InstanceTypeFilter: lambda.InstanceTypeFilter_Allow([]InstanceType{
//   		ec2.InstanceType_Of(ec2.InstanceClass_M5, ec2.InstanceSize_LARGE),
//   		ec2.InstanceType_*Of(ec2.InstanceClass_M5, ec2.InstanceSize_XLARGE),
//   	}),
//   })
//
//   // Or exclude specific instance types
//   excludeCapacityProvider := lambda.NewCapacityProvider(this, jsii.String("MyCapacityProviderExcluded"), &CapacityProviderProps{
//   	Subnets: vpc.*PrivateSubnets,
//   	SecurityGroups: []ISecurityGroup{
//   		securityGroup,
//   	},
//   	InstanceTypeFilter: lambda.InstanceTypeFilter_Exclude([]InstanceType{
//   		ec2.InstanceType_*Of(ec2.InstanceClass_T2, ec2.InstanceSize_MICRO),
//   	}),
//   })
//
type InstanceTypeFilter interface {
	// A list of instance types that the capacity provider is allowed to use.
	AllowedInstanceTypes() *[]awsec2.InstanceType
	// A list of instance types that the capacity provider should not use.
	ExcludedInstanceTypes() *[]awsec2.InstanceType
}

// The jsii proxy struct for InstanceTypeFilter
type jsiiProxy_InstanceTypeFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_InstanceTypeFilter) AllowedInstanceTypes() *[]awsec2.InstanceType {
	var returns *[]awsec2.InstanceType
	_jsii_.Get(
		j,
		"allowedInstanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_InstanceTypeFilter) ExcludedInstanceTypes() *[]awsec2.InstanceType {
	var returns *[]awsec2.InstanceType
	_jsii_.Get(
		j,
		"excludedInstanceTypes",
		&returns,
	)
	return returns
}


// Creates an instance type filter that allows only the specified instance types.
func InstanceTypeFilter_Allow(instanceTypes *[]awsec2.InstanceType) InstanceTypeFilter {
	_init_.Initialize()

	if err := validateInstanceTypeFilter_AllowParameters(instanceTypes); err != nil {
		panic(err)
	}
	var returns InstanceTypeFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InstanceTypeFilter",
		"allow",
		[]interface{}{instanceTypes},
		&returns,
	)

	return returns
}

// Creates an instance type filter that excludes the specified instance types.
func InstanceTypeFilter_Exclude(instanceTypes *[]awsec2.InstanceType) InstanceTypeFilter {
	_init_.Initialize()

	if err := validateInstanceTypeFilter_ExcludeParameters(instanceTypes); err != nil {
		panic(err)
	}
	var returns InstanceTypeFilter

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.InstanceTypeFilter",
		"exclude",
		[]interface{}{instanceTypes},
		&returns,
	)

	return returns
}


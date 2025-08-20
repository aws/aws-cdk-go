package awsstepfunctionstasks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// VPC configuration.
type IBedrockCreateModelCustomizationJobVpcConfig interface {
	// VPC configuration security groups.
	//
	// The minimum number of security groups is 1.
	// The maximum number of security groups is 5.
	SecurityGroups() *[]awsec2.ISecurityGroup
	// VPC configuration subnets.
	//
	// The minimum number of subnets is 1.
	// The maximum number of subnets is 16.
	Subnets() *[]awsec2.ISubnet
}

// The jsii proxy for IBedrockCreateModelCustomizationJobVpcConfig
type jsiiProxy_IBedrockCreateModelCustomizationJobVpcConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_IBedrockCreateModelCustomizationJobVpcConfig) SecurityGroups() *[]awsec2.ISecurityGroup {
	var returns *[]awsec2.ISecurityGroup
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IBedrockCreateModelCustomizationJobVpcConfig) Subnets() *[]awsec2.ISubnet {
	var returns *[]awsec2.ISubnet
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}


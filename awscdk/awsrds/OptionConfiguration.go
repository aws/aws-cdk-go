package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
)

// Configuration properties for an option.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//   var vpc vpc
//
//   optionConfiguration := &OptionConfiguration{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Port: jsii.Number(123),
//   	SecurityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   	Settings: map[string]*string{
//   		"settingsKey": jsii.String("settings"),
//   	},
//   	Version: jsii.String("version"),
//   	Vpc: vpc,
//   }
//
type OptionConfiguration struct {
	// The name of the option.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The port number that this option uses.
	//
	// If `port` is specified then `vpc`
	// must also be specified.
	// Default: - no port.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Optional list of security groups to use for this option, if `vpc` is specified.
	//
	// If no groups are provided, a default one will be created.
	// Default: - a default group will be created if `port` or `vpc` are specified.
	//
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The settings for the option.
	// Default: - no settings.
	//
	Settings *map[string]*string `field:"optional" json:"settings" yaml:"settings"`
	// The version for the option.
	// Default: - no version.
	//
	Version *string `field:"optional" json:"version" yaml:"version"`
	// The VPC where a security group should be created for this option.
	//
	// If `vpc`
	// is specified then `port` must also be specified.
	// Default: - no VPC.
	//
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}


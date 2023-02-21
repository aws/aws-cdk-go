package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOptionGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOptionGroupProps := &CfnOptionGroupProps{
//   	EngineName: jsii.String("engineName"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	OptionGroupDescription: jsii.String("optionGroupDescription"),
//
//   	// the properties below are optional
//   	OptionConfigurations: []interface{}{
//   		&OptionConfigurationProperty{
//   			OptionName: jsii.String("optionName"),
//
//   			// the properties below are optional
//   			DbSecurityGroupMemberships: []*string{
//   				jsii.String("dbSecurityGroupMemberships"),
//   			},
//   			OptionSettings: []interface{}{
//   				&OptionSettingProperty{
//   					Name: jsii.String("name"),
//   					Value: jsii.String("value"),
//   				},
//   			},
//   			OptionVersion: jsii.String("optionVersion"),
//   			Port: jsii.Number(123),
//   			VpcSecurityGroupMemberships: []*string{
//   				jsii.String("vpcSecurityGroupMemberships"),
//   			},
//   		},
//   	},
//   	OptionGroupName: jsii.String("optionGroupName"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnOptionGroupProps struct {
	// Specifies the name of the engine that this option group should be associated with.
	//
	// Valid Values:
	//
	// - `mariadb`
	// - `mysql`
	// - `oracle-ee`
	// - `oracle-ee-cdb`
	// - `oracle-se2`
	// - `oracle-se2-cdb`
	// - `postgres`
	// - `sqlserver-ee`
	// - `sqlserver-se`
	// - `sqlserver-ex`
	// - `sqlserver-web`.
	EngineName *string `field:"required" json:"engineName" yaml:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	MajorEngineVersion *string `field:"required" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// The description of the option group.
	OptionGroupDescription *string `field:"required" json:"optionGroupDescription" yaml:"optionGroupDescription"`
	// A list of options and the settings for each option.
	OptionConfigurations interface{} `field:"optional" json:"optionConfigurations" yaml:"optionConfigurations"`
	// The name of the option group to be created.
	//
	// Constraints:
	//
	// - Must be 1 to 255 letters, numbers, or hyphens
	// - First character must be a letter
	// - Can't end with a hyphen or contain two consecutive hyphens
	//
	// Example: `myoptiongroup`
	//
	// If you don't specify a value for `OptionGroupName` property, a name is automatically created for the option group.
	//
	// > This value is stored as a lowercase string.
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// An optional array of key-value pairs to apply to this option group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


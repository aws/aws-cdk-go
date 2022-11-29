package awsrds

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnOptionGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOptionGroupProps := &cfnOptionGroupProps{
//   	engineName: jsii.String("engineName"),
//   	majorEngineVersion: jsii.String("majorEngineVersion"),
//   	optionGroupDescription: jsii.String("optionGroupDescription"),
//
//   	// the properties below are optional
//   	optionConfigurations: []interface{}{
//   		&optionConfigurationProperty{
//   			optionName: jsii.String("optionName"),
//
//   			// the properties below are optional
//   			dbSecurityGroupMemberships: []*string{
//   				jsii.String("dbSecurityGroupMemberships"),
//   			},
//   			optionSettings: []interface{}{
//   				&optionSettingProperty{
//   					name: jsii.String("name"),
//   					value: jsii.String("value"),
//   				},
//   			},
//   			optionVersion: jsii.String("optionVersion"),
//   			port: jsii.Number(123),
//   			vpcSecurityGroupMemberships: []*string{
//   				jsii.String("vpcSecurityGroupMemberships"),
//   			},
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
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
	// - `oracle-se2`
	// - `oracle-se1`
	// - `oracle-se`
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
	// Tags to assign to the option group.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


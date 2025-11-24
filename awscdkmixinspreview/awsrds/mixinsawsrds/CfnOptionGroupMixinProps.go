package mixinsawsrds

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnOptionGroupPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnOptionGroupMixinProps := &CfnOptionGroupMixinProps{
//   	EngineName: jsii.String("engineName"),
//   	MajorEngineVersion: jsii.String("majorEngineVersion"),
//   	OptionConfigurations: []interface{}{
//   		&OptionConfigurationProperty{
//   			DbSecurityGroupMemberships: []*string{
//   				jsii.String("dbSecurityGroupMemberships"),
//   			},
//   			OptionName: jsii.String("optionName"),
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
//   	OptionGroupDescription: jsii.String("optionGroupDescription"),
//   	OptionGroupName: jsii.String("optionGroupName"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html
//
type CfnOptionGroupMixinProps struct {
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-enginename
	//
	EngineName *string `field:"optional" json:"engineName" yaml:"engineName"`
	// Specifies the major version of the engine that this option group should be associated with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-majorengineversion
	//
	MajorEngineVersion *string `field:"optional" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// A list of all available options for an option group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optionconfigurations
	//
	OptionConfigurations interface{} `field:"optional" json:"optionConfigurations" yaml:"optionConfigurations"`
	// The description of the option group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupdescription
	//
	OptionGroupDescription *string `field:"optional" json:"optionGroupDescription" yaml:"optionGroupDescription"`
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-optiongroupname
	//
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// Tags to assign to the option group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-optiongroup.html#cfn-rds-optiongroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


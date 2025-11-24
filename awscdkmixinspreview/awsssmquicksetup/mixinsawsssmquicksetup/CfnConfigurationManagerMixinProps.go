package mixinsawsssmquicksetup


// Properties for CfnConfigurationManagerPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnConfigurationManagerMixinProps := &CfnConfigurationManagerMixinProps{
//   	ConfigurationDefinitions: []interface{}{
//   		&ConfigurationDefinitionProperty{
//   			Id: jsii.String("id"),
//   			LocalDeploymentAdministrationRoleArn: jsii.String("localDeploymentAdministrationRoleArn"),
//   			LocalDeploymentExecutionRoleName: jsii.String("localDeploymentExecutionRoleName"),
//   			Parameters: map[string]*string{
//   				"parametersKey": jsii.String("parameters"),
//   			},
//   			Type: jsii.String("type"),
//   			TypeVersion: jsii.String("typeVersion"),
//   		},
//   	},
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html
//
type CfnConfigurationManagerMixinProps struct {
	// The definition of the Quick Setup configuration that the configuration manager deploys.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html#cfn-ssmquicksetup-configurationmanager-configurationdefinitions
	//
	ConfigurationDefinitions interface{} `field:"optional" json:"configurationDefinitions" yaml:"configurationDefinitions"`
	// The description of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html#cfn-ssmquicksetup-configurationmanager-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The name of the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html#cfn-ssmquicksetup-configurationmanager-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Key-value pairs of metadata to assign to the configuration manager.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssmquicksetup-configurationmanager.html#cfn-ssmquicksetup-configurationmanager-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


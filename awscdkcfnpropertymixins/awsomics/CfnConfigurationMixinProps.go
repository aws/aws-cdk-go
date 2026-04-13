package awsomics


// Properties for CfnConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnConfigurationMixinProps := &CfnConfigurationMixinProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   	RunConfigurations: &RunConfigurationsProperty{
//   		VpcConfig: &VpcConfigProperty{
//   			SecurityGroupIds: []*string{
//   				jsii.String("securityGroupIds"),
//   			},
//   			SubnetIds: []*string{
//   				jsii.String("subnetIds"),
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html
//
type CfnConfigurationMixinProps struct {
	// Optional description for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly name for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-runconfigurations
	//
	RunConfigurations interface{} `field:"optional" json:"runConfigurations" yaml:"runConfigurations"`
	// A map of resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


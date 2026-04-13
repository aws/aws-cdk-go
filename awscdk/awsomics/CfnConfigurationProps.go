package awsomics


// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &CfnConfigurationProps{
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
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html
//
type CfnConfigurationProps struct {
	// User-friendly name for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-runconfigurations
	//
	RunConfigurations interface{} `field:"required" json:"runConfigurations" yaml:"runConfigurations"`
	// Optional description for the configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A map of resource tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-omics-configuration.html#cfn-omics-configuration-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


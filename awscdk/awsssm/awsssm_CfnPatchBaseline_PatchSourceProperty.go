package awsssm


// `PatchSource` is the property type for the `Sources` resource of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.
//
// The AWS CloudFormation `AWS::SSM::PatchSource` resource is used to provide information about the patches to use to update target instances, including target operating systems and source repository. Applies to Linux instances only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchSourceProperty := &patchSourceProperty{
//   	configuration: jsii.String("configuration"),
//   	name: jsii.String("name"),
//   	products: []*string{
//   		jsii.String("products"),
//   	},
//   }
//
type CfnPatchBaseline_PatchSourceProperty struct {
	// The value of the yum repo configuration. For example:.
	//
	// `[main]`
	//
	// `name=MyCustomRepository`
	//
	// `baseurl=https://my-custom-repository`
	//
	// `enabled=1`
	//
	// > For information about other options available for your yum repository configuration, see [dnf.conf(5)](https://docs.aws.amazon.com/https://man7.org/linux/man-pages/man5/dnf.conf.5.html) .
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name specified to identify the patch source.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "AmazonLinux2016.09", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
}


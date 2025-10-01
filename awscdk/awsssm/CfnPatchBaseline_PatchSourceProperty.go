package awsssm


// `PatchSource` is the property type for the `Sources` resource of the [AWS::SSM::PatchBaseline](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-patchbaseline.html) resource.
//
// The AWS CloudFormation `AWS::SSM::PatchSource` resource is used to provide information about the patches to use to update target instances, including target operating systems and source repository. Applies to Linux managed nodes only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   patchSourceProperty := &PatchSourceProperty{
//   	Configuration: jsii.String("configuration"),
//   	Name: jsii.String("name"),
//   	Products: []*string{
//   		jsii.String("products"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchsource.html
//
type CfnPatchBaseline_PatchSourceProperty struct {
	// The value of the repo configuration.
	//
	// *Example for yum repositories*
	//
	// `[main]`
	//
	// `name=MyCustomRepository`
	//
	// `baseurl=https://my-custom-repository`
	//
	// `enabled=1`
	//
	// For information about other options available for your yum repository configuration, see [dnf.conf(5)](https://docs.aws.amazon.com/https://man7.org/linux/man-pages/man5/dnf.conf.5.html) on the *man7.org* website.
	//
	// *Examples for Ubuntu Server and Debian Server*
	//
	// `deb http://security.ubuntu.com/ubuntu jammy main`
	//
	// `deb https://site.example.com/debian distribution component1 component2 component3`
	//
	// Repo information for Ubuntu Server repositories must be specifed in a single line. For more examples and information, see [jammy (5) sources.list.5.gz](https://docs.aws.amazon.com/https://manpages.ubuntu.com/manpages/jammy/man5/sources.list.5.html) on the *Ubuntu Server Manuals* website and [sources.list format](https://docs.aws.amazon.com/https://wiki.debian.org/SourcesList#sources.list_format) on the *Debian Wiki* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchsource.html#cfn-ssm-patchbaseline-patchsource-configuration
	//
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// The name specified to identify the patch source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchsource.html#cfn-ssm-patchbaseline-patchsource-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The specific operating system versions a patch repository applies to, such as "Ubuntu16.04", "RedhatEnterpriseLinux7.2" or "Suse12.7". For lists of supported product values, see [PatchFilter](https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchFilter.html) in the *AWS Systems Manager API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-patchbaseline-patchsource.html#cfn-ssm-patchbaseline-patchsource-products
	//
	Products *[]*string `field:"optional" json:"products" yaml:"products"`
}


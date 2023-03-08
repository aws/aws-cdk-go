package awsopsworks


// Contains the information required to retrieve an app or cookbook from a repository.
//
// For more information, see [Creating Apps](https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-creating.html) or [Custom Recipes and Cookbooks](https://docs.aws.amazon.com/opsworks/latest/userguide/workingcookbook.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sourceProperty := &sourceProperty{
//   	password: jsii.String("password"),
//   	revision: jsii.String("revision"),
//   	sshKey: jsii.String("sshKey"),
//   	type: jsii.String("type"),
//   	url: jsii.String("url"),
//   	username: jsii.String("username"),
//   }
//
type CfnApp_SourceProperty struct {
	// When included in a request, the parameter depends on the repository type.
	//
	// - For Amazon S3 bundles, set `Password` to the appropriate IAM secret access key.
	// - For HTTP bundles and Subversion repositories, set `Password` to the password.
	//
	// For more information on how to safely handle IAM credentials, see [](https://docs.aws.amazon.com/general/latest/gr/aws-access-keys-best-practices.html) .
	//
	// In responses, AWS OpsWorks Stacks returns `*****FILTERED*****` instead of the actual value.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The application's version.
	//
	// AWS OpsWorks Stacks enables you to easily deploy new versions of an application. One of the simplest approaches is to have branches or revisions in your repository that represent different versions that can potentially be deployed.
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// In requests, the repository's SSH key.
	//
	// In responses, AWS OpsWorks Stacks returns `*****FILTERED*****` instead of the actual value.
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// The repository type.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The source URL.
	//
	// The following is an example of an Amazon S3 source URL: `https://s3.amazonaws.com/opsworks-demo-bucket/opsworks_cookbook_demo.tar.gz` .
	Url *string `field:"optional" json:"url" yaml:"url"`
	// This parameter depends on the repository type.
	//
	// - For Amazon S3 bundles, set `Username` to the appropriate IAM access key ID.
	// - For HTTP bundles, Git repositories, and Subversion repositories, set `Username` to the user name.
	Username *string `field:"optional" json:"username" yaml:"username"`
}


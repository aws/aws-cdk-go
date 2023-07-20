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
//   sourceProperty := &SourceProperty{
//   	Password: jsii.String("password"),
//   	Revision: jsii.String("revision"),
//   	SshKey: jsii.String("sshKey"),
//   	Type: jsii.String("type"),
//   	Url: jsii.String("url"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html
//
type CfnStack_SourceProperty struct {
	// When included in a request, the parameter depends on the repository type.
	//
	// - For Amazon S3 bundles, set `Password` to the appropriate IAM secret access key.
	// - For HTTP bundles and Subversion repositories, set `Password` to the password.
	//
	// For more information on how to safely handle IAM credentials, see [](https://docs.aws.amazon.com/general/latest/gr/aws-access-keys-best-practices.html) .
	//
	// In responses, AWS OpsWorks Stacks returns `*****FILTERED*****` instead of the actual value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-password
	//
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The application's version.
	//
	// AWS OpsWorks Stacks enables you to easily deploy new versions of an application. One of the simplest approaches is to have branches or revisions in your repository that represent different versions that can potentially be deployed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-revision
	//
	Revision *string `field:"optional" json:"revision" yaml:"revision"`
	// The repository's SSH key.
	//
	// For more information, see [Using Git Repository SSH Keys](https://docs.aws.amazon.com/opsworks/latest/userguide/workingapps-deploykeys.html) in the *AWS OpsWorks User Guide* . To pass in an SSH key as a parameter, see the following example:
	//
	// `"Parameters" : { "GitSSHKey" : { "Description" : "Change SSH key newlines to commas.", "Type" : "CommaDelimitedList", "NoEcho" : "true" }, ... "CustomCookbooksSource": { "Revision" : { "Ref": "GitRevision"}, "SshKey" : { "Fn::Join" : [ "\n", { "Ref": "GitSSHKey"} ] }, "Type": "git", "Url": { "Ref": "GitURL"} } ...`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-sshkey
	//
	SshKey *string `field:"optional" json:"sshKey" yaml:"sshKey"`
	// The repository type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The source URL.
	//
	// The following is an example of an Amazon S3 source URL: `https://s3.amazonaws.com/opsworks-demo-bucket/opsworks_cookbook_demo.tar.gz` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
	// This parameter depends on the repository type.
	//
	// - For Amazon S3 bundles, set `Username` to the appropriate IAM access key ID.
	// - For HTTP bundles, Git repositories, and Subversion repositories, set `Username` to the user name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-source.html#cfn-opsworks-stack-source-username
	//
	Username *string `field:"optional" json:"username" yaml:"username"`
}


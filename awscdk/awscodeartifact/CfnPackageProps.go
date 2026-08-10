package awscodeartifact


// Properties for defining a `CfnPackage`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPackageProps := &CfnPackageProps{
//   	DomainName: jsii.String("domainName"),
//   	Format: jsii.String("format"),
//   	Name: jsii.String("name"),
//   	Repository: jsii.String("repository"),
//
//   	// the properties below are optional
//   	Namespace: jsii.String("namespace"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html
//
type CfnPackageProps struct {
	// The name of the domain that contains the repository that contains the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The format of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-format
	//
	Format *string `field:"required" json:"format" yaml:"format"`
	// The name of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the repository that contains the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-repository
	//
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The namespace of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}


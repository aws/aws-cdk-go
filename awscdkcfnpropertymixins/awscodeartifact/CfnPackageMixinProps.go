package awscodeartifact


// Properties for CfnPackagePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnPackageMixinProps := &CfnPackageMixinProps{
//   	DomainName: jsii.String("domainName"),
//   	Format: jsii.String("format"),
//   	Name: jsii.String("name"),
//   	Namespace: jsii.String("namespace"),
//   	Repository: jsii.String("repository"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html
//
type CfnPackageMixinProps struct {
	// The name of the domain that contains the repository that contains the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-domainname
	//
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// The format of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-format
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The name of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The namespace of the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-namespace
	//
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The name of the repository that contains the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codeartifact-package.html#cfn-codeartifact-package-repository
	//
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}


package awscodeartifact


// The origin restrictions for the package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   restrictionsProperty := &RestrictionsProperty{
//   	Publish: jsii.String("publish"),
//   	Upstream: jsii.String("upstream"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-package-restrictions.html
//
type CfnPackage_RestrictionsProperty struct {
	// The package origin configuration that determines if new versions of the package can be published directly to the repository.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-package-restrictions.html#cfn-codeartifact-package-restrictions-publish
	//
	Publish *string `field:"optional" json:"publish" yaml:"publish"`
	// The package origin configuration that determines if new versions of the package can be added to the repository from an external connection or upstream source.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-package-restrictions.html#cfn-codeartifact-package-restrictions-upstream
	//
	Upstream *string `field:"optional" json:"upstream" yaml:"upstream"`
}


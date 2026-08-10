package awscodeartifact


// The package origin configuration for the package.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   originConfigurationProperty := &OriginConfigurationProperty{
//   	Restrictions: &RestrictionsProperty{
//   		Publish: jsii.String("publish"),
//   		Upstream: jsii.String("upstream"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-package-originconfiguration.html
//
type CfnPackage_OriginConfigurationProperty struct {
	// The origin restrictions for the package.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codeartifact-package-originconfiguration.html#cfn-codeartifact-package-originconfiguration-restrictions
	//
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
}


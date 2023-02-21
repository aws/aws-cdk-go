package awsopsworks


// Describes the Chef configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chefConfigurationProperty := &chefConfigurationProperty{
//   	berkshelfVersion: jsii.String("berkshelfVersion"),
//   	manageBerkshelf: jsii.Boolean(false),
//   }
//
type CfnStack_ChefConfigurationProperty struct {
	// The Berkshelf version.
	BerkshelfVersion *string `field:"optional" json:"berkshelfVersion" yaml:"berkshelfVersion"`
	// Whether to enable Berkshelf.
	ManageBerkshelf interface{} `field:"optional" json:"manageBerkshelf" yaml:"manageBerkshelf"`
}


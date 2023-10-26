package awsopsworks


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chefConfigurationProperty := &ChefConfigurationProperty{
//   	BerkshelfVersion: jsii.String("berkshelfVersion"),
//   	ManageBerkshelf: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
//
type CfnStack_ChefConfigurationProperty struct {
	// The Berkshelf version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-stack-chefconfiguration-berkshelfversion
	//
	BerkshelfVersion *string `field:"optional" json:"berkshelfVersion" yaml:"berkshelfVersion"`
	// Whether to enable Berkshelf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-stack-chefconfiguration-manageberkshelf
	//
	ManageBerkshelf interface{} `field:"optional" json:"manageBerkshelf" yaml:"manageBerkshelf"`
}


package awsimagebuilder


// Identifies an Amazon EC2 launch template to use for a specific account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   launchTemplateConfigurationProperty := &launchTemplateConfigurationProperty{
//   	accountId: jsii.String("accountId"),
//   	launchTemplateId: jsii.String("launchTemplateId"),
//   	setDefaultVersion: jsii.Boolean(false),
//   }
//
type CfnDistributionConfiguration_LaunchTemplateConfigurationProperty struct {
	// The account ID that this configuration applies to.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Identifies the Amazon EC2 launch template to use.
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Set the specified Amazon EC2 launch template as the default launch template for the specified account.
	SetDefaultVersion interface{} `field:"optional" json:"setDefaultVersion" yaml:"setDefaultVersion"`
}


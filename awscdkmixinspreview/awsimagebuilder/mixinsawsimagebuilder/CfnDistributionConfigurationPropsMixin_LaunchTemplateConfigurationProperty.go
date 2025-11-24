package mixinsawsimagebuilder


// Identifies an Amazon EC2 launch template to use for a specific account.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   launchTemplateConfigurationProperty := &LaunchTemplateConfigurationProperty{
//   	AccountId: jsii.String("accountId"),
//   	LaunchTemplateId: jsii.String("launchTemplateId"),
//   	SetDefaultVersion: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchtemplateconfiguration.html
//
type CfnDistributionConfigurationPropsMixin_LaunchTemplateConfigurationProperty struct {
	// The account ID that this configuration applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchtemplateconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchtemplateconfiguration-accountid
	//
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Identifies the Amazon EC2 launch template to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchtemplateconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchtemplateconfiguration-launchtemplateid
	//
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// Set the specified Amazon EC2 launch template as the default launch template for the specified account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-imagebuilder-distributionconfiguration-launchtemplateconfiguration.html#cfn-imagebuilder-distributionconfiguration-launchtemplateconfiguration-setdefaultversion
	//
	SetDefaultVersion interface{} `field:"optional" json:"setDefaultVersion" yaml:"setDefaultVersion"`
}


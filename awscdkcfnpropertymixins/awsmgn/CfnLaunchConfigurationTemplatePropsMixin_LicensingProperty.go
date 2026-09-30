package awsmgn


// Configuration of a machine's license.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   licensingProperty := &LicensingProperty{
//   	OsByol: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-licensing.html
//
type CfnLaunchConfigurationTemplatePropsMixin_LicensingProperty struct {
	// Whether to configure BYOL OS licensing.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mgn-launchconfigurationtemplate-licensing.html#cfn-mgn-launchconfigurationtemplate-licensing-osbyol
	//
	OsByol interface{} `field:"optional" json:"osByol" yaml:"osByol"`
}


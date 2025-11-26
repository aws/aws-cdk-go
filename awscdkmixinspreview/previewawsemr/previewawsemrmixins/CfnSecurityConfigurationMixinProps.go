package previewawsemrmixins


// Properties for CfnSecurityConfigurationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var securityConfiguration interface{}
//
//   cfnSecurityConfigurationMixinProps := &CfnSecurityConfigurationMixinProps{
//   	Name: jsii.String("name"),
//   	SecurityConfiguration: securityConfiguration,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
//
type CfnSecurityConfigurationMixinProps struct {
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The security configuration details in JSON format.
	//
	// For JSON parameters and examples, see [Use Security Configurations to Set Up Cluster Security](https://docs.aws.amazon.com/emr/latest/ManagementGuide/emr-security-configurations.html) in the *Amazon EMR Management Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
	//
	SecurityConfiguration interface{} `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
}


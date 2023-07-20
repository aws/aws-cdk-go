package awsemr


// Properties for defining a `CfnSecurityConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityConfiguration interface{}
//
//   cfnSecurityConfigurationProps := &CfnSecurityConfigurationProps{
//   	SecurityConfiguration: securityConfiguration,
//
//   	// the properties below are optional
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
//
type CfnSecurityConfigurationProps struct {
	// The security configuration details in JSON format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration
	//
	SecurityConfiguration interface{} `field:"required" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The name of the security configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


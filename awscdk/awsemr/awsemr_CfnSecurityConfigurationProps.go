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
//   cfnSecurityConfigurationProps := &cfnSecurityConfigurationProps{
//   	securityConfiguration: securityConfiguration,
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   }
//
type CfnSecurityConfigurationProps struct {
	// The security configuration details in JSON format.
	SecurityConfiguration interface{} `field:"required" json:"securityConfiguration" yaml:"securityConfiguration"`
	// The name of the security configuration.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


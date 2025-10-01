package awsglue


// A reference to a SecurityConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigurationReference := &SecurityConfigurationReference{
//   	SecurityConfigurationId: jsii.String("securityConfigurationId"),
//   }
//
type SecurityConfigurationReference struct {
	// The Id of the SecurityConfiguration resource.
	SecurityConfigurationId *string `field:"required" json:"securityConfigurationId" yaml:"securityConfigurationId"`
}


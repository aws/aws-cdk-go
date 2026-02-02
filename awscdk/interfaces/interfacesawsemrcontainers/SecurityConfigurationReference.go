package interfacesawsemrcontainers


// A reference to a SecurityConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigurationReference := &SecurityConfigurationReference{
//   	SecurityConfigurationArn: jsii.String("securityConfigurationArn"),
//   }
//
type SecurityConfigurationReference struct {
	// The Arn of the SecurityConfiguration resource.
	SecurityConfigurationArn *string `field:"required" json:"securityConfigurationArn" yaml:"securityConfigurationArn"`
}


package interfacesawsemr


// A reference to a SecurityConfiguration resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityConfigurationReference := &SecurityConfigurationReference{
//   	SecurityConfigurationName: jsii.String("securityConfigurationName"),
//   }
//
type SecurityConfigurationReference struct {
	// The Name of the SecurityConfiguration resource.
	SecurityConfigurationName *string `field:"required" json:"securityConfigurationName" yaml:"securityConfigurationName"`
}


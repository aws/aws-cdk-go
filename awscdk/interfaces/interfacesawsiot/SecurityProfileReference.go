package interfacesawsiot


// A reference to a SecurityProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityProfileReference := &SecurityProfileReference{
//   	SecurityProfileArn: jsii.String("securityProfileArn"),
//   	SecurityProfileName: jsii.String("securityProfileName"),
//   }
//
type SecurityProfileReference struct {
	// The ARN of the SecurityProfile resource.
	SecurityProfileArn *string `field:"required" json:"securityProfileArn" yaml:"securityProfileArn"`
	// The SecurityProfileName of the SecurityProfile resource.
	SecurityProfileName *string `field:"required" json:"securityProfileName" yaml:"securityProfileName"`
}


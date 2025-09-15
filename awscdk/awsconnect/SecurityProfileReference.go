package awsconnect


// A reference to a SecurityProfile resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityProfileReference := &SecurityProfileReference{
//   	SecurityProfileArn: jsii.String("securityProfileArn"),
//   }
//
type SecurityProfileReference struct {
	// The SecurityProfileArn of the SecurityProfile resource.
	SecurityProfileArn *string `field:"required" json:"securityProfileArn" yaml:"securityProfileArn"`
}


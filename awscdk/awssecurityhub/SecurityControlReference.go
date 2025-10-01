package awssecurityhub


// A reference to a SecurityControl resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   securityControlReference := &SecurityControlReference{
//   	SecurityControlId: jsii.String("securityControlId"),
//   }
//
type SecurityControlReference struct {
	// The SecurityControlId of the SecurityControl resource.
	SecurityControlId *string `field:"required" json:"securityControlId" yaml:"securityControlId"`
}


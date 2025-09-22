package awsshield


// A reference to a Protection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protectionReference := &ProtectionReference{
//   	ProtectionArn: jsii.String("protectionArn"),
//   }
//
type ProtectionReference struct {
	// The ProtectionArn of the Protection resource.
	ProtectionArn *string `field:"required" json:"protectionArn" yaml:"protectionArn"`
}


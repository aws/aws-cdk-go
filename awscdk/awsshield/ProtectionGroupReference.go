package awsshield


// A reference to a ProtectionGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   protectionGroupReference := &ProtectionGroupReference{
//   	ProtectionGroupArn: jsii.String("protectionGroupArn"),
//   }
//
type ProtectionGroupReference struct {
	// The ProtectionGroupArn of the ProtectionGroup resource.
	ProtectionGroupArn *string `field:"required" json:"protectionGroupArn" yaml:"protectionGroupArn"`
}


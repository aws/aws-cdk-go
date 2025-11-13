package interfacesawscloudfront


// A reference to a ConnectionGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionGroupReference := &ConnectionGroupReference{
//   	ConnectionGroupArn: jsii.String("connectionGroupArn"),
//   	ConnectionGroupId: jsii.String("connectionGroupId"),
//   }
//
type ConnectionGroupReference struct {
	// The ARN of the ConnectionGroup resource.
	ConnectionGroupArn *string `field:"required" json:"connectionGroupArn" yaml:"connectionGroupArn"`
	// The Id of the ConnectionGroup resource.
	ConnectionGroupId *string `field:"required" json:"connectionGroupId" yaml:"connectionGroupId"`
}


package interfacesawsmedialive


// A reference to a InputSecurityGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputSecurityGroupReference := &InputSecurityGroupReference{
//   	InputSecurityGroupArn: jsii.String("inputSecurityGroupArn"),
//   	InputSecurityGroupId: jsii.String("inputSecurityGroupId"),
//   }
//
type InputSecurityGroupReference struct {
	// The ARN of the InputSecurityGroup resource.
	InputSecurityGroupArn *string `field:"required" json:"inputSecurityGroupArn" yaml:"inputSecurityGroupArn"`
	// The Id of the InputSecurityGroup resource.
	InputSecurityGroupId *string `field:"required" json:"inputSecurityGroupId" yaml:"inputSecurityGroupId"`
}


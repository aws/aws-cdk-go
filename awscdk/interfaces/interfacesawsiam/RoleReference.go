package interfacesawsiam


// A reference to a Role resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   roleReference := &RoleReference{
//   	RoleArn: jsii.String("roleArn"),
//   	RoleName: jsii.String("roleName"),
//   }
//
type RoleReference struct {
	// The ARN of the Role resource.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The RoleName of the Role resource.
	RoleName *string `field:"required" json:"roleName" yaml:"roleName"`
}


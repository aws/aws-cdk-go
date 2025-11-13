package interfacesawsiotsitewise


// A reference to a AccessPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyReference := &AccessPolicyReference{
//   	AccessPolicyArn: jsii.String("accessPolicyArn"),
//   	AccessPolicyId: jsii.String("accessPolicyId"),
//   }
//
type AccessPolicyReference struct {
	// The ARN of the AccessPolicy resource.
	AccessPolicyArn *string `field:"required" json:"accessPolicyArn" yaml:"accessPolicyArn"`
	// The AccessPolicyId of the AccessPolicy resource.
	AccessPolicyId *string `field:"required" json:"accessPolicyId" yaml:"accessPolicyId"`
}


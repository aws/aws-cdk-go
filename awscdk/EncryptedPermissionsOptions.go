package awscdk


// Options for configuring permissions on encrypted resources.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   encryptedPermissionsOptions := &EncryptedPermissionsOptions{
//   	KeyActions: []*string{
//   		jsii.String("keyActions"),
//   	},
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   }
//
type EncryptedPermissionsOptions struct {
	// The ARNs of the resources to grant permissions on.
	// Default: - The ARN of the resource associated with the grant is used.
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
	// The KMS key actions to grant permissions for.
	// Default: - No permission is added to the KMS key, even if it exists.
	//
	KeyActions *[]*string `field:"optional" json:"keyActions" yaml:"keyActions"`
}


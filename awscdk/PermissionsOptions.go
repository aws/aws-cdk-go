package awscdk


// Options for configuring permissions in the `<Resource>.actions()` method.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   permissionsOptions := &PermissionsOptions{
//   	ResourceArns: []*string{
//   		jsii.String("resourceArns"),
//   	},
//   }
//
type PermissionsOptions struct {
	// The ARNs of the resources to grant permissions on.
	// Default: - The ARN of the resource associated with the grant is used.
	//
	ResourceArns *[]*string `field:"optional" json:"resourceArns" yaml:"resourceArns"`
}


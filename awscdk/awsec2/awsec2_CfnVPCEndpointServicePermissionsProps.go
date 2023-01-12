package awsec2


// Properties for defining a `CfnVPCEndpointServicePermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEndpointServicePermissionsProps := &cfnVPCEndpointServicePermissionsProps{
//   	serviceId: jsii.String("serviceId"),
//
//   	// the properties below are optional
//   	allowedPrincipals: []*string{
//   		jsii.String("allowedPrincipals"),
//   	},
//   }
//
type CfnVPCEndpointServicePermissionsProps struct {
	// The ID of the service.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// The Amazon Resource Names (ARN) of one or more principals (IAM users, IAM roles, and AWS accounts).
	//
	// Permissions are granted to the principals in this list. To grant permissions to all principals, specify an asterisk (*). Permissions are revoked for principals not in this list. If the list is empty, then all permissions are revoked.
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
}


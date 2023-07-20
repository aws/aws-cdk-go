package awsec2


// Properties for defining a `CfnVPCEndpointServicePermissions`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPCEndpointServicePermissionsProps := &CfnVPCEndpointServicePermissionsProps{
//   	ServiceId: jsii.String("serviceId"),
//
//   	// the properties below are optional
//   	AllowedPrincipals: []*string{
//   		jsii.String("allowedPrincipals"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html
//
type CfnVPCEndpointServicePermissionsProps struct {
	// The ID of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html#cfn-ec2-vpcendpointservicepermissions-serviceid
	//
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
	// The Amazon Resource Names (ARN) of one or more principals (for example, users, IAM roles, and AWS accounts ).
	//
	// Permissions are granted to the principals in this list. To grant permissions to all principals, specify an asterisk (*). Permissions are revoked for principals not in this list. If the list is empty, then all permissions are revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html#cfn-ec2-vpcendpointservicepermissions-allowedprincipals
	//
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
}


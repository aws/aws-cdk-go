package previewawsec2mixins


// Properties for CfnVPCEndpointServicePermissionsPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnVPCEndpointServicePermissionsMixinProps := &CfnVPCEndpointServicePermissionsMixinProps{
//   	AllowedPrincipals: []*string{
//   		jsii.String("allowedPrincipals"),
//   	},
//   	ServiceId: jsii.String("serviceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html
//
type CfnVPCEndpointServicePermissionsMixinProps struct {
	// The Amazon Resource Names (ARN) of one or more principals (for example, users, IAM roles, and AWS accounts ).
	//
	// Permissions are granted to the principals in this list. To grant permissions to all principals, specify an asterisk (*). Permissions are revoked for principals not in this list. If the list is empty, then all permissions are revoked.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html#cfn-ec2-vpcendpointservicepermissions-allowedprincipals
	//
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
	// The ID of the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpointservicepermissions.html#cfn-ec2-vpcendpointservicepermissions-serviceid
	//
	ServiceId *string `field:"optional" json:"serviceId" yaml:"serviceId"`
}


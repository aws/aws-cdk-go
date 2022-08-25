package awsec2


// Attributes when importing an existing client VPN endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   clientVpnEndpointAttributes := &clientVpnEndpointAttributes{
//   	endpointId: jsii.String("endpointId"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type ClientVpnEndpointAttributes struct {
	// The endpoint ID.
	EndpointId *string `field:"required" json:"endpointId" yaml:"endpointId"`
	// The security groups associated with the endpoint.
	SecurityGroups *[]ISecurityGroup `field:"required" json:"securityGroups" yaml:"securityGroups"`
}


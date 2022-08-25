package awsec2


// Construction properties for an ImportedInterfaceVpcEndpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var securityGroup securityGroup
//
//   interfaceVpcEndpointAttributes := &interfaceVpcEndpointAttributes{
//   	port: jsii.Number(123),
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//
//   	// the properties below are optional
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
type InterfaceVpcEndpointAttributes struct {
	// The port of the service of the interface VPC endpoint.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The interface VPC endpoint identifier.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The security groups associated with the interface VPC endpoint.
	//
	// If you wish to manage the network connections associated with this endpoint,
	// you will need to specify its security groups.
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


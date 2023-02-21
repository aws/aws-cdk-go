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
//   	securityGroupId: jsii.String("securityGroupId"),
//   	securityGroups: []iSecurityGroup{
//   		securityGroup,
//   	},
//   }
//
// Experimental.
type InterfaceVpcEndpointAttributes struct {
	// The port of the service of the interface VPC endpoint.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The interface VPC endpoint identifier.
	// Experimental.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The identifier of the security group associated with the interface VPC endpoint.
	// Deprecated: use `securityGroups` instead.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// The security groups associated with the interface VPC endpoint.
	// Experimental.
	SecurityGroups *[]ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


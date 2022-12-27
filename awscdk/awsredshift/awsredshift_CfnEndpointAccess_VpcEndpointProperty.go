package awsredshift


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEndpointProperty := &vpcEndpointProperty{
//   	networkInterfaces: []interface{}{
//   		&networkInterfaceProperty{
//   			availabilityZone: jsii.String("availabilityZone"),
//   			networkInterfaceId: jsii.String("networkInterfaceId"),
//   			privateIpAddress: jsii.String("privateIpAddress"),
//   			subnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	vpcEndpointId: jsii.String("vpcEndpointId"),
//   	vpcId: jsii.String("vpcId"),
//   }
//
type CfnEndpointAccess_VpcEndpointProperty struct {
	// `CfnEndpointAccess.VpcEndpointProperty.NetworkInterfaces`.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// `CfnEndpointAccess.VpcEndpointProperty.VpcEndpointId`.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// `CfnEndpointAccess.VpcEndpointProperty.VpcId`.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


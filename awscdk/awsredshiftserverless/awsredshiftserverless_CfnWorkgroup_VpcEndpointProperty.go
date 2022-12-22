package awsredshiftserverless


// The connection endpoint for connecting to Amazon Redshift Serverless through the proxy.
//
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
type CfnWorkgroup_VpcEndpointProperty struct {
	// One or more network interfaces of the endpoint.
	//
	// Also known as an interface endpoint.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The connection endpoint ID for connecting to Amazon Redshift Serverless.
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The VPC identifier that the endpoint is associated with.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


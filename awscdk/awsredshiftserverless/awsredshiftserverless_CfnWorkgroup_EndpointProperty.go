package awsredshiftserverless


// The VPC endpoint object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &endpointProperty{
//   	address: jsii.String("address"),
//   	port: jsii.Number(123),
//   	vpcEndpoints: []interface{}{
//   		&vpcEndpointProperty{
//   			networkInterfaces: []interface{}{
//   				&networkInterfaceProperty{
//   					availabilityZone: jsii.String("availabilityZone"),
//   					networkInterfaceId: jsii.String("networkInterfaceId"),
//   					privateIpAddress: jsii.String("privateIpAddress"),
//   					subnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			vpcEndpointId: jsii.String("vpcEndpointId"),
//   			vpcId: jsii.String("vpcId"),
//   		},
//   	},
//   }
//
type CfnWorkgroup_EndpointProperty struct {
	// The DNS address of the VPC endpoint.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that Amazon Redshift Serverless listens on.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An array of `VpcEndpoint` objects.
	VpcEndpoints interface{} `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}


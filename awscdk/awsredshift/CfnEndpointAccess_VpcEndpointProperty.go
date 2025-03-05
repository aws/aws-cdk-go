package awsredshift


// The connection endpoint for connecting to an Amazon Redshift cluster through the proxy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpcEndpointProperty := &VpcEndpointProperty{
//   	NetworkInterfaces: []interface{}{
//   		&NetworkInterfaceProperty{
//   			AvailabilityZone: jsii.String("availabilityZone"),
//   			NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   			PrivateIpAddress: jsii.String("privateIpAddress"),
//   			SubnetId: jsii.String("subnetId"),
//   		},
//   	},
//   	VpcEndpointId: jsii.String("vpcEndpointId"),
//   	VpcId: jsii.String("vpcId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html
//
type CfnEndpointAccess_VpcEndpointProperty struct {
	// One or more network interfaces of the endpoint.
	//
	// Also known as an interface endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-networkinterfaces
	//
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// The connection endpoint ID for connecting an Amazon Redshift cluster through the proxy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcendpointid
	//
	VpcEndpointId *string `field:"optional" json:"vpcEndpointId" yaml:"vpcEndpointId"`
	// The VPC identifier that the endpoint is associated.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshift-endpointaccess-vpcendpoint.html#cfn-redshift-endpointaccess-vpcendpoint-vpcid
	//
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


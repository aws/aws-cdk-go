package mixinsawsredshiftserverless


// The VPC endpoint object.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.Number(123),
//   	VpcEndpoints: []interface{}{
//   		&VpcEndpointProperty{
//   			NetworkInterfaces: []interface{}{
//   				&NetworkInterfaceProperty{
//   					AvailabilityZone: jsii.String("availabilityZone"),
//   					NetworkInterfaceId: jsii.String("networkInterfaceId"),
//   					PrivateIpAddress: jsii.String("privateIpAddress"),
//   					SubnetId: jsii.String("subnetId"),
//   				},
//   			},
//   			VpcEndpointId: jsii.String("vpcEndpointId"),
//   			VpcId: jsii.String("vpcId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html
//
type CfnWorkgroupPropsMixin_EndpointProperty struct {
	// The DNS address of the VPC endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port that Amazon Redshift Serverless listens on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-port
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// An array of `VpcEndpoint` objects.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-redshiftserverless-workgroup-endpoint.html#cfn-redshiftserverless-workgroup-endpoint-vpcendpoints
	//
	VpcEndpoints interface{} `field:"optional" json:"vpcEndpoints" yaml:"vpcEndpoints"`
}


package awsec2


// The DNS names of the endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConnectEndpointDnsNamesProperty := &InstanceConnectEndpointDnsNamesProperty{
//   	DnsName: jsii.String("dnsName"),
//   	FipsDnsName: jsii.String("fipsDnsName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointdnsnames.html
//
type CfnInstanceConnectEndpoint_InstanceConnectEndpointDnsNamesProperty struct {
	// The DNS name of the EC2 Instance Connect Endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointdnsnames.html#cfn-ec2-instanceconnectendpoint-instanceconnectendpointdnsnames-dnsname
	//
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// The Federal Information Processing Standards (FIPS) compliant DNS name of the EC2 Instance Connect Endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointdnsnames.html#cfn-ec2-instanceconnectendpoint-instanceconnectendpointdnsnames-fipsdnsname
	//
	FipsDnsName *string `field:"optional" json:"fipsDnsName" yaml:"fipsDnsName"`
}


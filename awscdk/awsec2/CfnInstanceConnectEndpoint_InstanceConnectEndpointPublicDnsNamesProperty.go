package awsec2


// The public DNS names of the endpoint, including IPv4-only and dualstack DNS names.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   instanceConnectEndpointPublicDnsNamesProperty := &InstanceConnectEndpointPublicDnsNamesProperty{
//   	Dualstack: &InstanceConnectEndpointDnsNamesProperty{
//   		DnsName: jsii.String("dnsName"),
//   		FipsDnsName: jsii.String("fipsDnsName"),
//   	},
//   	Ipv4: &InstanceConnectEndpointDnsNamesProperty{
//   		DnsName: jsii.String("dnsName"),
//   		FipsDnsName: jsii.String("fipsDnsName"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointpublicdnsnames.html
//
type CfnInstanceConnectEndpoint_InstanceConnectEndpointPublicDnsNamesProperty struct {
	// The DNS names of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointpublicdnsnames.html#cfn-ec2-instanceconnectendpoint-instanceconnectendpointpublicdnsnames-dualstack
	//
	Dualstack interface{} `field:"optional" json:"dualstack" yaml:"dualstack"`
	// The DNS names of the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instanceconnectendpoint-instanceconnectendpointpublicdnsnames.html#cfn-ec2-instanceconnectendpoint-instanceconnectendpointpublicdnsnames-ipv4
	//
	Ipv4 interface{} `field:"optional" json:"ipv4" yaml:"ipv4"`
}


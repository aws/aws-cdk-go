package awsdeadline


// The host property details.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   hostPropertiesRequestProperty := &HostPropertiesRequestProperty{
//   	HostName: jsii.String("hostName"),
//   	IpAddresses: &IpAddressesProperty{
//   		IpV4Addresses: []*string{
//   			jsii.String("ipV4Addresses"),
//   		},
//   		IpV6Addresses: []*string{
//   			jsii.String("ipV6Addresses"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-hostpropertiesrequest.html
//
type CfnWorkerPropsMixin_HostPropertiesRequestProperty struct {
	// The host name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-hostpropertiesrequest.html#cfn-deadline-worker-hostpropertiesrequest-hostname
	//
	HostName *string `field:"optional" json:"hostName" yaml:"hostName"`
	// The IP addresses for a host.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-worker-hostpropertiesrequest.html#cfn-deadline-worker-hostpropertiesrequest-ipaddresses
	//
	IpAddresses interface{} `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
}


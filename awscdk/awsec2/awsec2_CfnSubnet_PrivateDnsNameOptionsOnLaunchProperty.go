package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateDnsNameOptionsOnLaunchProperty := &privateDnsNameOptionsOnLaunchProperty{
//   	enableResourceNameDnsAaaaRecord: jsii.Boolean(false),
//   	enableResourceNameDnsARecord: jsii.Boolean(false),
//   	hostnameType: jsii.String("hostnameType"),
//   }
//
type CfnSubnet_PrivateDnsNameOptionsOnLaunchProperty struct {
	// `CfnSubnet.PrivateDnsNameOptionsOnLaunchProperty.EnableResourceNameDnsAAAARecord`.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// `CfnSubnet.PrivateDnsNameOptionsOnLaunchProperty.EnableResourceNameDnsARecord`.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// `CfnSubnet.PrivateDnsNameOptionsOnLaunchProperty.HostnameType`.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}


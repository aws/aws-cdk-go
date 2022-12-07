package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   serviceConnectClientAliasProperty := &serviceConnectClientAliasProperty{
//   	port: jsii.Number(123),
//
//   	// the properties below are optional
//   	dnsName: jsii.String("dnsName"),
//   }
//
type CfnService_ServiceConnectClientAliasProperty struct {
	// `CfnService.ServiceConnectClientAliasProperty.Port`.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// `CfnService.ServiceConnectClientAliasProperty.DnsName`.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
}


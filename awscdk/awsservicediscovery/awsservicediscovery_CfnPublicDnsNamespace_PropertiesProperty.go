package awsservicediscovery


// Properties for the public DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertiesProperty := &propertiesProperty{
//   	dnsProperties: &publicDnsPropertiesMutableProperty{
//   		soa: &sOAProperty{
//   			ttl: jsii.Number(123),
//   		},
//   	},
//   }
//
type CfnPublicDnsNamespace_PropertiesProperty struct {
	// DNS properties for the public DNS namespace.
	DnsProperties interface{} `field:"optional" json:"dnsProperties" yaml:"dnsProperties"`
}


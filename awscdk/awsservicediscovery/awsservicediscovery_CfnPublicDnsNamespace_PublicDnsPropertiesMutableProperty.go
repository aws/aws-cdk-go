package awsservicediscovery


// DNS properties for the public DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDnsPropertiesMutableProperty := &publicDnsPropertiesMutableProperty{
//   	soa: &sOAProperty{
//   		ttl: jsii.Number(123),
//   	},
//   }
//
type CfnPublicDnsNamespace_PublicDnsPropertiesMutableProperty struct {
	// Start of Authority (SOA) record for the hosted zone for the public DNS namespace.
	Soa interface{} `field:"optional" json:"soa" yaml:"soa"`
}


package awsservicediscovery


// DNS properties for the private DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   privateDnsPropertiesMutableProperty := &privateDnsPropertiesMutableProperty{
//   	soa: &sOAProperty{
//   		ttl: jsii.Number(123),
//   	},
//   }
//
type CfnPrivateDnsNamespace_PrivateDnsPropertiesMutableProperty struct {
	// Fields for the Start of Authority (SOA) record for the hosted zone for the private DNS namespace.
	Soa interface{} `field:"optional" json:"soa" yaml:"soa"`
}


package awsservicediscovery


// DNS properties for the public DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   publicDnsPropertiesMutableProperty := &PublicDnsPropertiesMutableProperty{
//   	Soa: &SOAProperty{
//   		Ttl: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-publicdnspropertiesmutable.html
//
type CfnPublicDnsNamespace_PublicDnsPropertiesMutableProperty struct {
	// Start of Authority (SOA) record for the hosted zone for the public DNS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-publicdnspropertiesmutable.html#cfn-servicediscovery-publicdnsnamespace-publicdnspropertiesmutable-soa
	//
	Soa interface{} `field:"optional" json:"soa" yaml:"soa"`
}


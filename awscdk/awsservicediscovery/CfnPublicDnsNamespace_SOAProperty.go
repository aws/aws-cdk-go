package awsservicediscovery


// Start of Authority (SOA) properties for a public or private DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sOAProperty := &SOAProperty{
//   	Ttl: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-soa.html
//
type CfnPublicDnsNamespace_SOAProperty struct {
	// The time to live (TTL) for purposes of negative caching.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-soa.html#cfn-servicediscovery-publicdnsnamespace-soa-ttl
	//
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}


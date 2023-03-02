package awsservicediscovery


// Start of Authority (SOA) properties for a public or private DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sOAProperty := &sOAProperty{
//   	ttl: jsii.Number(123),
//   }
//
type CfnPublicDnsNamespace_SOAProperty struct {
	// The time to live (TTL) for purposes of negative caching.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}


package interfacesawsroute53resolver


// A reference to a ResolverDNSSECConfig resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resolverDNSSECConfigReference := &ResolverDNSSECConfigReference{
//   	ResolverDnssecConfigId: jsii.String("resolverDnssecConfigId"),
//   }
//
type ResolverDNSSECConfigReference struct {
	// The Id of the ResolverDNSSECConfig resource.
	ResolverDnssecConfigId *string `field:"required" json:"resolverDnssecConfigId" yaml:"resolverDnssecConfigId"`
}


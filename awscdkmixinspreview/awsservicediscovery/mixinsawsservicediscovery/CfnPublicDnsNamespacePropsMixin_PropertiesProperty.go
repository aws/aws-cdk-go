package mixinsawsservicediscovery


// Properties for the public DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   propertiesProperty := &PropertiesProperty{
//   	DnsProperties: &PublicDnsPropertiesMutableProperty{
//   		Soa: &SOAProperty{
//   			Ttl: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-properties.html
//
type CfnPublicDnsNamespacePropsMixin_PropertiesProperty struct {
	// DNS properties for the public DNS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-publicdnsnamespace-properties.html#cfn-servicediscovery-publicdnsnamespace-properties-dnsproperties
	//
	DnsProperties interface{} `field:"optional" json:"dnsProperties" yaml:"dnsProperties"`
}


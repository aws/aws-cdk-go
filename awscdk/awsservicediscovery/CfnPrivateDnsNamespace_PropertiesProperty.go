package awsservicediscovery


// Properties for the private DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   propertiesProperty := &PropertiesProperty{
//   	DnsProperties: &PrivateDnsPropertiesMutableProperty{
//   		Soa: &SOAProperty{
//   			Ttl: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-properties.html
//
type CfnPrivateDnsNamespace_PropertiesProperty struct {
	// DNS properties for the private DNS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-properties.html#cfn-servicediscovery-privatednsnamespace-properties-dnsproperties
	//
	DnsProperties interface{} `field:"optional" json:"dnsProperties" yaml:"dnsProperties"`
}


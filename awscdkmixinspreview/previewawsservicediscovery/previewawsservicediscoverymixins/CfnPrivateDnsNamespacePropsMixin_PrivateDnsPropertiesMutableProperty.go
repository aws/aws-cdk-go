package previewawsservicediscoverymixins


// DNS properties for the private DNS namespace.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   privateDnsPropertiesMutableProperty := &PrivateDnsPropertiesMutableProperty{
//   	Soa: &SOAProperty{
//   		Ttl: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-privatednspropertiesmutable.html
//
type CfnPrivateDnsNamespacePropsMixin_PrivateDnsPropertiesMutableProperty struct {
	// Fields for the Start of Authority (SOA) record for the hosted zone for the private DNS namespace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-servicediscovery-privatednsnamespace-privatednspropertiesmutable.html#cfn-servicediscovery-privatednsnamespace-privatednspropertiesmutable-soa
	//
	Soa interface{} `field:"optional" json:"soa" yaml:"soa"`
}


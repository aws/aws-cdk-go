package awsmedialive


// Pair of multicast url and source ip address (optional) that make up a multicast source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   multicastSourceProperty := &MulticastSourceProperty{
//   	SourceIp: jsii.String("sourceIp"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html
//
type CfnInputPropsMixin_MulticastSourceProperty struct {
	// This represents the ip address of the device sending the multicast stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html#cfn-medialive-input-multicastsource-sourceip
	//
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// This represents the customer's source URL where multicast stream is pulled from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html#cfn-medialive-input-multicastsource-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


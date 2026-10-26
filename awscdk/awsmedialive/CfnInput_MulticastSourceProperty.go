package awsmedialive


// Pair of multicast url and source ip address (optional) that make up a multicast source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multicastSourceProperty := &MulticastSourceProperty{
//   	Url: jsii.String("url"),
//
//   	// the properties below are optional
//   	SourceIp: jsii.String("sourceIp"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html
//
type CfnInput_MulticastSourceProperty struct {
	// This represents the customer's source URL where multicast stream is pulled from.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html#cfn-medialive-input-multicastsource-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// This represents the ip address of the device sending the multicast stream.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-multicastsource.html#cfn-medialive-input-multicastsource-sourceip
	//
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
}


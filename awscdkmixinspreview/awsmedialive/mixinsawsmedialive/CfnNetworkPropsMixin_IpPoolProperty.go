package mixinsawsmedialive


// IP address cidr pool.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   ipPoolProperty := &IpPoolProperty{
//   	Cidr: jsii.String("cidr"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-network-ippool.html
//
type CfnNetworkPropsMixin_IpPoolProperty struct {
	// IP address cidr pool.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-network-ippool.html#cfn-medialive-network-ippool-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}


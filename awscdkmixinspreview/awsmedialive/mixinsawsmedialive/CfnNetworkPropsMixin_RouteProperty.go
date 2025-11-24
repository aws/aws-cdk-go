package mixinsawsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   routeProperty := &RouteProperty{
//   	Cidr: jsii.String("cidr"),
//   	Gateway: jsii.String("gateway"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-network-route.html
//
type CfnNetworkPropsMixin_RouteProperty struct {
	// Ip address cidr.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-network-route.html#cfn-medialive-network-route-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// IP address for the route packet paths.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-network-route.html#cfn-medialive-network-route-gateway
	//
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}


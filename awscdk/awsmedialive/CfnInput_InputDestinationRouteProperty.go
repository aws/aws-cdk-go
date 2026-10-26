package awsmedialive


// A network route configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputDestinationRouteProperty := &InputDestinationRouteProperty{
//   	Cidr: jsii.String("cidr"),
//   	Gateway: jsii.String("gateway"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationroute.html
//
type CfnInput_InputDestinationRouteProperty struct {
	// The CIDR of the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationroute.html#cfn-medialive-input-inputdestinationroute-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// An optional gateway for the route.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputdestinationroute.html#cfn-medialive-input-inputdestinationroute-gateway
	//
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}


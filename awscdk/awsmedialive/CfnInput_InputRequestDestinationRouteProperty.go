package awsmedialive


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputRequestDestinationRouteProperty := &InputRequestDestinationRouteProperty{
//   	Cidr: jsii.String("cidr"),
//   	Gateway: jsii.String("gateway"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputrequestdestinationroute.html
//
type CfnInput_InputRequestDestinationRouteProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputrequestdestinationroute.html#cfn-medialive-input-inputrequestdestinationroute-cidr
	//
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-input-inputrequestdestinationroute.html#cfn-medialive-input-inputrequestdestinationroute-gateway
	//
	Gateway *string `field:"optional" json:"gateway" yaml:"gateway"`
}


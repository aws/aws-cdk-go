package awsmediaconnect


// The configuration settings for a destination endpoint that represents a router input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerInputDestinationEndpointConfigurationProperty := &RouterInputDestinationEndpointConfigurationProperty{
//   	RouterInputArn: jsii.String("routerInputArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointconfiguration.html
//
type CfnDestinationEndpointPropsMixin_RouterInputDestinationEndpointConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the router input that the destination endpoint represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointconfiguration.html#cfn-mediaconnect-destinationendpoint-routerinputdestinationendpointconfiguration-routerinputarn
	//
	RouterInputArn *string `field:"optional" json:"routerInputArn" yaml:"routerInputArn"`
}


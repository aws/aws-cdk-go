package awsmediaconnect


// The details of the router input that is associated with a destination endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerInputDestinationEndpointAssociatedResourceDetailsProperty := &RouterInputDestinationEndpointAssociatedResourceDetailsProperty{
//   	MaximumBitrate: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails.html
//
type CfnDestinationEndpointPropsMixin_RouterInputDestinationEndpointAssociatedResourceDetailsProperty struct {
	// The maximum bitrate of the router input, in bits per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails.html#cfn-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails-maximumbitrate
	//
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails.html#cfn-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails.html#cfn-mediaconnect-destinationendpoint-routerinputdestinationendpointassociatedresourcedetails-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}


package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var standard interface{}
//
//   destinationEndpointAssociatedResourceDetailsProperty := &DestinationEndpointAssociatedResourceDetailsProperty{
//   	RouterInput: &RouterInputDestinationEndpointAssociatedResourceDetailsProperty{
//   		MaximumBitrate: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		State: jsii.String("state"),
//   	},
//   	Standard: standard,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointassociatedresourcedetails.html
//
type CfnDestinationEndpoint_DestinationEndpointAssociatedResourceDetailsProperty struct {
	// The details of the router input that is associated with a destination endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointassociatedresourcedetails.html#cfn-mediaconnect-destinationendpoint-destinationendpointassociatedresourcedetails-routerinput
	//
	RouterInput interface{} `field:"optional" json:"routerInput" yaml:"routerInput"`
	// The associated resource details for a standard destination endpoint.
	//
	// This type does not require any additional settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-destinationendpoint-destinationendpointassociatedresourcedetails.html#cfn-mediaconnect-destinationendpoint-destinationendpointassociatedresourcedetails-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}


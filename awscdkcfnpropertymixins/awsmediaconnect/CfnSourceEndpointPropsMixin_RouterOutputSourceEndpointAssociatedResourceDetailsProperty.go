package awsmediaconnect


// The details of the router output that is associated with a source endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   routerOutputSourceEndpointAssociatedResourceDetailsProperty := &RouterOutputSourceEndpointAssociatedResourceDetailsProperty{
//   	MaximumBitrate: jsii.Number(123),
//   	Name: jsii.String("name"),
//   	State: jsii.String("state"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails.html
//
type CfnSourceEndpointPropsMixin_RouterOutputSourceEndpointAssociatedResourceDetailsProperty struct {
	// The maximum bitrate of the router output, in bits per second.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails.html#cfn-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails-maximumbitrate
	//
	MaximumBitrate *float64 `field:"optional" json:"maximumBitrate" yaml:"maximumBitrate"`
	// The name of the router output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails.html#cfn-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails.html#cfn-mediaconnect-sourceendpoint-routeroutputsourceendpointassociatedresourcedetails-state
	//
	State *string `field:"optional" json:"state" yaml:"state"`
}


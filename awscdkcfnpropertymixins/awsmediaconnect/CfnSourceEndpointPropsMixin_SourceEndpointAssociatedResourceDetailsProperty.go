package awsmediaconnect


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var standard interface{}
//
//   sourceEndpointAssociatedResourceDetailsProperty := &SourceEndpointAssociatedResourceDetailsProperty{
//   	RouterOutput: &RouterOutputSourceEndpointAssociatedResourceDetailsProperty{
//   		MaximumBitrate: jsii.Number(123),
//   		Name: jsii.String("name"),
//   		State: jsii.String("state"),
//   	},
//   	Standard: standard,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointassociatedresourcedetails.html
//
type CfnSourceEndpointPropsMixin_SourceEndpointAssociatedResourceDetailsProperty struct {
	// The details of the router output that is associated with a source endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointassociatedresourcedetails.html#cfn-mediaconnect-sourceendpoint-sourceendpointassociatedresourcedetails-routeroutput
	//
	RouterOutput interface{} `field:"optional" json:"routerOutput" yaml:"routerOutput"`
	// The associated resource details for a standard source endpoint.
	//
	// This type does not require any additional settings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-sourceendpointassociatedresourcedetails.html#cfn-mediaconnect-sourceendpoint-sourceendpointassociatedresourcedetails-standard
	//
	Standard interface{} `field:"optional" json:"standard" yaml:"standard"`
}


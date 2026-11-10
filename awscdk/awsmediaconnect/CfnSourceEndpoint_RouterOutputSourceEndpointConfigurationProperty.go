package awsmediaconnect


// The configuration settings for a source endpoint that represents a router output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   routerOutputSourceEndpointConfigurationProperty := &RouterOutputSourceEndpointConfigurationProperty{
//   	RouterOutputArn: jsii.String("routerOutputArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointconfiguration.html
//
type CfnSourceEndpoint_RouterOutputSourceEndpointConfigurationProperty struct {
	// The Amazon Resource Name (ARN) of the router output that the source endpoint represents.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediaconnect-sourceendpoint-routeroutputsourceendpointconfiguration.html#cfn-mediaconnect-sourceendpoint-routeroutputsourceendpointconfiguration-routeroutputarn
	//
	RouterOutputArn *string `field:"required" json:"routerOutputArn" yaml:"routerOutputArn"`
}


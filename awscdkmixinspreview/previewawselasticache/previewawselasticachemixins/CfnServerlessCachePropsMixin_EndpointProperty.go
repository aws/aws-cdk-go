package previewawselasticachemixins


// Represents the information required for client programs to connect to a cache node.
//
// This value is read-only.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.String("port"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-endpoint.html
//
type CfnServerlessCachePropsMixin_EndpointProperty struct {
	// The DNS hostname of the cache node.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-endpoint.html#cfn-elasticache-serverlesscache-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The port number that the cache engine is listening on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-serverlesscache-endpoint.html#cfn-elasticache-serverlesscache-endpoint-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
}


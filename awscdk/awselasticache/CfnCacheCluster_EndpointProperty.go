package awselasticache


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   endpointProperty := &EndpointProperty{
//   	Address: jsii.String("address"),
//   	Port: jsii.String("port"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-endpoint.html
//
type CfnCacheCluster_EndpointProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-endpoint.html#cfn-elasticache-cachecluster-endpoint-address
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-cachecluster-endpoint.html#cfn-elasticache-cachecluster-endpoint-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
}


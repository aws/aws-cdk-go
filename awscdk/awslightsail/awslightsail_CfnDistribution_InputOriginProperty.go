package awslightsail


// `InputOrigin` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the origin resource of an Amazon Lightsail content delivery network (CDN) distribution.
//
// An origin can be a instance, bucket, or load balancer. A distribution pulls content from an origin, caches it, and serves it to viewers through a worldwide network of edge servers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputOriginProperty := &inputOriginProperty{
//   	name: jsii.String("name"),
//   	protocolPolicy: jsii.String("protocolPolicy"),
//   	regionName: jsii.String("regionName"),
//   }
//
type CfnDistribution_InputOriginProperty struct {
	// The name of the origin resource.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// The AWS Region name of the origin resource.
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}


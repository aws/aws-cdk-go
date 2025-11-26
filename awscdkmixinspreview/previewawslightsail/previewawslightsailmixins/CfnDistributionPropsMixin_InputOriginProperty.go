package previewawslightsailmixins


// `InputOrigin` is a property of the [AWS::Lightsail::Distribution](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-lightsail-distribution.html) resource. It describes the origin resource of an Amazon Lightsail content delivery network (CDN) distribution.
//
// An origin can be a instance, bucket, or load balancer. A distribution pulls content from an origin, caches it, and serves it to viewers through a worldwide network of edge servers.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   inputOriginProperty := &InputOriginProperty{
//   	Name: jsii.String("name"),
//   	ProtocolPolicy: jsii.String("protocolPolicy"),
//   	RegionName: jsii.String("regionName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-inputorigin.html
//
type CfnDistributionPropsMixin_InputOriginProperty struct {
	// The name of the origin resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-inputorigin.html#cfn-lightsail-distribution-inputorigin-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-inputorigin.html#cfn-lightsail-distribution-inputorigin-protocolpolicy
	//
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
	// The AWS Region name of the origin resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lightsail-distribution-inputorigin.html#cfn-lightsail-distribution-inputorigin-regionname
	//
	RegionName *string `field:"optional" json:"regionName" yaml:"regionName"`
}


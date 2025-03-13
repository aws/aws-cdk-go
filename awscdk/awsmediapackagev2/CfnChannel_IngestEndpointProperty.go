package awsmediapackagev2


// The input URL where the source stream should be sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestEndpointProperty := &IngestEndpointProperty{
//   	Id: jsii.String("id"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-ingestendpoint.html
//
type CfnChannel_IngestEndpointProperty struct {
	// The identifier associated with the ingest endpoint of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-ingestendpoint.html#cfn-mediapackagev2-channel-ingestendpoint-id
	//
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The URL associated with the ingest endpoint of the channel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-channel-ingestendpoint.html#cfn-mediapackagev2-channel-ingestendpoint-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


package awsmediapackage


// An endpoint for ingesting source content for a channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingestEndpointProperty := &IngestEndpointProperty{
//   	Id: jsii.String("id"),
//   	Password: jsii.String("password"),
//   	Url: jsii.String("url"),
//   	Username: jsii.String("username"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html
//
type CfnChannel_IngestEndpointProperty struct {
	// The endpoint identifier.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-id
	//
	Id *string `field:"required" json:"id" yaml:"id"`
	// The system-generated password for WebDAV input authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-password
	//
	Password *string `field:"required" json:"password" yaml:"password"`
	// The input URL where the source stream should be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-url
	//
	Url *string `field:"required" json:"url" yaml:"url"`
	// The system-generated username for WebDAV input authentication.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-ingestendpoint.html#cfn-mediapackage-channel-ingestendpoint-username
	//
	Username *string `field:"required" json:"username" yaml:"username"`
}


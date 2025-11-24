package mixinsawsmediapackage


// HLS ingest configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   hlsIngestProperty := &HlsIngestProperty{
//   	IngestEndpoints: []interface{}{
//   		&IngestEndpointProperty{
//   			Id: jsii.String("id"),
//   			Password: jsii.String("password"),
//   			Url: jsii.String("url"),
//   			Username: jsii.String("username"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-hlsingest.html
//
type CfnChannelPropsMixin_HlsIngestProperty struct {
	// The input URL where the source stream should be sent.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-channel-hlsingest.html#cfn-mediapackage-channel-hlsingest-ingestendpoints
	//
	IngestEndpoints interface{} `field:"optional" json:"ingestEndpoints" yaml:"ingestEndpoints"`
}


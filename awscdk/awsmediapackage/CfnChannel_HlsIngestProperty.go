package awsmediapackage


// HLS ingest configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnChannel_HlsIngestProperty struct {
	// The input URL where the source stream should be sent.
	IngestEndpoints interface{} `field:"optional" json:"ingestEndpoints" yaml:"ingestEndpoints"`
}


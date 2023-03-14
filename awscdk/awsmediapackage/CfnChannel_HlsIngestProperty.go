package awsmediapackage


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
	// `CfnChannel.HlsIngestProperty.ingestEndpoints`.
	IngestEndpoints interface{} `field:"optional" json:"ingestEndpoints" yaml:"ingestEndpoints"`
}


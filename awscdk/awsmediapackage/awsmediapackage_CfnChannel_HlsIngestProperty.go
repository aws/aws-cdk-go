package awsmediapackage


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   hlsIngestProperty := &hlsIngestProperty{
//   	ingestEndpoints: []interface{}{
//   		&ingestEndpointProperty{
//   			id: jsii.String("id"),
//   			password: jsii.String("password"),
//   			url: jsii.String("url"),
//   			username: jsii.String("username"),
//   		},
//   	},
//   }
//
type CfnChannel_HlsIngestProperty struct {
	// `CfnChannel.HlsIngestProperty.ingestEndpoints`.
	IngestEndpoints interface{} `field:"optional" json:"ingestEndpoints" yaml:"ingestEndpoints"`
}


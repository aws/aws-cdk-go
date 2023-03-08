package awsmedialive


// The settings for one RTMP output.
//
// The parent of this entity is OutputSettings.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   rtmpOutputSettingsProperty := &rtmpOutputSettingsProperty{
//   	certificateMode: jsii.String("certificateMode"),
//   	connectionRetryInterval: jsii.Number(123),
//   	destination: &outputLocationRefProperty{
//   		destinationRefId: jsii.String("destinationRefId"),
//   	},
//   	numRetries: jsii.Number(123),
//   }
//
type CfnChannel_RtmpOutputSettingsProperty struct {
	// If set to verifyAuthenticity, verifies the TLS certificate chain to a trusted certificate authority (CA).
	//
	// This causes RTMPS outputs with self-signed certificates to fail.
	CertificateMode *string `field:"optional" json:"certificateMode" yaml:"certificateMode"`
	// The number of seconds to wait before retrying a connection to the Flash Media server if the connection is lost.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// The RTMP endpoint excluding the stream name (for example, rtmp://host/appname).
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// The number of retry attempts.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}


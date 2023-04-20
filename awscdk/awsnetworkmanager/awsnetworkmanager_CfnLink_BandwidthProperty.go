package awsnetworkmanager


// Describes bandwidth information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   bandwidthProperty := &BandwidthProperty{
//   	DownloadSpeed: jsii.Number(123),
//   	UploadSpeed: jsii.Number(123),
//   }
//
type CfnLink_BandwidthProperty struct {
	// Download speed in Mbps.
	DownloadSpeed *float64 `field:"optional" json:"downloadSpeed" yaml:"downloadSpeed"`
	// Upload speed in Mbps.
	UploadSpeed *float64 `field:"optional" json:"uploadSpeed" yaml:"uploadSpeed"`
}


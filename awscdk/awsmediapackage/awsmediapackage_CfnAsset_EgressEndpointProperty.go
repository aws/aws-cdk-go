package awsmediapackage


// The playback endpoint for a packaging configuration on an asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressEndpointProperty := &egressEndpointProperty{
//   	packagingConfigurationId: jsii.String("packagingConfigurationId"),
//   	url: jsii.String("url"),
//   }
//
type CfnAsset_EgressEndpointProperty struct {
	// The ID of a packaging configuration that's applied to this asset.
	PackagingConfigurationId *string `field:"required" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The URL that's used to request content from this endpoint.
	Url *string `field:"required" json:"url" yaml:"url"`
}


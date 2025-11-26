package previewawsmediapackagemixins


// The playback endpoint for a packaging configuration on an asset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   egressEndpointProperty := &EgressEndpointProperty{
//   	PackagingConfigurationId: jsii.String("packagingConfigurationId"),
//   	Url: jsii.String("url"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html
//
type CfnAssetPropsMixin_EgressEndpointProperty struct {
	// The ID of a packaging configuration that's applied to this asset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html#cfn-mediapackage-asset-egressendpoint-packagingconfigurationid
	//
	PackagingConfigurationId *string `field:"optional" json:"packagingConfigurationId" yaml:"packagingConfigurationId"`
	// The URL that's used to request content from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-asset-egressendpoint.html#cfn-mediapackage-asset-egressendpoint-url
	//
	Url *string `field:"optional" json:"url" yaml:"url"`
}


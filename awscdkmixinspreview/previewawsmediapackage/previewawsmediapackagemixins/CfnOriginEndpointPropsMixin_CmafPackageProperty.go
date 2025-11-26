package previewawsmediapackagemixins


// Parameters for Common Media Application Format (CMAF) packaging.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cmafPackageProperty := &CmafPackageProperty{
//   	Encryption: &CmafEncryptionProperty{
//   		ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   		EncryptionMethod: jsii.String("encryptionMethod"),
//   		KeyRotationIntervalSeconds: jsii.Number(123),
//   		SpekeKeyProvider: &SpekeKeyProviderProperty{
//   			CertificateArn: jsii.String("certificateArn"),
//   			EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   				PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   				PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   			},
//   			ResourceId: jsii.String("resourceId"),
//   			RoleArn: jsii.String("roleArn"),
//   			SystemIds: []*string{
//   				jsii.String("systemIds"),
//   			},
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	HlsManifests: []interface{}{
//   		&HlsManifestProperty{
//   			AdMarkers: jsii.String("adMarkers"),
//   			AdsOnDeliveryRestrictions: jsii.String("adsOnDeliveryRestrictions"),
//   			AdTriggers: []*string{
//   				jsii.String("adTriggers"),
//   			},
//   			Id: jsii.String("id"),
//   			IncludeIframeOnlyStream: jsii.Boolean(false),
//   			ManifestName: jsii.String("manifestName"),
//   			PlaylistType: jsii.String("playlistType"),
//   			PlaylistWindowSeconds: jsii.Number(123),
//   			ProgramDateTimeIntervalSeconds: jsii.Number(123),
//   			Url: jsii.String("url"),
//   		},
//   	},
//   	SegmentDurationSeconds: jsii.Number(123),
//   	SegmentPrefix: jsii.String("segmentPrefix"),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html
//
type CfnOriginEndpointPropsMixin_CmafPackageProperty struct {
	// Parameters for encrypting content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-encryption
	//
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// A list of HLS manifest configurations that are available from this endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-hlsmanifests
	//
	HlsManifests interface{} `field:"optional" json:"hlsManifests" yaml:"hlsManifests"`
	// Duration (in seconds) of each segment.
	//
	// Actual segments are rounded to the nearest multiple of the source segment duration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-segmentdurationseconds
	//
	SegmentDurationSeconds *float64 `field:"optional" json:"segmentDurationSeconds" yaml:"segmentDurationSeconds"`
	// An optional custom string that is prepended to the name of each segment.
	//
	// If not specified, the segment prefix defaults to the ChannelId.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-segmentprefix
	//
	SegmentPrefix *string `field:"optional" json:"segmentPrefix" yaml:"segmentPrefix"`
	// Limitations for outputs from the endpoint, based on the video bitrate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-originendpoint-cmafpackage.html#cfn-mediapackage-originendpoint-cmafpackage-streamselection
	//
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}


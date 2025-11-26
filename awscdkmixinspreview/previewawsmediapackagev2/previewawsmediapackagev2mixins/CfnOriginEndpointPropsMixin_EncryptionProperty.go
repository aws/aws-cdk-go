package previewawsmediapackagev2mixins


// The parameters for encrypting content.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   encryptionProperty := &EncryptionProperty{
//   	CmafExcludeSegmentDrmMetadata: jsii.Boolean(false),
//   	ConstantInitializationVector: jsii.String("constantInitializationVector"),
//   	EncryptionMethod: &EncryptionMethodProperty{
//   		CmafEncryptionMethod: jsii.String("cmafEncryptionMethod"),
//   		IsmEncryptionMethod: jsii.String("ismEncryptionMethod"),
//   		TsEncryptionMethod: jsii.String("tsEncryptionMethod"),
//   	},
//   	KeyRotationIntervalSeconds: jsii.Number(123),
//   	SpekeKeyProvider: &SpekeKeyProviderProperty{
//   		DrmSystems: []*string{
//   			jsii.String("drmSystems"),
//   		},
//   		EncryptionContractConfiguration: &EncryptionContractConfigurationProperty{
//   			PresetSpeke20Audio: jsii.String("presetSpeke20Audio"),
//   			PresetSpeke20Video: jsii.String("presetSpeke20Video"),
//   		},
//   		ResourceId: jsii.String("resourceId"),
//   		RoleArn: jsii.String("roleArn"),
//   		Url: jsii.String("url"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html
//
type CfnOriginEndpointPropsMixin_EncryptionProperty struct {
	// Excludes SEIG and SGPD boxes from segment metadata in CMAF containers.
	//
	// When set to `true` , MediaPackage omits these DRM metadata boxes from CMAF segments, which can improve compatibility with certain devices and players that don't support these boxes.
	//
	// Important considerations:
	//
	// - This setting only affects CMAF container formats
	// - Key rotation can still be handled through media playlist signaling
	// - PSSH and TENC boxes remain unaffected
	// - Default behavior is preserved when this setting is disabled
	//
	// Valid values: `true` | `false`
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-cmafexcludesegmentdrmmetadata
	//
	CmafExcludeSegmentDrmMetadata interface{} `field:"optional" json:"cmafExcludeSegmentDrmMetadata" yaml:"cmafExcludeSegmentDrmMetadata"`
	// A 128-bit, 16-byte hex value represented by a 32-character string, used in conjunction with the key for encrypting content.
	//
	// If you don't specify a value, then MediaPackage creates the constant initialization vector (IV).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-constantinitializationvector
	//
	ConstantInitializationVector *string `field:"optional" json:"constantInitializationVector" yaml:"constantInitializationVector"`
	// The encryption method to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-encryptionmethod
	//
	EncryptionMethod interface{} `field:"optional" json:"encryptionMethod" yaml:"encryptionMethod"`
	// The interval, in seconds, to rotate encryption keys for the origin endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-keyrotationintervalseconds
	//
	KeyRotationIntervalSeconds *float64 `field:"optional" json:"keyRotationIntervalSeconds" yaml:"keyRotationIntervalSeconds"`
	// The SPEKE key provider to use for encryption.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-encryption.html#cfn-mediapackagev2-originendpoint-encryption-spekekeyprovider
	//
	SpekeKeyProvider interface{} `field:"optional" json:"spekeKeyProvider" yaml:"spekeKeyProvider"`
}


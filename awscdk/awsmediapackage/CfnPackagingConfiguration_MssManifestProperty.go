package awsmediapackage


// Parameters for a Microsoft Smooth manifest.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   mssManifestProperty := &MssManifestProperty{
//   	ManifestName: jsii.String("manifestName"),
//   	StreamSelection: &StreamSelectionProperty{
//   		MaxVideoBitsPerSecond: jsii.Number(123),
//   		MinVideoBitsPerSecond: jsii.Number(123),
//   		StreamOrder: jsii.String("streamOrder"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-mssmanifest.html
//
type CfnPackagingConfiguration_MssManifestProperty struct {
	// A short string that's appended to the end of the endpoint URL to create a unique path to this packaging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-mssmanifest.html#cfn-mediapackage-packagingconfiguration-mssmanifest-manifestname
	//
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// Video bitrate limitations for outputs from this packaging configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackage-packagingconfiguration-mssmanifest.html#cfn-mediapackage-packagingconfiguration-mssmanifest-streamselection
	//
	StreamSelection interface{} `field:"optional" json:"streamSelection" yaml:"streamSelection"`
}


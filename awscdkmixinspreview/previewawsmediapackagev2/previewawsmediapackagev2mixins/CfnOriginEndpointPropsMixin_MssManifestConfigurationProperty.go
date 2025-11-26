package previewawsmediapackagev2mixins


// <p>Configuration details for a Microsoft Smooth Streaming (MSS) manifest associated with an origin endpoint.
//
// This includes all the settings and properties that define how the MSS content is packaged and delivered.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   mssManifestConfigurationProperty := &MssManifestConfigurationProperty{
//   	FilterConfiguration: &FilterConfigurationProperty{
//   		ClipStartTime: jsii.String("clipStartTime"),
//   		End: jsii.String("end"),
//   		ManifestFilter: jsii.String("manifestFilter"),
//   		Start: jsii.String("start"),
//   		TimeDelaySeconds: jsii.Number(123),
//   	},
//   	ManifestLayout: jsii.String("manifestLayout"),
//   	ManifestName: jsii.String("manifestName"),
//   	ManifestWindowSeconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.html
//
type CfnOriginEndpointPropsMixin_MssManifestConfigurationProperty struct {
	// <p>Filter configuration includes settings for manifest filtering, start and end times, and time delay that apply to all of your egress requests for this manifest.
	//
	// </p>.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-mssmanifestconfiguration-filterconfiguration
	//
	FilterConfiguration interface{} `field:"optional" json:"filterConfiguration" yaml:"filterConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-mssmanifestconfiguration-manifestlayout
	//
	ManifestLayout *string `field:"optional" json:"manifestLayout" yaml:"manifestLayout"`
	// <p>The name of the MSS manifest.
	//
	// This name is appended to the origin endpoint URL to create the unique path for accessing this specific MSS manifest.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-mssmanifestconfiguration-manifestname
	//
	ManifestName *string `field:"optional" json:"manifestName" yaml:"manifestName"`
	// <p>The duration (in seconds) of the manifest window.
	//
	// This represents the total amount of content available in the manifest at any given time.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-mssmanifestconfiguration.html#cfn-mediapackagev2-originendpoint-mssmanifestconfiguration-manifestwindowseconds
	//
	ManifestWindowSeconds *float64 `field:"optional" json:"manifestWindowSeconds" yaml:"manifestWindowSeconds"`
}


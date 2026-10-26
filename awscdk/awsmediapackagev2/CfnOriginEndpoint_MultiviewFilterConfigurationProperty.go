package awsmediapackagev2


// <p>The multiview combination for a pinned manifest.
//
// MediaPackage serves the manifest with this layout and these sources, so players request it without an <code>aws.multiview</code> query parameter.</p> <p>If a request for a pinned manifest also includes an <code>aws.multiview</code> query parameter, MediaPackage rejects the request, even when that parameter requests the same combination.</p>
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   multiviewFilterConfigurationProperty := &MultiviewFilterConfigurationProperty{
//   	Layout: jsii.String("layout"),
//   	Sources: []*string{
//   		jsii.String("sources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-multiviewfilterconfiguration.html
//
type CfnOriginEndpoint_MultiviewFilterConfigurationProperty struct {
	// <p>A tile layout for a multiview channel.
	//
	// Each layout determines how many source tiles are composited into the output and how those tiles are arranged.</p> <p>The allowed values are:</p> <ul> <li> <p> <code>LAYOUT_2EH</code> – Two tiles of equal size, arranged horizontally.</p> </li> <li> <p> <code>LAYOUT_2PL</code> – Two tiles, with one larger primary tile.</p> </li> <li> <p> <code>LAYOUT_3EL</code> – Three tiles of equal size, arranged in two columns.</p> </li> <li> <p> <code>LAYOUT_3PL</code> – Three tiles, with one larger primary tile on the left and two stacked on the right.</p> </li> <li> <p> <code>LAYOUT_4E</code> – Four tiles of equal size, arranged in a two-by-two grid.</p> </li> <li> <p> <code>LAYOUT_4PL</code> – Four tiles, with one larger primary tile on the left and three stacked on the right.</p> </li> </ul>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-multiviewfilterconfiguration.html#cfn-mediapackagev2-originendpoint-multiviewfilterconfiguration-layout
	//
	Layout *string `field:"required" json:"layout" yaml:"layout"`
	// <p>The source channels to composite, in tile order.
	//
	// Each channel must be one of the <code>AvailableSources</code> of the channel that this origin endpoint is on, and the number of channels must equal the number of tiles in <code>Layout</code>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediapackagev2-originendpoint-multiviewfilterconfiguration.html#cfn-mediapackagev2-originendpoint-multiviewfilterconfiguration-sources
	//
	Sources *[]*string `field:"required" json:"sources" yaml:"sources"`
}


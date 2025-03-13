package awsmediatailor


// The configuration for using a content delivery network (CDN), like Amazon CloudFront, for content and ad segment management.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdnConfigurationProperty := &CdnConfigurationProperty{
//   	AdSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   	ContentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html
//
type CfnPlaybackConfiguration_CdnConfigurationProperty struct {
	// A non-default content delivery network (CDN) to serve ad segments.
	//
	// By default, AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings as its CDN for ad segments. To set up an alternate CDN, create a rule in your CDN for the origin ads.mediatailor. *<region>* .amazonaws.com. Then specify the rule's name in this `AdSegmentUrlPrefix` . When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html#cfn-mediatailor-playbackconfiguration-cdnconfiguration-adsegmenturlprefix
	//
	AdSegmentUrlPrefix *string `field:"optional" json:"adSegmentUrlPrefix" yaml:"adSegmentUrlPrefix"`
	// A content delivery network (CDN) to cache content segments, so that content requests don’t always have to go to the origin server.
	//
	// First, create a rule in your CDN for the content segment origin server. Then specify the rule's name in this `ContentSegmentUrlPrefix` . When AWS Elemental MediaTailor serves a manifest, it reports your CDN as the source for content segments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-cdnconfiguration.html#cfn-mediatailor-playbackconfiguration-cdnconfiguration-contentsegmenturlprefix
	//
	ContentSegmentUrlPrefix *string `field:"optional" json:"contentSegmentUrlPrefix" yaml:"contentSegmentUrlPrefix"`
}


package awsmediatailor


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cdnConfigurationProperty := &cdnConfigurationProperty{
//   	adSegmentUrlPrefix: jsii.String("adSegmentUrlPrefix"),
//   	contentSegmentUrlPrefix: jsii.String("contentSegmentUrlPrefix"),
//   }
//
type CfnPlaybackConfiguration_CdnConfigurationProperty struct {
	// `CfnPlaybackConfiguration.CdnConfigurationProperty.AdSegmentUrlPrefix`.
	AdSegmentUrlPrefix *string `field:"optional" json:"adSegmentUrlPrefix" yaml:"adSegmentUrlPrefix"`
	// `CfnPlaybackConfiguration.CdnConfigurationProperty.ContentSegmentUrlPrefix`.
	ContentSegmentUrlPrefix *string `field:"optional" json:"contentSegmentUrlPrefix" yaml:"contentSegmentUrlPrefix"`
}


package awsmediatailor


// The ad decision server (ADS) request timeouts and personalization time budgets for live, VOD, and prefetch workflows.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adsPersonalizationTimeoutsProperty := &AdsPersonalizationTimeoutsProperty{
//   	AdsRequestTimeoutMilliseconds: jsii.Number(123),
//   	LiveMaximumAdsPersonalizationTimeMilliseconds: jsii.Number(123),
//   	PrefetchAdsRequestTimeoutMilliseconds: jsii.Number(123),
//   	PrefetchMaximumAdsPersonalizationTimeMilliseconds: jsii.Number(123),
//   	VodMaximumAdsPersonalizationTimeMilliseconds: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html
//
type CfnPlaybackConfiguration_AdsPersonalizationTimeoutsProperty struct {
	// The maximum time, in milliseconds, that MediaTailor waits for a single ADS response during live or VOD playback.
	//
	// The default is 3000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html#cfn-mediatailor-playbackconfiguration-adspersonalizationtimeouts-adsrequesttimeoutmilliseconds
	//
	AdsRequestTimeoutMilliseconds *float64 `field:"optional" json:"adsRequestTimeoutMilliseconds" yaml:"adsRequestTimeoutMilliseconds"`
	// The maximum total time, in milliseconds, that MediaTailor spends on ADS activity for live manifests.
	//
	// The default is 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html#cfn-mediatailor-playbackconfiguration-adspersonalizationtimeouts-livemaximumadspersonalizationtimemilliseconds
	//
	LiveMaximumAdsPersonalizationTimeMilliseconds *float64 `field:"optional" json:"liveMaximumAdsPersonalizationTimeMilliseconds" yaml:"liveMaximumAdsPersonalizationTimeMilliseconds"`
	// The maximum time, in milliseconds, that MediaTailor waits for a single ADS response during prefetch retrieval.
	//
	// If not set, MediaTailor uses the AdsRequestTimeoutMilliseconds value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html#cfn-mediatailor-playbackconfiguration-adspersonalizationtimeouts-prefetchadsrequesttimeoutmilliseconds
	//
	PrefetchAdsRequestTimeoutMilliseconds *float64 `field:"optional" json:"prefetchAdsRequestTimeoutMilliseconds" yaml:"prefetchAdsRequestTimeoutMilliseconds"`
	// The maximum total time, in milliseconds, that MediaTailor spends on ADS activity during prefetch retrieval.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html#cfn-mediatailor-playbackconfiguration-adspersonalizationtimeouts-prefetchmaximumadspersonalizationtimemilliseconds
	//
	PrefetchMaximumAdsPersonalizationTimeMilliseconds *float64 `field:"optional" json:"prefetchMaximumAdsPersonalizationTimeMilliseconds" yaml:"prefetchMaximumAdsPersonalizationTimeMilliseconds"`
	// The maximum total time, in milliseconds, that MediaTailor spends on ADS activity for VOD manifests.
	//
	// The default is 10000.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationtimeouts.html#cfn-mediatailor-playbackconfiguration-adspersonalizationtimeouts-vodmaximumadspersonalizationtimemilliseconds
	//
	VodMaximumAdsPersonalizationTimeMilliseconds *float64 `field:"optional" json:"vodMaximumAdsPersonalizationTimeMilliseconds" yaml:"vodMaximumAdsPersonalizationTimeMilliseconds"`
}


package awsmediatailor


// The settings that control how many concurrent requests MediaTailor makes to the ad decision server (ADS).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   adsPersonalizationConcurrencyProperty := &AdsPersonalizationConcurrencyProperty{
//   	EnableVodVastParallelization: jsii.Boolean(false),
//   	MaxConcurrentAdsRequests: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationconcurrency.html
//
type CfnPlaybackConfiguration_AdsPersonalizationConcurrencyProperty struct {
	// Enables parallel processing of ADS requests in VOD workflows when the ADS returns VAST responses.
	//
	// The default is false.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationconcurrency.html#cfn-mediatailor-playbackconfiguration-adspersonalizationconcurrency-enablevodvastparallelization
	//
	EnableVodVastParallelization interface{} `field:"optional" json:"enableVodVastParallelization" yaml:"enableVodVastParallelization"`
	// The maximum number of simultaneous requests that MediaTailor makes to the ADS for each manifest request.
	//
	// The default is 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-mediatailor-playbackconfiguration-adspersonalizationconcurrency.html#cfn-mediatailor-playbackconfiguration-adspersonalizationconcurrency-maxconcurrentadsrequests
	//
	MaxConcurrentAdsRequests *float64 `field:"optional" json:"maxConcurrentAdsRequests" yaml:"maxConcurrentAdsRequests"`
}


package previewawsiotfleetwisemixins


// Information about a collection scheme that uses a time period to decide how often to collect data.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeBasedCollectionSchemeProperty := &TimeBasedCollectionSchemeProperty{
//   	PeriodMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedcollectionscheme.html
//
type CfnCampaignPropsMixin_TimeBasedCollectionSchemeProperty struct {
	// The time period (in milliseconds) to decide how often to collect data.
	//
	// For example, if the time period is `60000` , the Edge Agent software collects data once every minute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedcollectionscheme.html#cfn-iotfleetwise-campaign-timebasedcollectionscheme-periodms
	//
	PeriodMs *float64 `field:"optional" json:"periodMs" yaml:"periodMs"`
}


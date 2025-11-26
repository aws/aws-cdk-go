package previewawsiotfleetwisemixins


// Used to configure a frequency-based vehicle signal fetch.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   timeBasedSignalFetchConfigProperty := &TimeBasedSignalFetchConfigProperty{
//   	ExecutionFrequencyMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html
//
type CfnCampaignPropsMixin_TimeBasedSignalFetchConfigProperty struct {
	// The frequency with which the signal fetch will be executed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-timebasedsignalfetchconfig-executionfrequencyms
	//
	ExecutionFrequencyMs *float64 `field:"optional" json:"executionFrequencyMs" yaml:"executionFrequencyMs"`
}


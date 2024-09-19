package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timeBasedSignalFetchConfigProperty := &TimeBasedSignalFetchConfigProperty{
//   	ExecutionFrequencyMs: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html
//
type CfnCampaign_TimeBasedSignalFetchConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-timebasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-timebasedsignalfetchconfig-executionfrequencyms
	//
	ExecutionFrequencyMs *float64 `field:"required" json:"executionFrequencyMs" yaml:"executionFrequencyMs"`
}


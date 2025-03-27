package awsiotfleetwise


// The upload options for the data partition.
//
// If upload options are specified, you must also specify storage options. See [DataPartitionStorageOptions](https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_DataPartitionStorageOptions.html) .
//
// > Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the *AWS IoT FleetWise Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dataPartitionUploadOptionsProperty := &DataPartitionUploadOptionsProperty{
//   	Expression: jsii.String("expression"),
//
//   	// the properties below are optional
//   	ConditionLanguageVersion: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionuploadoptions.html
//
type CfnCampaign_DataPartitionUploadOptionsProperty struct {
	// The logical expression used to recognize what data to collect.
	//
	// For example, `$variable.`Vehicle.OutsideAirTemperature` >= 105.0` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionuploadoptions.html#cfn-iotfleetwise-campaign-datapartitionuploadoptions-expression
	//
	Expression *string `field:"required" json:"expression" yaml:"expression"`
	// The version of the condition language.
	//
	// Defaults to the most recent condition language version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-datapartitionuploadoptions.html#cfn-iotfleetwise-campaign-datapartitionuploadoptions-conditionlanguageversion
	//
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
}


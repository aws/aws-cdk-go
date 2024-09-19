package awsiotfleetwise


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   conditionBasedSignalFetchConfigProperty := &ConditionBasedSignalFetchConfigProperty{
//   	ConditionExpression: jsii.String("conditionExpression"),
//   	TriggerMode: jsii.String("triggerMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html
//
type CfnCampaign_ConditionBasedSignalFetchConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-conditionexpression
	//
	ConditionExpression *string `field:"required" json:"conditionExpression" yaml:"conditionExpression"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-triggermode
	//
	TriggerMode *string `field:"required" json:"triggerMode" yaml:"triggerMode"`
}


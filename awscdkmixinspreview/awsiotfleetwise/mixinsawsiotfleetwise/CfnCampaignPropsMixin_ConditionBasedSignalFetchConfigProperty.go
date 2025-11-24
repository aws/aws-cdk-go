package mixinsawsiotfleetwise


// Specifies the condition under which a signal fetch occurs.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   conditionBasedSignalFetchConfigProperty := &ConditionBasedSignalFetchConfigProperty{
//   	ConditionExpression: jsii.String("conditionExpression"),
//   	TriggerMode: jsii.String("triggerMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html
//
type CfnCampaignPropsMixin_ConditionBasedSignalFetchConfigProperty struct {
	// The condition that must be satisfied to trigger a signal fetch.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-conditionexpression
	//
	ConditionExpression *string `field:"optional" json:"conditionExpression" yaml:"conditionExpression"`
	// Indicates the mode in which the signal fetch is triggered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iotfleetwise-campaign-conditionbasedsignalfetchconfig.html#cfn-iotfleetwise-campaign-conditionbasedsignalfetchconfig-triggermode
	//
	TriggerMode *string `field:"optional" json:"triggerMode" yaml:"triggerMode"`
}


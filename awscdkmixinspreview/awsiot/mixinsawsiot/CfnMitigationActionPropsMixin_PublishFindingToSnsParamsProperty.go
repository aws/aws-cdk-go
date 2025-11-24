package mixinsawsiot


// Parameters to define a mitigation action that publishes findings to Amazon SNS.
//
// You can implement your own custom actions in response to the Amazon SNS messages.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   publishFindingToSnsParamsProperty := &PublishFindingToSnsParamsProperty{
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-publishfindingtosnsparams.html
//
type CfnMitigationActionPropsMixin_PublishFindingToSnsParamsProperty struct {
	// The ARN of the topic to which you want to publish the findings.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-mitigationaction-publishfindingtosnsparams.html#cfn-iot-mitigationaction-publishfindingtosnsparams-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}


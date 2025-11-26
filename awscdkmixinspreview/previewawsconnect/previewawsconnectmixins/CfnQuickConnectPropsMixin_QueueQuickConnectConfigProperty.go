package previewawsconnectmixins


// Contains information about a queue for a quick connect.
//
// The flow must be of type Transfer to Queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   queueQuickConnectConfigProperty := &QueueQuickConnectConfigProperty{
//   	ContactFlowArn: jsii.String("contactFlowArn"),
//   	QueueArn: jsii.String("queueArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html
//
type CfnQuickConnectPropsMixin_QueueQuickConnectConfigProperty struct {
	// The Amazon Resource Name (ARN) of the flow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-contactflowarn
	//
	ContactFlowArn *string `field:"optional" json:"contactFlowArn" yaml:"contactFlowArn"`
	// The Amazon Resource Name (ARN) of the queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-quickconnect-queuequickconnectconfig.html#cfn-connect-quickconnect-queuequickconnectconfig-queuearn
	//
	QueueArn *string `field:"optional" json:"queueArn" yaml:"queueArn"`
}


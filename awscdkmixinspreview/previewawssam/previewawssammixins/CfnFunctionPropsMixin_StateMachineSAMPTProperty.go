package previewawssammixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stateMachineSAMPTProperty := &StateMachineSAMPTProperty{
//   	StateMachineName: jsii.String("stateMachineName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-statemachinesampt.html
//
type CfnFunctionPropsMixin_StateMachineSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-statemachinesampt.html#cfn-serverless-function-statemachinesampt-statemachinename
	//
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
}


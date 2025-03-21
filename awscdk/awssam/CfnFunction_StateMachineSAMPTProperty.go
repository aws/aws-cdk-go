package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   stateMachineSAMPTProperty := &StateMachineSAMPTProperty{
//   	StateMachineName: jsii.String("stateMachineName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-statemachinesampt.html
//
type CfnFunction_StateMachineSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-statemachinesampt.html#cfn-serverless-function-statemachinesampt-statemachinename
	//
	StateMachineName *string `field:"required" json:"stateMachineName" yaml:"stateMachineName"`
}


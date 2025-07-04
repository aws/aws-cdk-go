package awsapptest


// Specifies the CloudFormation action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cloudFormationActionProperty := &CloudFormationActionProperty{
//   	Resource: jsii.String("resource"),
//
//   	// the properties below are optional
//   	ActionType: jsii.String("actionType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html
//
type CfnTestCase_CloudFormationActionProperty struct {
	// The resource of the CloudFormation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-resource
	//
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The action type of the CloudFormation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
}


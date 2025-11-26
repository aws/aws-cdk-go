package previewawsapptestmixins


// Specifies the CloudFormation action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cloudFormationActionProperty := &CloudFormationActionProperty{
//   	ActionType: jsii.String("actionType"),
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html
//
type CfnTestCasePropsMixin_CloudFormationActionProperty struct {
	// The action type of the CloudFormation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// The resource of the CloudFormation action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-cloudformationaction.html#cfn-apptest-testcase-cloudformationaction-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}


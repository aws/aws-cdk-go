package previewawsapptestmixins


// Specifies the Mainframe Modernization non-managed application action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   m2NonManagedApplicationActionProperty := &M2NonManagedApplicationActionProperty{
//   	ActionType: jsii.String("actionType"),
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2nonmanagedapplicationaction.html
//
type CfnTestCasePropsMixin_M2NonManagedApplicationActionProperty struct {
	// The action type of the Mainframe Modernization non-managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2nonmanagedapplicationaction.html#cfn-apptest-testcase-m2nonmanagedapplicationaction-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// The resource of the Mainframe Modernization non-managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2nonmanagedapplicationaction.html#cfn-apptest-testcase-m2nonmanagedapplicationaction-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}


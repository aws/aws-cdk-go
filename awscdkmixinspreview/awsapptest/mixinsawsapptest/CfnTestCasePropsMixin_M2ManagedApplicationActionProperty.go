package mixinsawsapptest


// Specifies the Mainframe Modernization managed application action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   m2ManagedApplicationActionProperty := &M2ManagedApplicationActionProperty{
//   	ActionType: jsii.String("actionType"),
//   	Properties: &M2ManagedActionPropertiesProperty{
//   		ForceStop: jsii.Boolean(false),
//   		ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   	},
//   	Resource: jsii.String("resource"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html
//
type CfnTestCasePropsMixin_M2ManagedApplicationActionProperty struct {
	// The action type of the Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-actiontype
	//
	ActionType *string `field:"optional" json:"actionType" yaml:"actionType"`
	// The properties of the Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
	// The resource of the Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-resource
	//
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}


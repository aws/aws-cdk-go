package previewawsapptestmixins


// Specifies a resource action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   resourceActionProperty := &ResourceActionProperty{
//   	CloudFormationAction: &CloudFormationActionProperty{
//   		ActionType: jsii.String("actionType"),
//   		Resource: jsii.String("resource"),
//   	},
//   	M2ManagedApplicationAction: &M2ManagedApplicationActionProperty{
//   		ActionType: jsii.String("actionType"),
//   		Properties: &M2ManagedActionPropertiesProperty{
//   			ForceStop: jsii.Boolean(false),
//   			ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   		},
//   		Resource: jsii.String("resource"),
//   	},
//   	M2NonManagedApplicationAction: &M2NonManagedApplicationActionProperty{
//   		ActionType: jsii.String("actionType"),
//   		Resource: jsii.String("resource"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html
//
type CfnTestCasePropsMixin_ResourceActionProperty struct {
	// The CloudFormation action of the resource action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-cloudformationaction
	//
	CloudFormationAction interface{} `field:"optional" json:"cloudFormationAction" yaml:"cloudFormationAction"`
	// The Mainframe Modernization managed application action of the resource action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-m2managedapplicationaction
	//
	M2ManagedApplicationAction interface{} `field:"optional" json:"m2ManagedApplicationAction" yaml:"m2ManagedApplicationAction"`
	// The Mainframe Modernization non-managed application action of the resource action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-resourceaction.html#cfn-apptest-testcase-resourceaction-m2nonmanagedapplicationaction
	//
	M2NonManagedApplicationAction interface{} `field:"optional" json:"m2NonManagedApplicationAction" yaml:"m2NonManagedApplicationAction"`
}


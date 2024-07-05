package awsapptest


// Specifies the AWS Mainframe Modernization managed application action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   m2ManagedApplicationActionProperty := &M2ManagedApplicationActionProperty{
//   	ActionType: jsii.String("actionType"),
//   	Resource: jsii.String("resource"),
//
//   	// the properties below are optional
//   	Properties: &M2ManagedActionPropertiesProperty{
//   		ForceStop: jsii.Boolean(false),
//   		ImportDataSetLocation: jsii.String("importDataSetLocation"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html
//
type CfnTestCase_M2ManagedApplicationActionProperty struct {
	// The action type of the AWS Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-actiontype
	//
	ActionType *string `field:"required" json:"actionType" yaml:"actionType"`
	// The resource of the AWS Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-resource
	//
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// The properties of the AWS Mainframe Modernization managed application action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-m2managedapplicationaction.html#cfn-apptest-testcase-m2managedapplicationaction-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}


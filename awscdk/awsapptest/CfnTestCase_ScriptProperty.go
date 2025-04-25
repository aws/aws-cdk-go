package awsapptest


// Specifies the script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scriptProperty := &ScriptProperty{
//   	ScriptLocation: jsii.String("scriptLocation"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html
//
type CfnTestCase_ScriptProperty struct {
	// The script location of the scripts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html#cfn-apptest-testcase-script-scriptlocation
	//
	ScriptLocation *string `field:"required" json:"scriptLocation" yaml:"scriptLocation"`
	// The type of the scripts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html#cfn-apptest-testcase-script-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


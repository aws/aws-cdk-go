package mixinsawsapptest


// Specifies the script.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   scriptProperty := &ScriptProperty{
//   	ScriptLocation: jsii.String("scriptLocation"),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html
//
type CfnTestCasePropsMixin_ScriptProperty struct {
	// The script location of the scripts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html#cfn-apptest-testcase-script-scriptlocation
	//
	ScriptLocation *string `field:"optional" json:"scriptLocation" yaml:"scriptLocation"`
	// The type of the scripts.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-script.html#cfn-apptest-testcase-script-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


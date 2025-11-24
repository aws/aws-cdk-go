package mixinsawsapptest


// Specifies the latest version of a test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   testCaseLatestVersionProperty := &TestCaseLatestVersionProperty{
//   	Status: jsii.String("status"),
//   	Version: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html
//
type CfnTestCasePropsMixin_TestCaseLatestVersionProperty struct {
	// The status of the test case latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html#cfn-apptest-testcase-testcaselatestversion-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The version of the test case latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html#cfn-apptest-testcase-testcaselatestversion-version
	//
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}


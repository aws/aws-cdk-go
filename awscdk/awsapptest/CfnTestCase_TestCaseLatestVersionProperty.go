package awsapptest


// Specifies the latest version of a test case.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   testCaseLatestVersionProperty := &TestCaseLatestVersionProperty{
//   	Status: jsii.String("status"),
//   	Version: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html
//
type CfnTestCase_TestCaseLatestVersionProperty struct {
	// The status of the test case latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html#cfn-apptest-testcase-testcaselatestversion-status
	//
	Status *string `field:"required" json:"status" yaml:"status"`
	// The version of the test case latest version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apptest-testcase-testcaselatestversion.html#cfn-apptest-testcase-testcaselatestversion-version
	//
	Version *float64 `field:"required" json:"version" yaml:"version"`
}


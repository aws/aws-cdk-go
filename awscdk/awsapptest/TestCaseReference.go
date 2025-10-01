package awsapptest


// A reference to a TestCase resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   testCaseReference := &TestCaseReference{
//   	TestCaseArn: jsii.String("testCaseArn"),
//   	TestCaseId: jsii.String("testCaseId"),
//   }
//
type TestCaseReference struct {
	// The ARN of the TestCase resource.
	TestCaseArn *string `field:"required" json:"testCaseArn" yaml:"testCaseArn"`
	// The TestCaseId of the TestCase resource.
	TestCaseId *string `field:"required" json:"testCaseId" yaml:"testCaseId"`
}


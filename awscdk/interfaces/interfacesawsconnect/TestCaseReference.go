package interfacesawsconnect


// A reference to a TestCase resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   testCaseReference := &TestCaseReference{
//   	TestCaseArn: jsii.String("testCaseArn"),
//   }
//
type TestCaseReference struct {
	// The TestCaseArn of the TestCase resource.
	TestCaseArn *string `field:"required" json:"testCaseArn" yaml:"testCaseArn"`
}


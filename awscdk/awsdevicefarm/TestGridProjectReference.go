package awsdevicefarm


// A reference to a TestGridProject resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   testGridProjectReference := &TestGridProjectReference{
//   	TestGridProjectArn: jsii.String("testGridProjectArn"),
//   }
//
type TestGridProjectReference struct {
	// The Arn of the TestGridProject resource.
	TestGridProjectArn *string `field:"required" json:"testGridProjectArn" yaml:"testGridProjectArn"`
}


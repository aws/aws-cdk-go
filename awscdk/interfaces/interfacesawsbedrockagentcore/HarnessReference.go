package interfacesawsbedrockagentcore


// A reference to a Harness resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessReference := &HarnessReference{
//   	HarnessArn: jsii.String("harnessArn"),
//   }
//
type HarnessReference struct {
	// The Arn of the Harness resource.
	HarnessArn *string `field:"required" json:"harnessArn" yaml:"harnessArn"`
}


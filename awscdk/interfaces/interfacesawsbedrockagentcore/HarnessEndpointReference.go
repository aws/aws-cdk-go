package interfacesawsbedrockagentcore


// A reference to a HarnessEndpoint resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessEndpointReference := &HarnessEndpointReference{
//   	HarnessEndpointArn: jsii.String("harnessEndpointArn"),
//   }
//
type HarnessEndpointReference struct {
	// The Arn of the HarnessEndpoint resource.
	HarnessEndpointArn *string `field:"required" json:"harnessEndpointArn" yaml:"harnessEndpointArn"`
}


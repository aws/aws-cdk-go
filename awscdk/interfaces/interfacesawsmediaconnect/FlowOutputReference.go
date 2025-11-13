package interfacesawsmediaconnect


// A reference to a FlowOutput resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   flowOutputReference := &FlowOutputReference{
//   	OutputArn: jsii.String("outputArn"),
//   }
//
type FlowOutputReference struct {
	// The OutputArn of the FlowOutput resource.
	OutputArn *string `field:"required" json:"outputArn" yaml:"outputArn"`
}


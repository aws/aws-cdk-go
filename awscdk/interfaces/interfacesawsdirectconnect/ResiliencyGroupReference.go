package interfacesawsdirectconnect


// A reference to a ResiliencyGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resiliencyGroupReference := &ResiliencyGroupReference{
//   	ResiliencyGroupArn: jsii.String("resiliencyGroupArn"),
//   }
//
type ResiliencyGroupReference struct {
	// The ResiliencyGroupArn of the ResiliencyGroup resource.
	ResiliencyGroupArn *string `field:"required" json:"resiliencyGroupArn" yaml:"resiliencyGroupArn"`
}


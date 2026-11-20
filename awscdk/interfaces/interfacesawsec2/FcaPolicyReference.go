package interfacesawsec2


// A reference to a FcaPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fcaPolicyReference := &FcaPolicyReference{
//   	FcaPolicyArn: jsii.String("fcaPolicyArn"),
//   }
//
type FcaPolicyReference struct {
	// The Arn of the FcaPolicy resource.
	FcaPolicyArn *string `field:"required" json:"fcaPolicyArn" yaml:"fcaPolicyArn"`
}


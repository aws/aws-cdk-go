package interfacesawscloudformation


// A reference to a TypeActivation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   typeActivationReference := &TypeActivationReference{
//   	TypeActivationArn: jsii.String("typeActivationArn"),
//   }
//
type TypeActivationReference struct {
	// The Arn of the TypeActivation resource.
	TypeActivationArn *string `field:"required" json:"typeActivationArn" yaml:"typeActivationArn"`
}


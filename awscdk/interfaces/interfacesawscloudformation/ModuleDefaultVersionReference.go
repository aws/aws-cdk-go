package interfacesawscloudformation


// A reference to a ModuleDefaultVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   moduleDefaultVersionReference := &ModuleDefaultVersionReference{
//   	ModuleDefaultVersionArn: jsii.String("moduleDefaultVersionArn"),
//   }
//
type ModuleDefaultVersionReference struct {
	// The Arn of the ModuleDefaultVersion resource.
	ModuleDefaultVersionArn *string `field:"required" json:"moduleDefaultVersionArn" yaml:"moduleDefaultVersionArn"`
}


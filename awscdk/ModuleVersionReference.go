package awscdk


// A reference to a ModuleVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//
//   moduleVersionReference := &ModuleVersionReference{
//   	ModuleVersionArn: jsii.String("moduleVersionArn"),
//   }
//
type ModuleVersionReference struct {
	// The Arn of the ModuleVersion resource.
	ModuleVersionArn *string `field:"required" json:"moduleVersionArn" yaml:"moduleVersionArn"`
}


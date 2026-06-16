package interfacesawsresiliencehubv2


// A reference to a System resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   systemReference := &SystemReference{
//   	SystemArn: jsii.String("systemArn"),
//   }
//
type SystemReference struct {
	// The SystemArn of the System resource.
	SystemArn *string `field:"required" json:"systemArn" yaml:"systemArn"`
}


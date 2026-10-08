package interfacesawsredshift


// A reference to a RedshiftIdcApplication resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftIdcApplicationReference := &RedshiftIdcApplicationReference{
//   	RedshiftIdcApplicationArn: jsii.String("redshiftIdcApplicationArn"),
//   }
//
type RedshiftIdcApplicationReference struct {
	// The RedshiftIdcApplicationArn of the RedshiftIdcApplication resource.
	RedshiftIdcApplicationArn *string `field:"required" json:"redshiftIdcApplicationArn" yaml:"redshiftIdcApplicationArn"`
}


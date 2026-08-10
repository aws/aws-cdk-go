package interfacesawsssm


// A reference to a OpsItem resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   opsItemReference := &OpsItemReference{
//   	OpsItemArn: jsii.String("opsItemArn"),
//   }
//
type OpsItemReference struct {
	// The OpsItemArn of the OpsItem resource.
	OpsItemArn *string `field:"required" json:"opsItemArn" yaml:"opsItemArn"`
}


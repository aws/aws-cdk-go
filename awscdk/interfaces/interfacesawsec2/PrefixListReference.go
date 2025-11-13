package interfacesawsec2


// A reference to a PrefixList resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prefixListReference := &PrefixListReference{
//   	PrefixListArn: jsii.String("prefixListArn"),
//   	PrefixListId: jsii.String("prefixListId"),
//   }
//
type PrefixListReference struct {
	// The ARN of the PrefixList resource.
	PrefixListArn *string `field:"required" json:"prefixListArn" yaml:"prefixListArn"`
	// The PrefixListId of the PrefixList resource.
	PrefixListId *string `field:"required" json:"prefixListId" yaml:"prefixListId"`
}


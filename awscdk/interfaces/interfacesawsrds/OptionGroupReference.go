package interfacesawsrds


// A reference to a OptionGroup resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   optionGroupReference := &OptionGroupReference{
//   	OptionGroupName: jsii.String("optionGroupName"),
//   }
//
type OptionGroupReference struct {
	// The OptionGroupName of the OptionGroup resource.
	OptionGroupName *string `field:"required" json:"optionGroupName" yaml:"optionGroupName"`
}


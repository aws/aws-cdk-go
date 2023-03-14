package awslex


// Specifies the default value to use when a user doesn't provide a value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotDefaultValueProperty := &slotDefaultValueProperty{
//   	defaultValue: jsii.String("defaultValue"),
//   }
//
type CfnBot_SlotDefaultValueProperty struct {
	// The default value to use when a user doesn't provide a value for a slot.
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
}


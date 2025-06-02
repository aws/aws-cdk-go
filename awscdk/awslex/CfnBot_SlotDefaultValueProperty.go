package awslex


// Specifies the default value to use when a user doesn't provide a value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotDefaultValueProperty := &SlotDefaultValueProperty{
//   	DefaultValue: jsii.String("defaultValue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotdefaultvalue.html
//
type CfnBot_SlotDefaultValueProperty struct {
	// The default value to use when a user doesn't provide a value for a slot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotdefaultvalue.html#cfn-lex-bot-slotdefaultvalue-defaultvalue
	//
	DefaultValue *string `field:"required" json:"defaultValue" yaml:"defaultValue"`
}


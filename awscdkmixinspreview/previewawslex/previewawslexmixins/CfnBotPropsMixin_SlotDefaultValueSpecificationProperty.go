package previewawslexmixins


// The default value to use when a user doesn't provide a value for a slot.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   slotDefaultValueSpecificationProperty := &SlotDefaultValueSpecificationProperty{
//   	DefaultValueList: []interface{}{
//   		&SlotDefaultValueProperty{
//   			DefaultValue: jsii.String("defaultValue"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotdefaultvaluespecification.html
//
type CfnBotPropsMixin_SlotDefaultValueSpecificationProperty struct {
	// A list of default values.
	//
	// Amazon Lex chooses the default value to use in the order that they are presented in the list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lex-bot-slotdefaultvaluespecification.html#cfn-lex-bot-slotdefaultvaluespecification-defaultvaluelist
	//
	DefaultValueList interface{} `field:"optional" json:"defaultValueList" yaml:"defaultValueList"`
}


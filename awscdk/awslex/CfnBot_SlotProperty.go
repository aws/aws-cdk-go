package awslex


// Specifies the definition of a slot.
//
// Amazon Lex elicits slot values from uses to fulfill the user's intent.
//
// Example:
//
//
type CfnBot_SlotProperty struct {
	// The name given to the slot.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The name of the slot type that this slot is based on.
	//
	// The slot type defines the acceptable values for the slot.
	SlotTypeName *string `field:"required" json:"slotTypeName" yaml:"slotTypeName"`
	// Determines the slot resolution strategy that Amazon Lex uses to return slot type values.
	//
	// The field can be set to one of the following values:
	//
	// - ORIGINAL_VALUE - Returns the value entered by the user, if the user value is similar to a slot value.
	// - TOP_RESOLUTION - If there is a resolution list for the slot, return the first value in the resolution list as the slot type value. If there is no resolution list, null is returned.
	//
	// If you don't specify the `valueSelectionStrategy` , the default is `ORIGINAL_VALUE` .
	ValueElicitationSetting interface{} `field:"required" json:"valueElicitationSetting" yaml:"valueElicitationSetting"`
	// The description of the slot.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether a slot can return multiple values.
	MultipleValuesSetting interface{} `field:"optional" json:"multipleValuesSetting" yaml:"multipleValuesSetting"`
	// Determines whether the contents of the slot are obfuscated in Amazon CloudWatch Logs logs.
	//
	// Use obfuscated slots to protect information such as personally identifiable information (PII) in logs.
	ObfuscationSetting interface{} `field:"optional" json:"obfuscationSetting" yaml:"obfuscationSetting"`
}


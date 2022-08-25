package awslex


// Each slot type can have a set of values.
//
// The `SlotTypeValue` represents a value that the slot type can take.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   slotTypeValueProperty := &slotTypeValueProperty{
//   	sampleValue: &sampleValueProperty{
//   		value: jsii.String("value"),
//   	},
//
//   	// the properties below are optional
//   	synonyms: []interface{}{
//   		&sampleValueProperty{
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBot_SlotTypeValueProperty struct {
	// The value of the slot type entry.
	SampleValue interface{} `field:"required" json:"sampleValue" yaml:"sampleValue"`
	// Additional values related to the slot type entry.
	Synonyms interface{} `field:"optional" json:"synonyms" yaml:"synonyms"`
}


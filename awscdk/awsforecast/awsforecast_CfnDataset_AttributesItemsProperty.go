package awsforecast


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   attributesItemsProperty := &attributesItemsProperty{
//   	attributeName: jsii.String("attributeName"),
//   	attributeType: jsii.String("attributeType"),
//   }
//
type CfnDataset_AttributesItemsProperty struct {
	// `CfnDataset.AttributesItemsProperty.AttributeName`.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// `CfnDataset.AttributesItemsProperty.AttributeType`.
	AttributeType *string `field:"optional" json:"attributeType" yaml:"attributeType"`
}


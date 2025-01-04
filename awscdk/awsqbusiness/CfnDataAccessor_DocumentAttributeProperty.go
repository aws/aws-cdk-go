package awsqbusiness


// A document attribute or metadata field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   documentAttributeProperty := &DocumentAttributeProperty{
//   	Name: jsii.String("name"),
//   	Value: &DocumentAttributeValueProperty{
//   		DateValue: jsii.String("dateValue"),
//   		LongValue: jsii.Number(123),
//   		StringListValue: []*string{
//   			jsii.String("stringListValue"),
//   		},
//   		StringValue: jsii.String("stringValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattribute.html
//
type CfnDataAccessor_DocumentAttributeProperty struct {
	// The identifier for the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattribute.html#cfn-qbusiness-dataaccessor-documentattribute-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value of the attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-dataaccessor-documentattribute.html#cfn-qbusiness-dataaccessor-documentattribute-value
	//
	Value interface{} `field:"required" json:"value" yaml:"value"`
}


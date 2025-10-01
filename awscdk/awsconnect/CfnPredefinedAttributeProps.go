package awsconnect


// Properties for defining a `CfnPredefinedAttribute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPredefinedAttributeProps := &CfnPredefinedAttributeProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AttributeConfiguration: &AttributeConfigurationProperty{
//   		EnableValueValidationOnAssociation: jsii.Boolean(false),
//   		IsReadOnly: jsii.Boolean(false),
//   	},
//   	Purposes: []*string{
//   		jsii.String("purposes"),
//   	},
//   	Values: &ValuesProperty{
//   		StringList: []*string{
//   			jsii.String("stringList"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html
//
type CfnPredefinedAttributeProps struct {
	// The Amazon Resource Name (ARN) of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-instancearn
	//
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the predefined attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Custom metadata that is associated to predefined attributes to control behavior in upstream services, such as controlling how a predefined attribute should be displayed in the Amazon Connect admin website.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-attributeconfiguration
	//
	AttributeConfiguration interface{} `field:"optional" json:"attributeConfiguration" yaml:"attributeConfiguration"`
	// Values that enable you to categorize your predefined attributes.
	//
	// You can use them in custom UI elements across the Amazon Connect admin website.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-purposes
	//
	Purposes *[]*string `field:"optional" json:"purposes" yaml:"purposes"`
	// The values of a predefined attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}


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
	// The values of a predefined attribute.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connect-predefinedattribute.html#cfn-connect-predefinedattribute-values
	//
	Values interface{} `field:"optional" json:"values" yaml:"values"`
}


package awsconnect


// The values of a predefined attribute.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   valuesProperty := &ValuesProperty{
//   	StringList: []*string{
//   		jsii.String("stringList"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-predefinedattribute-values.html
//
type CfnPredefinedAttribute_ValuesProperty struct {
	// Predefined attribute values of type string list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-predefinedattribute-values.html#cfn-connect-predefinedattribute-values-stringlist
	//
	StringList *[]*string `field:"optional" json:"stringList" yaml:"stringList"`
}


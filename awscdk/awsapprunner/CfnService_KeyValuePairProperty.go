package awsapprunner


// Describes a key-value pair, which is a string-to-string mapping.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   keyValuePairProperty := &KeyValuePairProperty{
//   	Name: jsii.String("name"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-keyvaluepair.html
//
type CfnService_KeyValuePairProperty struct {
	// The key name string to map to a value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-keyvaluepair.html#cfn-apprunner-service-keyvaluepair-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value string to which the key name is mapped.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apprunner-service-keyvaluepair.html#cfn-apprunner-service-keyvaluepair-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}


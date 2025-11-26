package awsathena


// A classification refers to a set of specific configurations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   classificationProperty := &ClassificationProperty{
//   	Name: jsii.String("name"),
//   	Properties: map[string]*string{
//   		"propertiesKey": jsii.String("properties"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-classification.html
//
type CfnWorkGroup_ClassificationProperty struct {
	// The name of the configuration classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-classification.html#cfn-athena-workgroup-classification-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A set of properties specified within a configuration classification.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-athena-workgroup-classification.html#cfn-athena-workgroup-classification-properties
	//
	Properties interface{} `field:"optional" json:"properties" yaml:"properties"`
}


package awscomprehend


// An entity type within a labeled training dataset that Amazon Comprehend uses to train a custom entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityTypesListItemProperty := &EntityTypesListItemProperty{
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entitytypeslistitem.html
//
type CfnEntityRecognizer_EntityTypesListItemProperty struct {
	// An entity type within a labeled training dataset.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-entityrecognizer-entitytypeslistitem.html#cfn-comprehend-entityrecognizer-entitytypeslistitem-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-entitytypeslistitem.html
//
type CfnFlywheel_EntityTypesListItemProperty struct {
	// An entity type within a labeled training dataset that Amazon Comprehend uses to train a custom entity recognizer.
	//
	// Entity types must not contain the following invalid characters: \n (line break), \\n (escaped line break, \r (carriage return), \\r (escaped carriage return), \t (tab), \\t (escaped tab), and , (comma).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-entitytypeslistitem.html#cfn-comprehend-flywheel-entitytypeslistitem-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
}


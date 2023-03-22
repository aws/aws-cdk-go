package awscomprehend


// An entity type within a labeled training dataset that Amazon Comprehend uses to train a custom entity recognizer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityTypesListItemProperty := &entityTypesListItemProperty{
//   	type: jsii.String("type"),
//   }
//
type CfnFlywheel_EntityTypesListItemProperty struct {
	// An entity type within a labeled training dataset that Amazon Comprehend uses to train a custom entity recognizer.
	//
	// Entity types must not contain the following invalid characters: \n (line break), \\n (escaped line break, \r (carriage return), \\r (escaped carriage return), \t (tab), \\t (escaped tab), space, and , (comma).
	Type *string `field:"required" json:"type" yaml:"type"`
}


package awscomprehend


// Configuration required for an entity recognition model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognitionConfigProperty := &entityRecognitionConfigProperty{
//   	entityTypes: []interface{}{
//   		&entityTypesListItemProperty{
//   			type: jsii.String("type"),
//   		},
//   	},
//   }
//
type CfnFlywheel_EntityRecognitionConfigProperty struct {
	// Up to 25 entity types that the model is trained to recognize.
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
}


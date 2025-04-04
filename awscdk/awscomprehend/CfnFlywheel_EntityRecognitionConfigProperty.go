package awscomprehend


// Configuration required for an entity recognition model.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   entityRecognitionConfigProperty := &EntityRecognitionConfigProperty{
//   	EntityTypes: []interface{}{
//   		&EntityTypesListItemProperty{
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-entityrecognitionconfig.html
//
type CfnFlywheel_EntityRecognitionConfigProperty struct {
	// Up to 25 entity types that the model is trained to recognize.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-entityrecognitionconfig.html#cfn-comprehend-flywheel-entityrecognitionconfig-entitytypes
	//
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
}


package awscomprehend


// Configuration about the model associated with a flywheel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskConfigProperty := &TaskConfigProperty{
//   	LanguageCode: jsii.String("languageCode"),
//
//   	// the properties below are optional
//   	DocumentClassificationConfig: &DocumentClassificationConfigProperty{
//   		Mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		Labels: []*string{
//   			jsii.String("labels"),
//   		},
//   	},
//   	EntityRecognitionConfig: &EntityRecognitionConfigProperty{
//   		EntityTypes: []interface{}{
//   			&EntityTypesListItemProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html
//
type CfnFlywheel_TaskConfigProperty struct {
	// Language code for the language that the model supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-languagecode
	//
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Configuration required for a document classification model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-documentclassificationconfig
	//
	DocumentClassificationConfig interface{} `field:"optional" json:"documentClassificationConfig" yaml:"documentClassificationConfig"`
	// Configuration required for an entity recognition model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-entityrecognitionconfig
	//
	EntityRecognitionConfig interface{} `field:"optional" json:"entityRecognitionConfig" yaml:"entityRecognitionConfig"`
}


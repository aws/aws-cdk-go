package mixinsawscomprehend


// Configuration about the model associated with a flywheel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   taskConfigProperty := &TaskConfigProperty{
//   	DocumentClassificationConfig: &DocumentClassificationConfigProperty{
//   		Labels: []*string{
//   			jsii.String("labels"),
//   		},
//   		Mode: jsii.String("mode"),
//   	},
//   	EntityRecognitionConfig: &EntityRecognitionConfigProperty{
//   		EntityTypes: []interface{}{
//   			&EntityTypesListItemProperty{
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	LanguageCode: jsii.String("languageCode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html
//
type CfnFlywheelPropsMixin_TaskConfigProperty struct {
	// Configuration required for a document classification model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-documentclassificationconfig
	//
	DocumentClassificationConfig interface{} `field:"optional" json:"documentClassificationConfig" yaml:"documentClassificationConfig"`
	// Configuration required for an entity recognition model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-entityrecognitionconfig
	//
	EntityRecognitionConfig interface{} `field:"optional" json:"entityRecognitionConfig" yaml:"entityRecognitionConfig"`
	// Language code for the language that the model supports.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-comprehend-flywheel-taskconfig.html#cfn-comprehend-flywheel-taskconfig-languagecode
	//
	LanguageCode *string `field:"optional" json:"languageCode" yaml:"languageCode"`
}


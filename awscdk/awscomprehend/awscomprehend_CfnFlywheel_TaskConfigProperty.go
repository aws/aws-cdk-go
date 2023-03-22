package awscomprehend


// Configuration about the custom classifier associated with the flywheel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   taskConfigProperty := &taskConfigProperty{
//   	languageCode: jsii.String("languageCode"),
//
//   	// the properties below are optional
//   	documentClassificationConfig: &documentClassificationConfigProperty{
//   		mode: jsii.String("mode"),
//
//   		// the properties below are optional
//   		labels: []*string{
//   			jsii.String("labels"),
//   		},
//   	},
//   	entityRecognitionConfig: &entityRecognitionConfigProperty{
//   		entityTypes: []interface{}{
//   			&entityTypesListItemProperty{
//   				type: jsii.String("type"),
//   			},
//   		},
//   	},
//   }
//
type CfnFlywheel_TaskConfigProperty struct {
	// Language code for the language that the model supports.
	LanguageCode *string `field:"required" json:"languageCode" yaml:"languageCode"`
	// Configuration required for a classification model.
	DocumentClassificationConfig interface{} `field:"optional" json:"documentClassificationConfig" yaml:"documentClassificationConfig"`
	// Configuration required for an entity recognition model.
	EntityRecognitionConfig interface{} `field:"optional" json:"entityRecognitionConfig" yaml:"entityRecognitionConfig"`
}


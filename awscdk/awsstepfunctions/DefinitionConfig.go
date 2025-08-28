package awsstepfunctions


// Partial object from the StateMachine L1 construct properties containing definition information.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var definition interface{}
//
//   definitionConfig := &DefinitionConfig{
//   	Definition: definition,
//   	DefinitionS3Location: &S3LocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   	DefinitionString: jsii.String("definitionString"),
//   }
//
type DefinitionConfig struct {
	Definition interface{} `field:"optional" json:"definition" yaml:"definition"`
	DefinitionS3Location *CfnStateMachine_S3LocationProperty `field:"optional" json:"definitionS3Location" yaml:"definitionS3Location"`
	DefinitionString *string `field:"optional" json:"definitionString" yaml:"definitionString"`
}


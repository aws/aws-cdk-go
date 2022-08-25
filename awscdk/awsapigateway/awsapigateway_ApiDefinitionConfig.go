package awsapigateway


// Post-Binding Configuration for a CDK construct.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var inlineDefinition interface{}
//
//   apiDefinitionConfig := &apiDefinitionConfig{
//   	inlineDefinition: inlineDefinition,
//   	s3Location: &apiDefinitionS3Location{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//
//   		// the properties below are optional
//   		version: jsii.String("version"),
//   	},
//   }
//
type ApiDefinitionConfig struct {
	// Inline specification (mutually exclusive with `s3Location`).
	InlineDefinition interface{} `field:"optional" json:"inlineDefinition" yaml:"inlineDefinition"`
	// The location of the specification in S3 (mutually exclusive with `inlineDefinition`).
	S3Location *ApiDefinitionS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}


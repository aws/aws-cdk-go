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
//   apiDefinitionConfig := &ApiDefinitionConfig{
//   	InlineDefinition: inlineDefinition,
//   	S3Location: &ApiDefinitionS3Location{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//
//   		// the properties below are optional
//   		Version: jsii.String("version"),
//   	},
//   }
//
type ApiDefinitionConfig struct {
	// Inline specification (mutually exclusive with `s3Location`).
	// Default: - API definition is not defined inline.
	//
	InlineDefinition interface{} `field:"optional" json:"inlineDefinition" yaml:"inlineDefinition"`
	// The location of the specification in S3 (mutually exclusive with `inlineDefinition`).
	// Default: - API definition is not an S3 location.
	//
	S3Location *ApiDefinitionS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}


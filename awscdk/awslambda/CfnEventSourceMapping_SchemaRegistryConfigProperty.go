package awslambda


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   schemaRegistryConfigProperty := &SchemaRegistryConfigProperty{
//   	AccessConfigs: []interface{}{
//   		&SchemaRegistryAccessConfigProperty{
//   			Type: jsii.String("type"),
//   			Uri: jsii.String("uri"),
//   		},
//   	},
//   	EventRecordFormat: jsii.String("eventRecordFormat"),
//   	SchemaRegistryUri: jsii.String("schemaRegistryUri"),
//   	SchemaValidationConfigs: []interface{}{
//   		&SchemaValidationConfigProperty{
//   			Attribute: jsii.String("attribute"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemaregistryconfig.html
//
type CfnEventSourceMapping_SchemaRegistryConfigProperty struct {
	// An array of access configuration objects that tell Lambda how to authenticate with your schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemaregistryconfig.html#cfn-lambda-eventsourcemapping-schemaregistryconfig-accessconfigs
	//
	AccessConfigs interface{} `field:"optional" json:"accessConfigs" yaml:"accessConfigs"`
	// The record format that Lambda delivers to your function after schema validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemaregistryconfig.html#cfn-lambda-eventsourcemapping-schemaregistryconfig-eventrecordformat
	//
	EventRecordFormat *string `field:"optional" json:"eventRecordFormat" yaml:"eventRecordFormat"`
	// The URI for your schema registry.
	//
	// The correct URI format depends on the type of schema registry you're using.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemaregistryconfig.html#cfn-lambda-eventsourcemapping-schemaregistryconfig-schemaregistryuri
	//
	SchemaRegistryUri *string `field:"optional" json:"schemaRegistryUri" yaml:"schemaRegistryUri"`
	// An array of schema validation configuration objects, which tell Lambda the message attributes you want to validate and filter using your schema registry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-eventsourcemapping-schemaregistryconfig.html#cfn-lambda-eventsourcemapping-schemaregistryconfig-schemavalidationconfigs
	//
	SchemaValidationConfigs interface{} `field:"optional" json:"schemaValidationConfigs" yaml:"schemaValidationConfigs"`
}


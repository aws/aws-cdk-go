package awsbedrock


// Contains details about the OpenAPI schema for the action group.
//
// For more information, see [Action group OpenAPI schemas](https://docs.aws.amazon.com//bedrock/latest/userguide/agents-api-schema.html) . You can either include the schema directly in the payload field or you can upload it to an S3 bucket and specify the S3 bucket location in the s3 field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPISchemaProperty := &APISchemaProperty{
//   	Payload: jsii.String("payload"),
//   	S3: &S3IdentifierProperty{
//   		S3BucketName: jsii.String("s3BucketName"),
//   		S3ObjectKey: jsii.String("s3ObjectKey"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-apischema.html
//
type CfnAgent_APISchemaProperty struct {
	// The JSON or YAML-formatted payload defining the OpenAPI schema for the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-apischema.html#cfn-bedrock-agent-apischema-payload
	//
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Contains details about the S3 object containing the OpenAPI schema for the action group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-apischema.html#cfn-bedrock-agent-apischema-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}


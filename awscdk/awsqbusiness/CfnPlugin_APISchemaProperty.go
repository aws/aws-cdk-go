package awsqbusiness


// Contains details about the OpenAPI schema for a custom plugin.
//
// For more information, see [custom plugin OpenAPI schemas](https://docs.aws.amazon.com/amazonq/latest/qbusiness-ug/custom-plugin.html#plugins-api-schema) . You can either include the schema directly in the payload field or you can upload it to an S3 bucket and specify the S3 bucket location in the `s3` field.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aPISchemaProperty := &APISchemaProperty{
//   	Payload: jsii.String("payload"),
//   	S3: &S3Property{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-apischema.html
//
type CfnPlugin_APISchemaProperty struct {
	// The JSON or YAML-formatted payload defining the OpenAPI schema for a custom plugin.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-apischema.html#cfn-qbusiness-plugin-apischema-payload
	//
	Payload *string `field:"optional" json:"payload" yaml:"payload"`
	// Contains details about the S3 object containing the OpenAPI schema for a custom plugin.
	//
	// The schema could be in either JSON or YAML format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-plugin-apischema.html#cfn-qbusiness-plugin-apischema-s3
	//
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}


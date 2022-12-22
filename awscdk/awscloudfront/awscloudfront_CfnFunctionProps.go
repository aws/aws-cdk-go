package awscloudfront


// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionProps := &cfnFunctionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	autoPublish: jsii.Boolean(false),
//   	functionCode: jsii.String("functionCode"),
//   	functionConfig: &functionConfigProperty{
//   		comment: jsii.String("comment"),
//   		runtime: jsii.String("runtime"),
//   	},
//   	functionMetadata: &functionMetadataProperty{
//   		functionArn: jsii.String("functionArn"),
//   	},
//   }
//
type CfnFunctionProps struct {
	// A name to identify the function.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A flag that determines whether to automatically publish the function to the `LIVE` stage when it’s created.
	//
	// To automatically publish to the `LIVE` stage, set this property to `true` .
	AutoPublish interface{} `field:"optional" json:"autoPublish" yaml:"autoPublish"`
	// The function code.
	//
	// For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide* .
	FunctionCode *string `field:"optional" json:"functionCode" yaml:"functionCode"`
	// Contains configuration information about a CloudFront function.
	FunctionConfig interface{} `field:"optional" json:"functionConfig" yaml:"functionConfig"`
	// Contains metadata about a CloudFront function.
	FunctionMetadata interface{} `field:"optional" json:"functionMetadata" yaml:"functionMetadata"`
}


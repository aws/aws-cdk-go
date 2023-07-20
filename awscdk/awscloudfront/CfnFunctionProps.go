package awscloudfront


// Properties for defining a `CfnFunction`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnFunctionProps := &CfnFunctionProps{
//   	FunctionCode: jsii.String("functionCode"),
//   	FunctionConfig: &FunctionConfigProperty{
//   		Comment: jsii.String("comment"),
//   		Runtime: jsii.String("runtime"),
//   	},
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	AutoPublish: jsii.Boolean(false),
//   	FunctionMetadata: &FunctionMetadataProperty{
//   		FunctionArn: jsii.String("functionArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html
//
type CfnFunctionProps struct {
	// The function code.
	//
	// For more information about writing a CloudFront function, see [Writing function code for CloudFront Functions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/writing-function-code.html) in the *Amazon CloudFront Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html#cfn-cloudfront-function-functioncode
	//
	FunctionCode *string `field:"required" json:"functionCode" yaml:"functionCode"`
	// Contains configuration information about a CloudFront function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html#cfn-cloudfront-function-functionconfig
	//
	FunctionConfig interface{} `field:"required" json:"functionConfig" yaml:"functionConfig"`
	// A name to identify the function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html#cfn-cloudfront-function-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A flag that determines whether to automatically publish the function to the `LIVE` stage when it’s created.
	//
	// To automatically publish to the `LIVE` stage, set this property to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html#cfn-cloudfront-function-autopublish
	//
	AutoPublish interface{} `field:"optional" json:"autoPublish" yaml:"autoPublish"`
	// Contains metadata about a CloudFront function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudfront-function.html#cfn-cloudfront-function-functionmetadata
	//
	FunctionMetadata interface{} `field:"optional" json:"functionMetadata" yaml:"functionMetadata"`
}


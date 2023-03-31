package awsappsync


// Use the `PipelineConfig` property type to specify `PipelineConfig` for an AWS AppSync resolver.
//
// `PipelineConfig` is a property of the [AWS::AppSync::Resolver](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appsync-resolver.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pipelineConfigProperty := &pipelineConfigProperty{
//   	functions: []*string{
//   		jsii.String("functions"),
//   	},
//   }
//
type CfnResolver_PipelineConfigProperty struct {
	// A list of `Function` objects.
	Functions *[]*string `field:"optional" json:"functions" yaml:"functions"`
}


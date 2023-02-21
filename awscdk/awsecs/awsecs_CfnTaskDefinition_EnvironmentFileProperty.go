package awsecs


// A list of files containing the environment variables to pass to a container.
//
// You can specify up to ten environment files. The file must have a `.env` file extension. Each line in an environment file should contain an environment variable in `VARIABLE=VALUE` format. Lines beginning with `#` are treated as comments and are ignored. For more information about the environment variable file syntax, see [Declare default environment variables in file](https://docs.aws.amazon.com/https://docs.docker.com/compose/env-file/) .
//
// If there are environment variables specified using the `environment` parameter in a container definition, they take precedence over the variables contained within an environment file. If multiple environment files are specified that contain the same variable, they're processed from the top down. We recommend that you use unique variable names. For more information, see [Specifying environment variables](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/taskdef-envfiles.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// This parameter is only supported for tasks hosted on Fargate using the following platform versions:
//
// - Linux platform version `1.4.0` or later.
// - Windows platform version `1.0.0` or later.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   environmentFileProperty := &environmentFileProperty{
//   	type: jsii.String("type"),
//   	value: jsii.String("value"),
//   }
//
type CfnTaskDefinition_EnvironmentFileProperty struct {
	// The file type to use.
	//
	// The only supported value is `s3` .
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The Amazon Resource Name (ARN) of the Amazon S3 object containing the environment variable file.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


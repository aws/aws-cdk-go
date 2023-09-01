package awss3deployment

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Example:
//   import iam "github.com/aws/aws-cdk-go/awscdk"
//   import lambda "github.com/aws/aws-cdk-go/awscdk"
//
//   var myLambdaFunction function
//   var destinationBucket bucket
//   var role role
//
//
//   s3deploy.NewDeployTimeSubstitutedFile(this, jsii.String("MyFile"), &DeployTimeSubstitutedFileProps{
//   	Source: jsii.String("my-file.yaml"),
//   	DestinationBucket: destinationBucket,
//   	Substitutions: map[string]*string{
//   		"variableName": myLambdaFunction.functionName,
//   	},
//   	Role: role,
//   })
//
type DeployTimeSubstitutedFileProps struct {
	// The S3 bucket to sync the contents of the zip file to.
	DestinationBucket awss3.IBucket `field:"required" json:"destinationBucket" yaml:"destinationBucket"`
	// Path to the user's local file.
	Source *string `field:"required" json:"source" yaml:"source"`
	// User-defined substitutions to make in the file.
	//
	// Placeholders in the user's local file must be specified with double curly
	// brackets and spaces. For example, if you use the key 'xxxx' in the file,
	// it must be written as: {{ xxxx }} to be recognized by the construct as a
	// substitution.
	Substitutions *map[string]*string `field:"required" json:"substitutions" yaml:"substitutions"`
	// Execution role associated with this function.
	// Default: - A role is automatically created.
	//
	Role awsiam.IRole `field:"optional" json:"role" yaml:"role"`
}


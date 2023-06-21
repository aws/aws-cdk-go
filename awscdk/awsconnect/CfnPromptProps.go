package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnPrompt`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnPromptProps := &CfnPromptProps{
//   	InstanceArn: jsii.String("instanceArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	S3Uri: jsii.String("s3Uri"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnPromptProps struct {
	// The identifier of the Amazon Connect instance.
	InstanceArn *string `field:"required" json:"instanceArn" yaml:"instanceArn"`
	// The name of the prompt.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the prompt.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The URI for the S3 bucket where the prompt is stored.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsevidently

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnProject`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnProjectProps := &cfnProjectProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	appConfigResource: &appConfigResourceObjectProperty{
//   		applicationId: jsii.String("applicationId"),
//   		environmentId: jsii.String("environmentId"),
//   	},
//   	dataDelivery: &dataDeliveryObjectProperty{
//   		logGroup: jsii.String("logGroup"),
//   		s3: &s3DestinationProperty{
//   			bucketName: jsii.String("bucketName"),
//
//   			// the properties below are optional
//   			prefix: jsii.String("prefix"),
//   		},
//   	},
//   	description: jsii.String("description"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnProjectProps struct {
	// The name for the project.
	//
	// It can include up to 127 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Evidently::Project.AppConfigResource`.
	AppConfigResource interface{} `field:"optional" json:"appConfigResource" yaml:"appConfigResource"`
	// A structure that contains information about where Evidently is to store evaluation events for longer term storage, if you choose to do so.
	//
	// If you choose not to store these events, Evidently deletes them after using them to produce metrics and other experiment results that you can view.
	//
	// You can't specify both `CloudWatchLogs` and `S3Destination` in the same operation.
	DataDelivery interface{} `field:"optional" json:"dataDelivery" yaml:"dataDelivery"`
	// An optional description of the project.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Assigns one or more tags (key-value pairs) to the project.
	//
	// Tags can help you organize and categorize your resources. You can also use them to scope user permissions by granting a user permission to access or change only resources with certain tag values.
	//
	// Tags don't have any semantic meaning to AWS and are interpreted strictly as strings of characters.
	//
	// You can associate as many as 50 tags with a project.
	//
	// For more information, see [Tagging AWS resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) .
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnScript`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnScriptProps := &cfnScriptProps{
//   	storageLocation: &s3LocationProperty{
//   		bucket: jsii.String("bucket"),
//   		key: jsii.String("key"),
//   		roleArn: jsii.String("roleArn"),
//
//   		// the properties below are optional
//   		objectVersion: jsii.String("objectVersion"),
//   	},
//
//   	// the properties below are optional
//   	name: jsii.String("name"),
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	version: jsii.String("version"),
//   }
//
type CfnScriptProps struct {
	// The location of the Amazon S3 bucket where a zipped file containing your Realtime scripts is stored.
	//
	// The storage location must specify the Amazon S3 bucket name, the zip file name (the "key"), and a role ARN that allows Amazon GameLift to access the Amazon S3 storage location. The S3 bucket must be in the same Region where you want to create a new script. By default, Amazon GameLift uploads the latest version of the zip file; if you have S3 object versioning turned on, you can use the `ObjectVersion` parameter to specify an earlier version.
	StorageLocation interface{} `field:"required" json:"storageLocation" yaml:"storageLocation"`
	// A descriptive label that is associated with a script.
	//
	// Script names do not need to be unique.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// A list of labels to assign to the new script resource.
	//
	// Tags are developer-defined key-value pairs. Tagging AWS resources are useful for resource management, access management and cost allocation. For more information, see [Tagging AWS Resources](https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the *AWS General Reference* . Once the resource is created, you can use TagResource, UntagResource, and ListTagsForResource to add, remove, and view tags. The maximum tag limit may be lower than stated. See the AWS General Reference for actual tagging limits.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The version that is associated with a build or script.
	//
	// Version strings do not need to be unique.
	Version *string `field:"optional" json:"version" yaml:"version"`
}


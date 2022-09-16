package awsappmesh


// An object that represents an access log file.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileAccessLogProperty := &fileAccessLogProperty{
//   	path: jsii.String("path"),
//
//   	// the properties below are optional
//   	format: &loggingFormatProperty{
//   		json: []interface{}{
//   			&jsonFormatRefProperty{
//   				key: jsii.String("key"),
//   				value: jsii.String("value"),
//   			},
//   		},
//   		text: jsii.String("text"),
//   	},
//   }
//
type CfnVirtualNode_FileAccessLogProperty struct {
	// The file path to write access logs to.
	//
	// You can use `/dev/stdout` to send access logs to standard out and configure your Envoy container to use a log driver, such as `awslogs` , to export the access logs to a log storage service such as Amazon CloudWatch Logs. You can also specify a path in the Envoy container's file system to write the files to disk.
	//
	// > The Envoy process must have write permissions to the path that you specify here. Otherwise, Envoy fails to bootstrap properly.
	Path *string `field:"required" json:"path" yaml:"path"`
	// `CfnVirtualNode.FileAccessLogProperty.Format`.
	Format interface{} `field:"optional" json:"format" yaml:"format"`
}


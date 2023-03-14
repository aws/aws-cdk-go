package awsecs


// The FireLens configuration for the container.
//
// This is used to specify and configure a log router for container logs. For more information, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   firelensConfigurationProperty := &firelensConfigurationProperty{
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	type: jsii.String("type"),
//   }
//
type CfnTaskDefinition_FirelensConfigurationProperty struct {
	// The options to use when configuring the log router.
	//
	// This field is optional and can be used to add additional metadata, such as the task, task definition, cluster, and container instance details to the log event.
	//
	// If specified, valid option keys are:
	//
	// - `enable-ecs-log-metadata` , which can be `true` or `false`
	// - `config-file-type` , which can be `s3` or `file`
	// - `config-file-value` , which is either an S3 ARN or a file path.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The log router to use.
	//
	// The valid values are `fluentd` or `fluentbit` .
	Type *string `field:"optional" json:"type" yaml:"type"`
}


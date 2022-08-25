package awsbatch


// Log configuration options to send to a custom log driver for the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var options interface{}
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: options,
//   	secretOptions: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
type CfnJobDefinition_LogConfigurationProperty struct {
	// The log driver to use for the container.
	//
	// The valid values listed for this parameter are log drivers that the Amazon ECS container agent can communicate with by default.
	//
	// The supported log drivers are `awslogs` , `fluentd` , `gelf` , `json-file` , `journald` , `logentries` , `syslog` , and `splunk` .
	//
	// > Jobs that are running on Fargate resources are restricted to the `awslogs` and `splunk` log drivers.
	//
	// - **awslogs** - Specifies the Amazon CloudWatch Logs logging driver. For more information, see [Using the awslogs log driver](https://docs.aws.amazon.com/batch/latest/userguide/using_awslogs.html) in the *AWS Batch User Guide* and [Amazon CloudWatch Logs logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/awslogs/) in the Docker documentation.
	// - **fluentd** - Specifies the Fluentd logging driver. For more information, including usage and options, see [Fluentd logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/fluentd/) in the Docker documentation.
	// - **gelf** - Specifies the Graylog Extended Format (GELF) logging driver. For more information, including usage and options, see [Graylog Extended Format logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/gelf/) in the Docker documentation.
	// - **journald** - Specifies the journald logging driver. For more information, including usage and options, see [Journald logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/journald/) in the Docker documentation.
	// - **json-file** - Specifies the JSON file logging driver. For more information, including usage and options, see [JSON File logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/json-file/) in the Docker documentation.
	// - **splunk** - Specifies the Splunk logging driver. For more information, including usage and options, see [Splunk logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/splunk/) in the Docker documentation.
	// - **syslog** - Specifies the syslog logging driver. For more information, including usage and options, see [Syslog logging driver](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/syslog/) in the Docker documentation.
	//
	// > If you have a custom driver that's not listed earlier that you want to work with the Amazon ECS container agent, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you want to have included. However, Amazon Web Services doesn't currently support running modified copies of this software.
	//
	// This parameter requires version 1.18 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log into your container instance and run the following command: `sudo docker version | grep "Server API version"`
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	//
	// This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log into your container instance and run the following command: `sudo docker version | grep "Server API version"`
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	//
	// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data.html) in the *AWS Batch User Guide* .
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


package mixinsawsecs


// The `LogConfiguration` property specifies log configuration options to send to a custom log driver for the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   logConfigurationProperty := &LogConfigurationProperty{
//   	LogDriver: jsii.String("logDriver"),
//   	Options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	SecretOptions: []interface{}{
//   		&SecretProperty{
//   			Name: jsii.String("name"),
//   			ValueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html
//
type CfnTaskDefinitionPropsMixin_LogConfigurationProperty struct {
	// The log driver to use for the container.
	//
	// For tasks on AWS Fargate , the supported log drivers are `awslogs` , `splunk` , and `awsfirelens` .
	//
	// For tasks hosted on Amazon EC2 instances, the supported log drivers are `awslogs` , `fluentd` , `gelf` , `json-file` , `journald` , `syslog` , `splunk` , and `awsfirelens` .
	//
	// For more information about using the `awslogs` log driver, see [Send Amazon ECS logs to CloudWatch](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For more information about using the `awsfirelens` log driver, see [Send Amazon ECS logs to an AWS service or AWS Partner](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) .
	//
	// > If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-logdriver
	//
	LogDriver *string `field:"optional" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	//
	// The options you can specify depend on the log driver. Some of the options you can specify when you use the `awslogs` log driver to route logs to Amazon CloudWatch include the following:
	//
	// - **awslogs-create-group** - Required: No
	//
	// Specify whether you want the log group to be created automatically. If this option isn't specified, it defaults to `false` .
	//
	// > Your IAM policy must include the `logs:CreateLogGroup` permission before you attempt to use `awslogs-create-group` .
	// - **awslogs-region** - Required: Yes
	//
	// Specify the AWS Region that the `awslogs` log driver is to send your Docker logs to. You can choose to send all of your logs from clusters in different Regions to a single region in CloudWatch Logs. This is so that they're all visible in one location. Otherwise, you can separate them by Region for more granularity. Make sure that the specified log group exists in the Region that you specify with this option.
	// - **awslogs-group** - Required: Yes
	//
	// Make sure to specify a log group that the `awslogs` log driver sends its log streams to.
	// - **awslogs-stream-prefix** - Required: Yes, when using Fargate.Optional when using EC2.
	//
	// Use the `awslogs-stream-prefix` option to associate a log stream with the specified prefix, the container name, and the ID of the Amazon ECS task that the container belongs to. If you specify a prefix with this option, then the log stream takes the format `prefix-name/container-name/ecs-task-id` .
	//
	// If you don't specify a prefix with this option, then the log stream is named after the container ID that's assigned by the Docker daemon on the container instance. Because it's difficult to trace logs back to the container that sent them with just the Docker container ID (which is only available on the container instance), we recommend that you specify a prefix with this option.
	//
	// For Amazon ECS services, you can use the service name as the prefix. Doing so, you can trace log streams to the service that the container belongs to, the name of the container that sent them, and the ID of the task that the container belongs to.
	//
	// You must specify a stream-prefix for your logs to have your logs appear in the Log pane when using the Amazon ECS console.
	// - **awslogs-datetime-format** - Required: No
	//
	// This option defines a multiline start pattern in Python `strftime` format. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages.
	//
	// One example of a use case for using this format is for parsing output such as a stack dump, which might otherwise be logged in multiple entries. The correct pattern allows it to be captured in a single entry.
	//
	// For more information, see [awslogs-datetime-format](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/awslogs/#awslogs-datetime-format) .
	//
	// You cannot configure both the `awslogs-datetime-format` and `awslogs-multiline-pattern` options.
	//
	// > Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance.
	// - **awslogs-multiline-pattern** - Required: No
	//
	// This option defines a multiline start pattern that uses a regular expression. A log message consists of a line that matches the pattern and any following lines that don’t match the pattern. The matched line is the delimiter between log messages.
	//
	// For more information, see [awslogs-multiline-pattern](https://docs.aws.amazon.com/https://docs.docker.com/config/containers/logging/awslogs/#awslogs-multiline-pattern) .
	//
	// This option is ignored if `awslogs-datetime-format` is also configured.
	//
	// You cannot configure both the `awslogs-datetime-format` and `awslogs-multiline-pattern` options.
	//
	// > Multiline logging performs regular expression parsing and matching of all log messages. This might have a negative impact on logging performance.
	//
	// The following options apply to all supported log drivers.
	//
	// - **mode** - Required: No
	//
	// Valid values: `non-blocking` | `blocking`
	//
	// This option defines the delivery mode of log messages from the container to the log driver specified using `logDriver` . The delivery mode you choose affects application availability when the flow of logs from container is interrupted.
	//
	// If you use the `blocking` mode and the flow of logs is interrupted, calls from container code to write to the `stdout` and `stderr` streams will block. The logging thread of the application will block as a result. This may cause the application to become unresponsive and lead to container healthcheck failure.
	//
	// If you use the `non-blocking` mode, the container's logs are instead stored in an in-memory intermediate buffer configured with the `max-buffer-size` option. This prevents the application from becoming unresponsive when logs cannot be sent. We recommend using this mode if you want to ensure service availability and are okay with some log loss. For more information, see [Preventing log loss with non-blocking mode in the `awslogs` container log driver](https://docs.aws.amazon.com/containers/preventing-log-loss-with-non-blocking-mode-in-the-awslogs-container-log-driver/) .
	//
	// You can set a default `mode` for all containers in a specific AWS Region by using the `defaultLogDriverMode` account setting. If you don't specify the `mode` option or configure the account setting, Amazon ECS will default to the `non-blocking` mode. For more information about the account setting, see [Default log driver mode](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#default-log-driver-mode) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > On June 25, 2025, Amazon ECS changed the default log driver mode from `blocking` to `non-blocking` to prioritize task availability over logging. To continue using the `blocking` mode after this change, do one of the following:
	// >
	// > - Set the `mode` option in your container definition's `logConfiguration` as `blocking` .
	// > - Set the `defaultLogDriverMode` account setting to `blocking` .
	// - **max-buffer-size** - Required: No
	//
	// Default value: `10m`
	//
	// When `non-blocking` mode is used, the `max-buffer-size` log option controls the size of the buffer that's used for intermediate message storage. Make sure to specify an adequate buffer size based on your application. When the buffer fills up, further logs cannot be stored. Logs that cannot be stored are lost.
	//
	// To route logs using the `splunk` log router, you need to specify a `splunk-token` and a `splunk-url` .
	//
	// When you use the `awsfirelens` log router to route logs to an AWS Service or AWS Partner Network destination for log storage and analytics, you can set the `log-driver-buffer-limit` option to limit the number of events that are buffered in memory, before being sent to the log router container. It can help to resolve potential log loss issue because high throughput might result in memory running out for the buffer inside of Docker.
	//
	// Other options you can specify when using `awsfirelens` to route logs depend on the destination. When you export logs to Amazon Data Firehose, you can specify the AWS Region with `region` and a name for the log stream with `delivery_stream` .
	//
	// When you export logs to Amazon Kinesis Data Streams, you can specify an AWS Region with `region` and a data stream name with `stream` .
	//
	// When you export logs to Amazon OpenSearch Service, you can specify options like `Name` , `Host` (OpenSearch Service endpoint without protocol), `Port` , `Index` , `Type` , `Aws_auth` , `Aws_region` , `Suppress_Type_Name` , and `tls` . For more information, see [Under the hood: FireLens for Amazon ECS Tasks](https://docs.aws.amazon.com/containers/under-the-hood-firelens-for-amazon-ecs-tasks/) .
	//
	// When you export logs to Amazon S3, you can specify the bucket using the `bucket` option. You can also specify `region` , `total_file_size` , `upload_timeout` , and `use_put_object` as options.
	//
	// This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-options
	//
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	//
	// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-logconfiguration.html#cfn-ecs-taskdefinition-logconfiguration-secretoptions
	//
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


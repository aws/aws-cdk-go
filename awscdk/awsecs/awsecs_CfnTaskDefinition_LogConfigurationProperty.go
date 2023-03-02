package awsecs


// The `LogConfiguration` property specifies log configuration options to send to a custom log driver for the container.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logConfigurationProperty := &logConfigurationProperty{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: []interface{}{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
type CfnTaskDefinition_LogConfigurationProperty struct {
	// The log driver to use for the container.
	//
	// For tasks on AWS Fargate , the supported log drivers are `awslogs` , `splunk` , and `awsfirelens` .
	//
	// For tasks hosted on Amazon EC2 instances, the supported log drivers are `awslogs` , `fluentd` , `gelf` , `json-file` , `journald` , `logentries` , `syslog` , `splunk` , and `awsfirelens` .
	//
	// For more information about using the `awslogs` log driver, see [Using the awslogs log driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// For more information about using the `awsfirelens` log driver, see [Custom log routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html) in the *Amazon Elastic Container Service Developer Guide* .
	//
	// > If you have a custom driver that isn't listed, you can fork the Amazon ECS container agent project that's [available on GitHub](https://docs.aws.amazon.com/https://github.com/aws/amazon-ecs-agent) and customize it to work with that driver. We encourage you to submit pull requests for changes that you would like to have included. However, we don't currently provide support for running modified copies of this software.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	//
	// This parameter requires version 1.19 of the Docker Remote API or greater on your container instance. To check the Docker Remote API version on your container instance, log in to your container instance and run the following command: `sudo docker version --format '{{.Server.APIVersion}}'`
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	//
	// For more information, see [Specifying sensitive data](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/specifying-sensitive-data.html) in the *Amazon Elastic Container Service Developer Guide* .
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


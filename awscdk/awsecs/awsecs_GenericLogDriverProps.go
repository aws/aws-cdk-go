package awsecs


// The configuration to use when creating a log driver.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.NewGenericLogDriver(&genericLogDriverProps{
//   		logDriver: jsii.String("fluentd"),
//   		options: map[string]*string{
//   			"tag": jsii.String("example-tag"),
//   		},
//   	}),
//   })
//
type GenericLogDriverProps struct {
	// The log driver to use for the container.
	//
	// The valid values listed for this parameter are log drivers
	// that the Amazon ECS container agent can communicate with by default.
	//
	// For tasks using the Fargate launch type, the supported log drivers are awslogs and splunk.
	// For tasks using the EC2 launch type, the supported log drivers are awslogs, syslog, gelf, fluentd, splunk, journald, and json-file.
	//
	// For more information about using the awslogs log driver, see
	// [Using the awslogs Log Driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *map[string]Secret `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


package awsecs


// Specifies the firelens log driver configuration options.
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_Firelens(&FireLensLogDriverProps{
//   		Options: map[string]*string{
//   			"Name": jsii.String("firehose"),
//   			"region": jsii.String("us-west-2"),
//   			"delivery_stream": jsii.String("my-stream"),
//   		},
//   	}),
//   })
//
type FireLensLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	// Default: - No env.
	//
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	// Default: - No envRegex.
	//
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	// Default: - No labels.
	//
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	// Default: - The first 12 characters of the container ID.
	//
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// The configuration options to send to the log driver.
	// Default: - the log driver options.
	//
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	// Default: - No secret options provided.
	//
	SecretOptions *map[string]Secret `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


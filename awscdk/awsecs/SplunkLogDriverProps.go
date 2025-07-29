package awsecs


// Specifies the splunk log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/splunk/)
//
// Example:
//   var secret secret
//
//
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_Splunk(&SplunkLogDriverProps{
//   		SecretToken: secret,
//   		Url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type SplunkLogDriverProps struct {
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
	// Splunk HTTP Event Collector token (Secret).
	//
	// The splunk-token is added to the SecretOptions property of the Log Driver Configuration. So the secret value will not be
	// resolved or viewable as plain text.
	SecretToken Secret `field:"required" json:"secretToken" yaml:"secretToken"`
	// Path to your Splunk Enterprise, self-service Splunk Cloud instance, or Splunk Cloud managed cluster (including port and scheme used by HTTP Event Collector) in one of the following formats: https://your_splunk_instance:8088 or https://input-prd-p-XXXXXXX.cloud.splunk.com:8088 or https://http-inputs-XXXXXXXX.splunkcloud.com.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Name to use for validating server certificate.
	// Default: - The hostname of the splunk-url.
	//
	CaName *string `field:"optional" json:"caName" yaml:"caName"`
	// Path to root certificate.
	// Default: - caPath not set.
	//
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// Message format.
	//
	// Can be inline, json or raw.
	// Default: - inline.
	//
	Format SplunkLogFormat `field:"optional" json:"format" yaml:"format"`
	// Enable/disable gzip compression to send events to Splunk Enterprise or Splunk Cloud instance.
	// Default: - false.
	//
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// Set compression level for gzip.
	//
	// Valid values are -1 (default), 0 (no compression),
	// 1 (best speed) ... 9 (best compression).
	// Default: - -1 (Default Compression).
	//
	GzipLevel *float64 `field:"optional" json:"gzipLevel" yaml:"gzipLevel"`
	// Event index.
	// Default: - index not set.
	//
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Ignore server certificate validation.
	// Default: - insecureSkipVerify not set.
	//
	InsecureSkipVerify *string `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// Event source.
	// Default: - source not set.
	//
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Event source type.
	// Default: - sourceType not set.
	//
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Verify on start, that docker can connect to Splunk server.
	// Default: - true.
	//
	VerifyConnection *bool `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}


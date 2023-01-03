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
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		secretToken: secret,
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
type SplunkLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Splunk HTTP Event Collector token (Secret).
	//
	// The splunk-token is added to the SecretOptions property of the Log Driver Configuration. So the secret value will not be
	// resolved or viewable as plain text.
	SecretToken Secret `field:"required" json:"secretToken" yaml:"secretToken"`
	// Path to your Splunk Enterprise, self-service Splunk Cloud instance, or Splunk Cloud managed cluster (including port and scheme used by HTTP Event Collector) in one of the following formats: https://your_splunk_instance:8088 or https://input-prd-p-XXXXXXX.cloud.splunk.com:8088 or https://http-inputs-XXXXXXXX.splunkcloud.com.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Name to use for validating server certificate.
	CaName *string `field:"optional" json:"caName" yaml:"caName"`
	// Path to root certificate.
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// Message format.
	//
	// Can be inline, json or raw.
	Format SplunkLogFormat `field:"optional" json:"format" yaml:"format"`
	// Enable/disable gzip compression to send events to Splunk Enterprise or Splunk Cloud instance.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// Set compression level for gzip.
	//
	// Valid values are -1 (default), 0 (no compression),
	// 1 (best speed) ... 9 (best compression).
	GzipLevel *float64 `field:"optional" json:"gzipLevel" yaml:"gzipLevel"`
	// Event index.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Ignore server certificate validation.
	InsecureSkipVerify *string `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// Event source.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Event source type.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Verify on start, that docker can connect to Splunk server.
	VerifyConnection *bool `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}


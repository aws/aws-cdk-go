package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Specifies the splunk log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/splunk/)
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.addContainer(jsii.String("TheContainer"), &containerDefinitionOptions{
//   	image: ecs.containerImage.fromRegistry(jsii.String("example-image")),
//   	memoryLimitMiB: jsii.Number(256),
//   	logging: ecs.logDrivers.splunk(&splunkLogDriverProps{
//   		token: awscdk.SecretValue.secretsManager(jsii.String("my-splunk-token")),
//   		url: jsii.String("my-splunk-url"),
//   	}),
//   })
//
// Experimental.
type SplunkLogDriverProps struct {
	// The env option takes an array of keys.
	//
	// If there is collision between
	// label and env keys, the value of the env takes precedence. Adds additional fields
	// to the extra attributes of a logging message.
	// Experimental.
	Env *[]*string `field:"optional" json:"env" yaml:"env"`
	// The env-regex option is similar to and compatible with env.
	//
	// Its value is a regular
	// expression to match logging-related environment variables. It is used for advanced
	// log tag options.
	// Experimental.
	EnvRegex *string `field:"optional" json:"envRegex" yaml:"envRegex"`
	// The labels option takes an array of keys.
	//
	// If there is collision
	// between label and env keys, the value of the env takes precedence. Adds additional
	// fields to the extra attributes of a logging message.
	// Experimental.
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// By default, Docker uses the first 12 characters of the container ID to tag log messages.
	//
	// Refer to the log tag option documentation for customizing the
	// log tag format.
	// Experimental.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Path to your Splunk Enterprise, self-service Splunk Cloud instance, or Splunk Cloud managed cluster (including port and scheme used by HTTP Event Collector) in one of the following formats: https://your_splunk_instance:8088 or https://input-prd-p-XXXXXXX.cloud.splunk.com:8088 or https://http-inputs-XXXXXXXX.splunkcloud.com.
	// Experimental.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Name to use for validating server certificate.
	// Experimental.
	CaName *string `field:"optional" json:"caName" yaml:"caName"`
	// Path to root certificate.
	// Experimental.
	CaPath *string `field:"optional" json:"caPath" yaml:"caPath"`
	// Message format.
	//
	// Can be inline, json or raw.
	// Experimental.
	Format SplunkLogFormat `field:"optional" json:"format" yaml:"format"`
	// Enable/disable gzip compression to send events to Splunk Enterprise or Splunk Cloud instance.
	// Experimental.
	Gzip *bool `field:"optional" json:"gzip" yaml:"gzip"`
	// Set compression level for gzip.
	//
	// Valid values are -1 (default), 0 (no compression),
	// 1 (best speed) ... 9 (best compression).
	// Experimental.
	GzipLevel *float64 `field:"optional" json:"gzipLevel" yaml:"gzipLevel"`
	// Event index.
	// Experimental.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Ignore server certificate validation.
	// Experimental.
	InsecureSkipVerify *string `field:"optional" json:"insecureSkipVerify" yaml:"insecureSkipVerify"`
	// Splunk HTTP Event Collector token (Secret).
	//
	// The splunk-token is added to the SecretOptions property of the Log Driver Configuration. So the secret value will not be
	// resolved or viewable as plain text.
	//
	// Please provide at least one of `token` or `secretToken`.
	// Experimental.
	SecretToken Secret `field:"optional" json:"secretToken" yaml:"secretToken"`
	// Event source.
	// Experimental.
	Source *string `field:"optional" json:"source" yaml:"source"`
	// Event source type.
	// Experimental.
	SourceType *string `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Splunk HTTP Event Collector token.
	//
	// The splunk-token is added to the Options property of the Log Driver Configuration. So the secret value will be resolved and
	// viewable in plain text in the console.
	//
	// Please provide at least one of `token` or `secretToken`.
	// Deprecated: Use {@link SplunkLogDriverProps.secretToken} instead.
	Token awscdk.SecretValue `field:"optional" json:"token" yaml:"token"`
	// Verify on start, that docker can connect to Splunk server.
	// Experimental.
	VerifyConnection *bool `field:"optional" json:"verifyConnection" yaml:"verifyConnection"`
}


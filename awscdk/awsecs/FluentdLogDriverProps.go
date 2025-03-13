package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies the fluentd log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/fluentd/)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fluentdLogDriverProps := &FluentdLogDriverProps{
//   	Address: jsii.String("address"),
//   	AsyncConnect: jsii.Boolean(false),
//   	BufferLimit: jsii.Number(123),
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	MaxRetries: jsii.Number(123),
//   	RetryWait: cdk.Duration_Minutes(jsii.Number(30)),
//   	SubSecondPrecision: jsii.Boolean(false),
//   	Tag: jsii.String("tag"),
//   }
//
type FluentdLogDriverProps struct {
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
	// By default, the logging driver connects to localhost:24224.
	//
	// Supply the
	// address option to connect to a different address. tcp(default) and unix
	// sockets are supported.
	// Default: - address not set.
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docker connects to Fluentd in the background.
	//
	// Messages are buffered until
	// the connection is established.
	// Default: - false.
	//
	AsyncConnect *bool `field:"optional" json:"asyncConnect" yaml:"asyncConnect"`
	// The amount of data to buffer before flushing to disk.
	// Default: - The amount of RAM available to the container.
	//
	BufferLimit *float64 `field:"optional" json:"bufferLimit" yaml:"bufferLimit"`
	// The maximum number of retries.
	// Default: - 4294967295 (2**32 - 1).
	//
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// How long to wait between retries.
	// Default: - 1 second.
	//
	RetryWait awscdk.Duration `field:"optional" json:"retryWait" yaml:"retryWait"`
	// Generates event logs in nanosecond resolution.
	// Default: - false.
	//
	SubSecondPrecision *bool `field:"optional" json:"subSecondPrecision" yaml:"subSecondPrecision"`
}


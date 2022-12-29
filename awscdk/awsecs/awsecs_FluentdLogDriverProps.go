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
//   fluentdLogDriverProps := &fluentdLogDriverProps{
//   	address: jsii.String("address"),
//   	asyncConnect: jsii.Boolean(false),
//   	bufferLimit: jsii.Number(123),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxRetries: jsii.Number(123),
//   	retryWait: cdk.duration.minutes(jsii.Number(30)),
//   	subSecondPrecision: jsii.Boolean(false),
//   	tag: jsii.String("tag"),
//   }
//
type FluentdLogDriverProps struct {
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
	// By default, the logging driver connects to localhost:24224.
	//
	// Supply the
	// address option to connect to a different address. tcp(default) and unix
	// sockets are supported.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docker connects to Fluentd in the background.
	//
	// Messages are buffered until
	// the connection is established.
	AsyncConnect *bool `field:"optional" json:"asyncConnect" yaml:"asyncConnect"`
	// The amount of data to buffer before flushing to disk.
	BufferLimit *float64 `field:"optional" json:"bufferLimit" yaml:"bufferLimit"`
	// The maximum number of retries.
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// How long to wait between retries.
	RetryWait awscdk.Duration `field:"optional" json:"retryWait" yaml:"retryWait"`
	// Generates event logs in nanosecond resolution.
	SubSecondPrecision *bool `field:"optional" json:"subSecondPrecision" yaml:"subSecondPrecision"`
}


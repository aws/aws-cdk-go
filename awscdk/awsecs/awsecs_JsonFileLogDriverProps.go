package awsecs


// Specifies the json-file log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/json-file/)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jsonFileLogDriverProps := &jsonFileLogDriverProps{
//   	compress: jsii.Boolean(false),
//   	env: []*string{
//   		jsii.String("env"),
//   	},
//   	envRegex: jsii.String("envRegex"),
//   	labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	maxFile: jsii.Number(123),
//   	maxSize: jsii.String("maxSize"),
//   	tag: jsii.String("tag"),
//   }
//
type JsonFileLogDriverProps struct {
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
	// Toggles compression for rotated logs.
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The maximum number of log files that can be present.
	//
	// If rolling the logs creates
	// excess files, the oldest file is removed. Only effective when max-size is also set.
	// A positive integer.
	MaxFile *float64 `field:"optional" json:"maxFile" yaml:"maxFile"`
	// The maximum size of the log before it is rolled.
	//
	// A positive integer plus a modifier
	// representing the unit of measure (k, m, or g).
	MaxSize *string `field:"optional" json:"maxSize" yaml:"maxSize"`
}


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
//   jsonFileLogDriverProps := &JsonFileLogDriverProps{
//   	Compress: jsii.Boolean(false),
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	MaxFile: jsii.Number(123),
//   	MaxSize: jsii.String("maxSize"),
//   	Tag: jsii.String("tag"),
//   }
//
type JsonFileLogDriverProps struct {
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
	// Toggles compression for rotated logs.
	// Default: - false.
	//
	Compress *bool `field:"optional" json:"compress" yaml:"compress"`
	// The maximum number of log files that can be present.
	//
	// If rolling the logs creates
	// excess files, the oldest file is removed. Only effective when max-size is also set.
	// A positive integer.
	// Default: - 1.
	//
	MaxFile *float64 `field:"optional" json:"maxFile" yaml:"maxFile"`
	// The maximum size of the log before it is rolled.
	//
	// A positive integer plus a modifier
	// representing the unit of measure (k, m, or g).
	// Default: - -1 (unlimited).
	//
	MaxSize *string `field:"optional" json:"maxSize" yaml:"maxSize"`
}


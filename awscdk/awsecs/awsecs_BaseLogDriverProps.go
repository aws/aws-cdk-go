package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   baseLogDriverProps := &BaseLogDriverProps{
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	Tag: jsii.String("tag"),
//   }
//
// Experimental.
type BaseLogDriverProps struct {
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
}


package awsecs


// Specifies the syslog log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/syslog/)
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   syslogLogDriverProps := &SyslogLogDriverProps{
//   	Address: jsii.String("address"),
//   	Env: []*string{
//   		jsii.String("env"),
//   	},
//   	EnvRegex: jsii.String("envRegex"),
//   	Facility: jsii.String("facility"),
//   	Format: jsii.String("format"),
//   	Labels: []*string{
//   		jsii.String("labels"),
//   	},
//   	Tag: jsii.String("tag"),
//   	TlsCaCert: jsii.String("tlsCaCert"),
//   	TlsCert: jsii.String("tlsCert"),
//   	TlsKey: jsii.String("tlsKey"),
//   	TlsSkipVerify: jsii.Boolean(false),
//   }
//
type SyslogLogDriverProps struct {
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
	// The address of an external syslog server.
	//
	// The URI specifier may be
	// [tcp|udp|tcp+tls]://host:port, unix://path, or unixgram://path.
	// Default: - If the transport is tcp, udp, or tcp+tls, the default port is 514.
	//
	Address *string `field:"optional" json:"address" yaml:"address"`
	// The syslog facility to use.
	//
	// Can be the number or name for any valid
	// syslog facility. See the syslog documentation:
	// https://tools.ietf.org/html/rfc5424#section-6.2.1.
	// Default: - facility not set.
	//
	Facility *string `field:"optional" json:"facility" yaml:"facility"`
	// The syslog message format to use.
	//
	// If not specified the local UNIX syslog
	// format is used, without a specified hostname. Specify rfc3164 for the RFC-3164
	// compatible format, rfc5424 for RFC-5424 compatible format, or rfc5424micro
	// for RFC-5424 compatible format with microsecond timestamp resolution.
	// Default: - format not set.
	//
	Format *string `field:"optional" json:"format" yaml:"format"`
	// The absolute path to the trust certificates signed by the CA.
	//
	// Ignored
	// if the address protocol is not tcp+tls.
	// Default: - tlsCaCert not set.
	//
	TlsCaCert *string `field:"optional" json:"tlsCaCert" yaml:"tlsCaCert"`
	// The absolute path to the TLS certificate file.
	//
	// Ignored if the address
	// protocol is not tcp+tls.
	// Default: - tlsCert not set.
	//
	TlsCert *string `field:"optional" json:"tlsCert" yaml:"tlsCert"`
	// The absolute path to the TLS key file.
	//
	// Ignored if the address protocol
	// is not tcp+tls.
	// Default: - tlsKey not set.
	//
	TlsKey *string `field:"optional" json:"tlsKey" yaml:"tlsKey"`
	// If set to true, TLS verification is skipped when connecting to the syslog daemon.
	//
	// Ignored if the address protocol is not tcp+tls.
	// Default: - false.
	//
	TlsSkipVerify *bool `field:"optional" json:"tlsSkipVerify" yaml:"tlsSkipVerify"`
}


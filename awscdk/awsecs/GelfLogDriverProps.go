package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Specifies the journald log driver configuration options.
//
// [Source](https://docs.docker.com/config/containers/logging/gelf/)
//
// Example:
//   // Create a Task Definition for the container to start
//   taskDefinition := ecs.NewEc2TaskDefinition(this, jsii.String("TaskDef"))
//   taskDefinition.AddContainer(jsii.String("TheContainer"), &ContainerDefinitionOptions{
//   	Image: ecs.ContainerImage_FromRegistry(jsii.String("example-image")),
//   	MemoryLimitMiB: jsii.Number(256),
//   	Logging: ecs.LogDrivers_Gelf(&GelfLogDriverProps{
//   		Address: jsii.String("my-gelf-address"),
//   	}),
//   })
//
type GelfLogDriverProps struct {
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
	// The address of the GELF server.
	//
	// tcp and udp are the only supported URI
	// specifier and you must specify the port.
	Address *string `field:"required" json:"address" yaml:"address"`
	// UDP Only The level of compression when gzip or zlib is the gelf-compression-type.
	//
	// An integer in the range of -1 to 9 (BestCompression). Higher levels provide more
	// compression at lower speed. Either -1 or 0 disables compression.
	// Default: - 1.
	//
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// UDP Only The type of compression the GELF driver uses to compress each log message.
	//
	// Allowed values are gzip, zlib and none.
	// Default: - gzip.
	//
	CompressionType GelfCompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// TCP Only The maximum number of reconnection attempts when the connection drop.
	//
	// A positive integer.
	// Default: - 3.
	//
	TcpMaxReconnect *float64 `field:"optional" json:"tcpMaxReconnect" yaml:"tcpMaxReconnect"`
	// TCP Only The number of seconds to wait between reconnection attempts.
	//
	// A positive integer.
	// Default: - 1.
	//
	TcpReconnectDelay awscdk.Duration `field:"optional" json:"tcpReconnectDelay" yaml:"tcpReconnectDelay"`
}


package awsmedialive


// Properties for defining a `CfnChannel`.
//
// Example:
//
//
type CfnChannelProps struct {
	// Specification of CDI inputs for this channel.
	CdiInputSpecification interface{} `field:"optional" json:"cdiInputSpecification" yaml:"cdiInputSpecification"`
	// The class for this channel.
	//
	// For a channel with two pipelines, the class is STANDARD. For a channel with one pipeline, the class is SINGLE_PIPELINE.
	ChannelClass *string `field:"optional" json:"channelClass" yaml:"channelClass"`
	// The settings that identify the destination for the outputs in this MediaLive output package.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
	// The encoding configuration for the output content.
	EncoderSettings interface{} `field:"optional" json:"encoderSettings" yaml:"encoderSettings"`
	// The list of input attachments for the channel.
	InputAttachments interface{} `field:"optional" json:"inputAttachments" yaml:"inputAttachments"`
	// The input specification for this channel.
	//
	// It specifies the key characteristics of the inputs for this channel: the maximum bitrate, the resolution, and the codec.
	InputSpecification interface{} `field:"optional" json:"inputSpecification" yaml:"inputSpecification"`
	// The verbosity for logging activity for this channel.
	//
	// Charges for logging (which are generated through Amazon CloudWatch Logging) are higher for higher verbosities.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// `AWS::MediaLive::Channel.Maintenance`.
	Maintenance interface{} `field:"optional" json:"maintenance" yaml:"maintenance"`
	// A name for this audio selector.
	//
	// The AudioDescription (in an output) references this name in order to identify a specific input audio to include in that output.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The IAM role for MediaLive to assume when running this channel.
	//
	// The role is identified by its ARN.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// A collection of tags for this channel.
	//
	// Each tag is a key-value pair.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Settings to enable VPC mode in the channel, so that the endpoints for all outputs are in your VPC.
	Vpc interface{} `field:"optional" json:"vpc" yaml:"vpc"`
}


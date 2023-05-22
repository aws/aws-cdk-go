package awsmsk


// Properties for defining a `CfnConfiguration`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnConfigurationProps := &CfnConfigurationProps{
//   	Name: jsii.String("name"),
//   	ServerProperties: jsii.String("serverProperties"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	KafkaVersionsList: []*string{
//   		jsii.String("kafkaVersionsList"),
//   	},
//   }
//
type CfnConfigurationProps struct {
	// The name of the configuration.
	//
	// Configuration names are strings that match the regex "^[0-9A-Za-z][0-9A-Za-z-]{0,}$".
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contents of the server.properties file. When using the API, you must ensure that the contents of the file are base64 encoded. When using the console, the SDK, or the CLI, the contents of server.properties can be in plaintext.
	ServerProperties *string `field:"required" json:"serverProperties" yaml:"serverProperties"`
	// The description of the configuration.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// `AWS::MSK::Configuration.KafkaVersionsList`.
	KafkaVersionsList *[]*string `field:"optional" json:"kafkaVersionsList" yaml:"kafkaVersionsList"`
}


package awsecs


// The configuration to use when creating a log driver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   logDriverConfig := &logDriverConfig{
//   	logDriver: jsii.String("logDriver"),
//
//   	// the properties below are optional
//   	options: map[string]*string{
//   		"optionsKey": jsii.String("options"),
//   	},
//   	secretOptions: []secretProperty{
//   		&secretProperty{
//   			name: jsii.String("name"),
//   			valueFrom: jsii.String("valueFrom"),
//   		},
//   	},
//   }
//
type LogDriverConfig struct {
	// The log driver to use for the container.
	//
	// The valid values listed for this parameter are log drivers
	// that the Amazon ECS container agent can communicate with by default.
	//
	// For tasks using the Fargate launch type, the supported log drivers are awslogs, splunk, and awsfirelens.
	// For tasks using the EC2 launch type, the supported log drivers are awslogs, fluentd, gelf, json-file, journald,
	// logentries,syslog, splunk, and awsfirelens.
	//
	// For more information about using the awslogs log driver, see
	// [Using the awslogs Log Driver](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_awslogs.html)
	// in the Amazon Elastic Container Service Developer Guide.
	//
	// For more information about using the awsfirelens log driver, see
	// [Custom Log Routing](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_firelens.html)
	// in the Amazon Elastic Container Service Developer Guide.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration.
	SecretOptions *[]*CfnTaskDefinition_SecretProperty `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


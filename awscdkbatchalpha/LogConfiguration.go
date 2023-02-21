// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Log configuration options to send to a custom log driver for the container.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &JobDefinitionProps{
//   	Container: &JobDefinitionContainer{
//   		Image: ecs.EcrImage_FromRegistry(jsii.String("docker/whalesay")),
//   		LogConfiguration: &LogConfiguration{
//   			LogDriver: batch.LogDriver_AWSLOGS,
//   			Options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			SecretOptions: []exposedSecret{
//   				batch.*exposedSecret_FromParametersStore(jsii.String("xyz"), ssm.StringParameter_FromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type LogConfiguration struct {
	// The log driver to use for the container.
	// Experimental.
	LogDriver LogDriver `field:"required" json:"logDriver" yaml:"logDriver"`
	// The configuration options to send to the log driver.
	// Experimental.
	Options interface{} `field:"optional" json:"options" yaml:"options"`
	// The secrets to pass to the log configuration as options.
	//
	// For more information, see https://docs.aws.amazon.com/batch/latest/userguide/specifying-sensitive-data-secrets.html#secrets-logconfig
	// Experimental.
	SecretOptions *[]ExposedSecret `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}


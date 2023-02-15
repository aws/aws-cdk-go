// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// The log driver to use for the container.
//
// Example:
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   batch.NewJobDefinition(this, jsii.String("job-def"), &jobDefinitionProps{
//   	container: &jobDefinitionContainer{
//   		image: ecs.ecrImage.fromRegistry(jsii.String("docker/whalesay")),
//   		logConfiguration: &logConfiguration{
//   			logDriver: batch.logDriver_AWSLOGS,
//   			options: map[string]*string{
//   				"awslogs-region": jsii.String("us-east-1"),
//   			},
//   			secretOptions: []exposedSecret{
//   				batch.*exposedSecret.fromParametersStore(jsii.String("xyz"), ssm.stringParameter.fromStringParameterName(this, jsii.String("parameter"), jsii.String("xyz"))),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type LogDriver string

const (
	// Specifies the Amazon CloudWatch Logs logging driver.
	// Experimental.
	LogDriver_AWSLOGS LogDriver = "AWSLOGS"
	// Specifies the Fluentd logging driver.
	// Experimental.
	LogDriver_FLUENTD LogDriver = "FLUENTD"
	// Specifies the Graylog Extended Format (GELF) logging driver.
	// Experimental.
	LogDriver_GELF LogDriver = "GELF"
	// Specifies the journald logging driver.
	// Experimental.
	LogDriver_JOURNALD LogDriver = "JOURNALD"
	// Specifies the logentries logging driver.
	// Experimental.
	LogDriver_LOGENTRIES LogDriver = "LOGENTRIES"
	// Specifies the JSON file logging driver.
	// Experimental.
	LogDriver_JSON_FILE LogDriver = "JSON_FILE"
	// Specifies the Splunk logging driver.
	// Experimental.
	LogDriver_SPLUNK LogDriver = "SPLUNK"
	// Specifies the syslog logging driver.
	// Experimental.
	LogDriver_SYSLOG LogDriver = "SYSLOG"
)


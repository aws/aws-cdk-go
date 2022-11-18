package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Construction properties of the {@link JobDefinition} construct.
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
type JobDefinitionProps struct {
	// An object with various properties specific to container-based jobs.
	// Experimental.
	Container *JobDefinitionContainer `field:"required" json:"container" yaml:"container"`
	// The name of the job definition.
	//
	// Up to 128 letters (uppercase and lowercase), numbers, hyphens, and underscores are allowed.
	// Experimental.
	JobDefinitionName *string `field:"optional" json:"jobDefinitionName" yaml:"jobDefinitionName"`
	// An object with various properties specific to multi-node parallel jobs.
	// Experimental.
	NodeProps IMultiNodeProps `field:"optional" json:"nodeProps" yaml:"nodeProps"`
	// When you submit a job, you can specify parameters that should replace the placeholders or override the default job definition parameters.
	//
	// Parameters
	// in job submission requests take precedence over the defaults in a job definition.
	// This allows you to use the same job definition for multiple jobs that use the same
	// format, and programmatically change values in the command at submission time.
	// Experimental.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The platform capabilities required by the job definition.
	// Experimental.
	PlatformCapabilities *[]PlatformCapabilities `field:"optional" json:"platformCapabilities" yaml:"platformCapabilities"`
	// The number of times to move a job to the RUNNABLE status.
	//
	// You may specify between 1 and
	// 10 attempts. If the value of attempts is greater than one, the job is retried on failure
	// the same number of attempts as the value.
	// Experimental.
	RetryAttempts *float64 `field:"optional" json:"retryAttempts" yaml:"retryAttempts"`
	// The timeout configuration for jobs that are submitted with this job definition.
	//
	// You can specify
	// a timeout duration after which AWS Batch terminates your jobs if they have not finished.
	// Experimental.
	Timeout awscdk.Duration `field:"optional" json:"timeout" yaml:"timeout"`
}


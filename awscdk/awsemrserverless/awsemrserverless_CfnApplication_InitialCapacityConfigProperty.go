package awsemrserverless


// The initial capacity configuration per worker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initialCapacityConfigProperty := &initialCapacityConfigProperty{
//   	workerConfiguration: &workerConfigurationProperty{
//   		cpu: jsii.String("cpu"),
//   		memory: jsii.String("memory"),
//
//   		// the properties below are optional
//   		disk: jsii.String("disk"),
//   	},
//   	workerCount: jsii.Number(123),
//   }
//
type CfnApplication_InitialCapacityConfigProperty struct {
	// The resource configuration of the initial capacity configuration.
	WorkerConfiguration interface{} `field:"required" json:"workerConfiguration" yaml:"workerConfiguration"`
	// The number of workers in the initial capacity configuration.
	//
	// *Minimum* : 1
	//
	// *Maximum* : 1000000.
	WorkerCount *float64 `field:"required" json:"workerCount" yaml:"workerCount"`
}


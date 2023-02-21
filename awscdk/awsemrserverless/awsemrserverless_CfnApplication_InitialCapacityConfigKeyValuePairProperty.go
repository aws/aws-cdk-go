package awsemrserverless


// The initial capacity configuration per worker.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   initialCapacityConfigKeyValuePairProperty := &initialCapacityConfigKeyValuePairProperty{
//   	key: jsii.String("key"),
//   	value: &initialCapacityConfigProperty{
//   		workerConfiguration: &workerConfigurationProperty{
//   			cpu: jsii.String("cpu"),
//   			memory: jsii.String("memory"),
//
//   			// the properties below are optional
//   			disk: jsii.String("disk"),
//   		},
//   		workerCount: jsii.Number(123),
//   	},
//   }
//
type CfnApplication_InitialCapacityConfigKeyValuePairProperty struct {
	// The worker type for an analytics framework.
	//
	// For Spark applications, the key can either be set to `Driver` or `Executor` . For Hive applications, it can be set to `HiveDriver` or `TezTask` .
	//
	// *Minimum* : 1
	//
	// *Maximum* : 50
	//
	// *Pattern* : `^[a-zA-Z]+[-_]*[a-zA-Z]+$`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the initial capacity configuration per worker.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}


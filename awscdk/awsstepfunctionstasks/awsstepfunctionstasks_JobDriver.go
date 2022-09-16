package awsstepfunctionstasks


// Specify the driver that the EMR Containers job runs on.
//
// The job driver is used to provide an input for the job that will be run.
//
// Example:
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &emrContainersStartJobRunProps{
//   	virtualCluster: tasks.virtualClusterInput.fromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	releaseLabel: tasks.releaseLabel_EMR_6_2_0(),
//   	jobName: jsii.String("EMR-Containers-Job"),
//   	jobDriver: &jobDriver{
//   		sparkSubmitJobDriver: &sparkSubmitJobDriver{
//   			entryPoint: sfn.taskInput.fromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   		},
//   	},
//   	applicationConfig: []applicationConfiguration{
//   		&applicationConfiguration{
//   			classification: tasks.classification_SPARK_DEFAULTS(),
//   			properties: map[string]*string{
//   				"spark.executor.instances": jsii.String("1"),
//   				"spark.executor.memory": jsii.String("512M"),
//   			},
//   		},
//   	},
//   })
//
type JobDriver struct {
	// The job driver parameters specified for spark submit.
	// See: https://docs.aws.amazon.com/emr-on-eks/latest/APIReference/API_SparkSubmitJobDriver.html
	//
	SparkSubmitJobDriver *SparkSubmitJobDriver `field:"required" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}


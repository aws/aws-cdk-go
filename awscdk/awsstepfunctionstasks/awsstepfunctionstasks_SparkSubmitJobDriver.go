package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// The information about job driver for Spark submit.
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
type SparkSubmitJobDriver struct {
	// The entry point of job application.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 256.
	EntryPoint awsstepfunctions.TaskInput `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// The arguments for a job application in a task input object containing an array of strings.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 10280.
	EntryPointArguments awsstepfunctions.TaskInput `field:"optional" json:"entryPointArguments" yaml:"entryPointArguments"`
	// The Spark submit parameters that are used for job runs.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 102400.
	SparkSubmitParameters *string `field:"optional" json:"sparkSubmitParameters" yaml:"sparkSubmitParameters"`
}


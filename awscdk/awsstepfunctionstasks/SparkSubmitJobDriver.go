package awsstepfunctionstasks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// The information about job driver for Spark submit.
//
// Example:
//   tasks.NewEmrContainersStartJobRun(this, jsii.String("EMR Containers Start Job Run"), &EmrContainersStartJobRunProps{
//   	VirtualCluster: tasks.VirtualClusterInput_FromVirtualClusterId(jsii.String("de92jdei2910fwedz")),
//   	ReleaseLabel: tasks.ReleaseLabel_EMR_6_2_0(),
//   	JobName: jsii.String("EMR-Containers-Job"),
//   	JobDriver: &JobDriver{
//   		SparkSubmitJobDriver: &SparkSubmitJobDriver{
//   			EntryPoint: sfn.TaskInput_FromText(jsii.String("local:///usr/lib/spark/examples/src/main/python/pi.py")),
//   		},
//   	},
//   	ApplicationConfig: []ApplicationConfiguration{
//   		&ApplicationConfiguration{
//   			Classification: tasks.Classification_SPARK_DEFAULTS(),
//   			Properties: map[string]*string{
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
	// Default: - No arguments defined.
	//
	EntryPointArguments awsstepfunctions.TaskInput `field:"optional" json:"entryPointArguments" yaml:"entryPointArguments"`
	// The Spark submit parameters that are used for job runs.
	//
	// Length Constraints: Minimum length of 1. Maximum length of 102400.
	// Default: - No spark submit parameters.
	//
	SparkSubmitParameters *string `field:"optional" json:"sparkSubmitParameters" yaml:"sparkSubmitParameters"`
}


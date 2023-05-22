package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The classification within a EMR Containers application configuration.
//
// Class can be extended to add other classifications.
// For example, new Classification('xxx-yyy');.
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
//   	ApplicationConfig: []applicationConfiguration{
//   		&applicationConfiguration{
//   			Classification: tasks.Classification_SPARK_DEFAULTS(),
//   			Properties: map[string]*string{
//   				"spark.executor.instances": jsii.String("1"),
//   				"spark.executor.memory": jsii.String("512M"),
//   			},
//   		},
//   	},
//   })
//
// Experimental.
type Classification interface {
	// A literal string in case a new EMR classification is released, if not already defined.
	// Experimental.
	ClassificationStatement() *string
}

// The jsii proxy struct for Classification
type jsiiProxy_Classification struct {
	_ byte // padding
}

func (j *jsiiProxy_Classification) ClassificationStatement() *string {
	var returns *string
	_jsii_.Get(
		j,
		"classificationStatement",
		&returns,
	)
	return returns
}


// Creates a new Classification.
// Experimental.
func NewClassification(classificationStatement *string) Classification {
	_init_.Initialize()

	if err := validateNewClassificationParameters(classificationStatement); err != nil {
		panic(err)
	}
	j := jsiiProxy_Classification{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.Classification",
		[]interface{}{classificationStatement},
		&j,
	)

	return &j
}

// Creates a new Classification.
// Experimental.
func NewClassification_Override(c Classification, classificationStatement *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.Classification",
		[]interface{}{classificationStatement},
		c,
	)
}

func Classification_SPARK() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK",
		&returns,
	)
	return returns
}

func Classification_SPARK_DEFAULTS() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK_DEFAULTS",
		&returns,
	)
	return returns
}

func Classification_SPARK_ENV() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK_ENV",
		&returns,
	)
	return returns
}

func Classification_SPARK_HIVE_SITE() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK_HIVE_SITE",
		&returns,
	)
	return returns
}

func Classification_SPARK_LOG4J() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK_LOG4J",
		&returns,
	)
	return returns
}

func Classification_SPARK_METRICS() Classification {
	_init_.Initialize()
	var returns Classification
	_jsii_.StaticGet(
		"monocdk.aws_stepfunctions_tasks.Classification",
		"SPARK_METRICS",
		&returns,
	)
	return returns
}


package awsemrcontainers


// Specify the driver that the job runs on.
//
// Exactly one of the two available job drivers is required, either SparkSqlJobDriver or SparkSubmitJobDriver.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   jobDriverProperty := &JobDriverProperty{
//   	SparkSqlJobDriver: &SparkSqlJobDriverProperty{
//   		EntryPoint: jsii.String("entryPoint"),
//   		SparkSqlParameters: jsii.String("sparkSqlParameters"),
//   	},
//   	SparkSubmitJobDriver: &SparkSubmitJobDriverProperty{
//   		EntryPoint: jsii.String("entryPoint"),
//
//   		// the properties below are optional
//   		EntryPointArguments: []*string{
//   			jsii.String("entryPointArguments"),
//   		},
//   		SparkSubmitParameters: jsii.String("sparkSubmitParameters"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobdriver.html
//
type CfnJobTemplate_JobDriverProperty struct {
	// The job driver for job type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobdriver.html#cfn-emrcontainers-jobtemplate-jobdriver-sparksqljobdriver
	//
	SparkSqlJobDriver interface{} `field:"optional" json:"sparkSqlJobDriver" yaml:"sparkSqlJobDriver"`
	// The information about job driver for Spark submit.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-jobdriver.html#cfn-emrcontainers-jobtemplate-jobdriver-sparksubmitjobdriver
	//
	SparkSubmitJobDriver interface{} `field:"optional" json:"sparkSubmitJobDriver" yaml:"sparkSubmitJobDriver"`
}


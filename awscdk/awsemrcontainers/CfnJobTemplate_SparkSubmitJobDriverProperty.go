package awsemrcontainers


// The information about job driver for Spark submit.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparkSubmitJobDriverProperty := &SparkSubmitJobDriverProperty{
//   	EntryPoint: jsii.String("entryPoint"),
//
//   	// the properties below are optional
//   	EntryPointArguments: []*string{
//   		jsii.String("entryPointArguments"),
//   	},
//   	SparkSubmitParameters: jsii.String("sparkSubmitParameters"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksubmitjobdriver.html
//
type CfnJobTemplate_SparkSubmitJobDriverProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksubmitjobdriver.html#cfn-emrcontainers-jobtemplate-sparksubmitjobdriver-entrypoint
	//
	EntryPoint *string `field:"required" json:"entryPoint" yaml:"entryPoint"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksubmitjobdriver.html#cfn-emrcontainers-jobtemplate-sparksubmitjobdriver-entrypointarguments
	//
	EntryPointArguments *[]*string `field:"optional" json:"entryPointArguments" yaml:"entryPointArguments"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emrcontainers-jobtemplate-sparksubmitjobdriver.html#cfn-emrcontainers-jobtemplate-sparksubmitjobdriver-sparksubmitparameters
	//
	SparkSubmitParameters *string `field:"optional" json:"sparkSubmitParameters" yaml:"sparkSubmitParameters"`
}


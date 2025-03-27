package awsdatazone


// Spark Glue Properties Input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparkGluePropertiesInputProperty := &SparkGluePropertiesInputProperty{
//   	AdditionalArgs: &SparkGlueArgsProperty{
//   		Connection: jsii.String("connection"),
//   	},
//   	GlueConnectionName: jsii.String("glueConnectionName"),
//   	GlueVersion: jsii.String("glueVersion"),
//   	IdleTimeout: jsii.Number(123),
//   	JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   	NumberOfWorkers: jsii.Number(123),
//   	PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   	WorkerType: jsii.String("workerType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html
//
type CfnConnection_SparkGluePropertiesInputProperty struct {
	// Spark Glue Args.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-additionalargs
	//
	AdditionalArgs interface{} `field:"optional" json:"additionalArgs" yaml:"additionalArgs"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-glueconnectionname
	//
	GlueConnectionName *string `field:"optional" json:"glueConnectionName" yaml:"glueConnectionName"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-glueversion
	//
	GlueVersion *string `field:"optional" json:"glueVersion" yaml:"glueVersion"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-idletimeout
	//
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-javavirtualenv
	//
	JavaVirtualEnv *string `field:"optional" json:"javaVirtualEnv" yaml:"javaVirtualEnv"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-numberofworkers
	//
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-pythonvirtualenv
	//
	PythonVirtualEnv *string `field:"optional" json:"pythonVirtualEnv" yaml:"pythonVirtualEnv"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkgluepropertiesinput.html#cfn-datazone-connection-sparkgluepropertiesinput-workertype
	//
	WorkerType *string `field:"optional" json:"workerType" yaml:"workerType"`
}


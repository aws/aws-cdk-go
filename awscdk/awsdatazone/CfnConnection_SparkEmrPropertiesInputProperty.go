package awsdatazone


// The Spark EMR properties.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sparkEmrPropertiesInputProperty := &SparkEmrPropertiesInputProperty{
//   	ComputeArn: jsii.String("computeArn"),
//   	InstanceProfileArn: jsii.String("instanceProfileArn"),
//   	JavaVirtualEnv: jsii.String("javaVirtualEnv"),
//   	LogUri: jsii.String("logUri"),
//   	ManagedEndpointArn: jsii.String("managedEndpointArn"),
//   	PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   	RuntimeRole: jsii.String("runtimeRole"),
//   	TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html
//
type CfnConnection_SparkEmrPropertiesInputProperty struct {
	// The compute ARN of Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-computearn
	//
	ComputeArn *string `field:"optional" json:"computeArn" yaml:"computeArn"`
	// The instance profile ARN of Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-instanceprofilearn
	//
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The java virtual env of the Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-javavirtualenv
	//
	JavaVirtualEnv *string `field:"optional" json:"javaVirtualEnv" yaml:"javaVirtualEnv"`
	// The log URI of the Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-managedendpointarn
	//
	ManagedEndpointArn *string `field:"optional" json:"managedEndpointArn" yaml:"managedEndpointArn"`
	// The Python virtual env of the Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-pythonvirtualenv
	//
	PythonVirtualEnv *string `field:"optional" json:"pythonVirtualEnv" yaml:"pythonVirtualEnv"`
	// The runtime role of the Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-runtimerole
	//
	RuntimeRole *string `field:"optional" json:"runtimeRole" yaml:"runtimeRole"`
	// The certificates S3 URI of the Spark EMR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-trustedcertificatess3uri
	//
	TrustedCertificatesS3Uri *string `field:"optional" json:"trustedCertificatesS3Uri" yaml:"trustedCertificatesS3Uri"`
}


package awsdatazone


// Spark EMR Properties Input.
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
//   	PythonVirtualEnv: jsii.String("pythonVirtualEnv"),
//   	RuntimeRole: jsii.String("runtimeRole"),
//   	TrustedCertificatesS3Uri: jsii.String("trustedCertificatesS3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html
//
type CfnConnection_SparkEmrPropertiesInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-computearn
	//
	ComputeArn *string `field:"optional" json:"computeArn" yaml:"computeArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-instanceprofilearn
	//
	InstanceProfileArn *string `field:"optional" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-javavirtualenv
	//
	JavaVirtualEnv *string `field:"optional" json:"javaVirtualEnv" yaml:"javaVirtualEnv"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-loguri
	//
	LogUri *string `field:"optional" json:"logUri" yaml:"logUri"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-pythonvirtualenv
	//
	PythonVirtualEnv *string `field:"optional" json:"pythonVirtualEnv" yaml:"pythonVirtualEnv"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-runtimerole
	//
	RuntimeRole *string `field:"optional" json:"runtimeRole" yaml:"runtimeRole"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datazone-connection-sparkemrpropertiesinput.html#cfn-datazone-connection-sparkemrpropertiesinput-trustedcertificatess3uri
	//
	TrustedCertificatesS3Uri *string `field:"optional" json:"trustedCertificatesS3Uri" yaml:"trustedCertificatesS3Uri"`
}


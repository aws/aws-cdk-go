package awsecs


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FilesVolumeConfigurationProperty := &S3FilesVolumeConfigurationProperty{
//   	FileSystemArn: jsii.String("fileSystemArn"),
//
//   	// the properties below are optional
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryptionPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html
//
type CfnTaskDefinition_S3FilesVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-filesystemarn
	//
	FileSystemArn *string `field:"required" json:"fileSystemArn" yaml:"fileSystemArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-s3filesvolumeconfiguration.html#cfn-ecs-taskdefinition-s3filesvolumeconfiguration-transitencryptionport
	//
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}


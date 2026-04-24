package awsbatch


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   s3FilesVolumeConfigurationProperty := &S3FilesVolumeConfigurationProperty{
//   	AccessPointArn: jsii.String("accessPointArn"),
//   	FileSystemArn: jsii.String("fileSystemArn"),
//   	RootDirectory: jsii.String("rootDirectory"),
//   	TransitEncryptionPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-s3filesvolumeconfiguration.html
//
type CfnJobDefinitionPropsMixin_S3FilesVolumeConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-s3filesvolumeconfiguration.html#cfn-batch-jobdefinition-s3filesvolumeconfiguration-accesspointarn
	//
	AccessPointArn *string `field:"optional" json:"accessPointArn" yaml:"accessPointArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-s3filesvolumeconfiguration.html#cfn-batch-jobdefinition-s3filesvolumeconfiguration-filesystemarn
	//
	FileSystemArn *string `field:"optional" json:"fileSystemArn" yaml:"fileSystemArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-s3filesvolumeconfiguration.html#cfn-batch-jobdefinition-s3filesvolumeconfiguration-rootdirectory
	//
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-s3filesvolumeconfiguration.html#cfn-batch-jobdefinition-s3filesvolumeconfiguration-transitencryptionport
	//
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}


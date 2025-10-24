package awsefs

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// EFS Replication Configuration.
//
// Example:
//   var vpc Vpc
//
//
//   // auto generate a regional replication destination file system
//   // auto generate a regional replication destination file system
//   efs.NewFileSystem(this, jsii.String("RegionalReplicationFileSystem"), &FileSystemProps{
//   	Vpc: Vpc,
//   	ReplicationConfiguration: efs.ReplicationConfiguration_RegionalFileSystem(jsii.String("us-west-2")),
//   })
//
//   // auto generate a one zone replication destination file system
//   // auto generate a one zone replication destination file system
//   efs.NewFileSystem(this, jsii.String("OneZoneReplicationFileSystem"), &FileSystemProps{
//   	Vpc: Vpc,
//   	ReplicationConfiguration: efs.ReplicationConfiguration_OneZoneFileSystem(jsii.String("us-east-1"), jsii.String("us-east-1a")),
//   })
//
//   destinationFileSystem := efs.NewFileSystem(this, jsii.String("DestinationFileSystem"), &FileSystemProps{
//   	Vpc: Vpc,
//   	// set as the read-only file system for use as a replication destination
//   	ReplicationOverwriteProtection: efs.ReplicationOverwriteProtection_DISABLED,
//   })
//   // specify the replication destination file system
//   // specify the replication destination file system
//   efs.NewFileSystem(this, jsii.String("ReplicationFileSystem"), &FileSystemProps{
//   	Vpc: Vpc,
//   	ReplicationConfiguration: efs.ReplicationConfiguration_ExistingFileSystem(destinationFileSystem),
//   })
//
type ReplicationConfiguration interface {
	// The availability zone name of the destination file system.
	//
	// One zone file system is used as the destination file system when this property is set.
	AvailabilityZone() *string
	// The existing destination file system for the replication.
	DestinationFileSystem() IFileSystem
	// AWS KMS key used to protect the encrypted file system.
	KmsKey() awskms.IKey
	// The AWS Region in which the destination file system is located.
	Region() *string
}

// The jsii proxy struct for ReplicationConfiguration
type jsiiProxy_ReplicationConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_ReplicationConfiguration) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationConfiguration) DestinationFileSystem() IFileSystem {
	var returns IFileSystem
	_jsii_.Get(
		j,
		"destinationFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationConfiguration) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ReplicationConfiguration) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}


func NewReplicationConfiguration_Override(r ReplicationConfiguration, options *ReplicationConfigurationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_efs.ReplicationConfiguration",
		[]interface{}{options},
		r,
	)
}

// Specify the existing destination file system for the replication.
func ReplicationConfiguration_ExistingFileSystem(destinationFileSystem IFileSystem) ReplicationConfiguration {
	_init_.Initialize()

	if err := validateReplicationConfiguration_ExistingFileSystemParameters(destinationFileSystem); err != nil {
		panic(err)
	}
	var returns ReplicationConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.ReplicationConfiguration",
		"existingFileSystem",
		[]interface{}{destinationFileSystem},
		&returns,
	)

	return returns
}

// Create a new one zone destination file system for the replication.
func ReplicationConfiguration_OneZoneFileSystem(region *string, availabilityZone *string, kmsKey awskms.IKey) ReplicationConfiguration {
	_init_.Initialize()

	if err := validateReplicationConfiguration_OneZoneFileSystemParameters(region, availabilityZone); err != nil {
		panic(err)
	}
	var returns ReplicationConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.ReplicationConfiguration",
		"oneZoneFileSystem",
		[]interface{}{region, availabilityZone, kmsKey},
		&returns,
	)

	return returns
}

// Create a new regional destination file system for the replication.
func ReplicationConfiguration_RegionalFileSystem(region *string, kmsKey awskms.IKey) ReplicationConfiguration {
	_init_.Initialize()

	var returns ReplicationConfiguration

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_efs.ReplicationConfiguration",
		"regionalFileSystem",
		[]interface{}{region, kmsKey},
		&returns,
	)

	return returns
}


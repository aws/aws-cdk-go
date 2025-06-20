# Amazon FSx Construct Library

[Amazon FSx](https://docs.aws.amazon.com/fsx/?id=docs_gateway) provides fully managed third-party file systems with the
native compatibility and feature sets for workloads such as Microsoft Windows–based storage, high-performance computing,
machine learning, and electronic design automation.

Amazon FSx supports two file system types: [Lustre](https://docs.aws.amazon.com/fsx/latest/LustreGuide/index.html) and
[Windows](https://docs.aws.amazon.com/fsx/latest/WindowsGuide/index.html) File Server.

## FSx for Lustre

Amazon FSx for Lustre makes it easy and cost-effective to launch and run the popular, high-performance Lustre file
system. You use Lustre for workloads where speed matters, such as machine learning, high performance computing (HPC),
video processing, and financial modeling.

The open-source Lustre file system is designed for applications that require fast storage—where you want your storage
to keep up with your compute. Lustre was built to solve the problem of quickly and cheaply processing the world's
ever-growing datasets. It's a widely used file system designed for the fastest computers in the world. It provides
submillisecond latencies, up to hundreds of GBps of throughput, and up to millions of IOPS. For more information on
Lustre, see the [Lustre website](http://lustre.org/).

As a fully managed service, Amazon FSx makes it easier for you to use Lustre for workloads where storage speed matters.
Amazon FSx for Lustre eliminates the traditional complexity of setting up and managing Lustre file systems, enabling
you to spin up and run a battle-tested high-performance file system in minutes. It also provides multiple deployment
options so you can optimize cost for your needs.

Amazon FSx for Lustre is POSIX-compliant, so you can use your current Linux-based applications without having to make
any changes. Amazon FSx for Lustre provides a native file system interface and works as any file system does with your
Linux operating system. It also provides read-after-write consistency and supports file locking.

### Installation

Import to your project:

```go
import fsx "github.com/aws/aws-cdk-go/awscdk"
```

### Basic Usage

Setup required properties and create:

```go
var vpc vpc


fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	LustreConfiguration: &LustreConfiguration{
		DeploymentType: fsx.LustreDeploymentType_SCRATCH_2,
	},
	StorageCapacityGiB: jsii.Number(1200),
	Vpc: Vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
})
```

### File System Type Version

You can set [the Lustre version for the file system](https://docs.aws.amazon.com/ja_jp/AWSCloudFormation/latest/UserGuide/aws-resource-fsx-filesystem.html#cfn-fsx-filesystem-filesystemtypeversion). To do this, use the `fileSystemTypeVersion` property:

```go
var vpc vpc


fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	LustreConfiguration: &LustreConfiguration{
		DeploymentType: fsx.LustreDeploymentType_SCRATCH_2,
	},
	StorageCapacityGiB: jsii.Number(1200),
	Vpc: Vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
	FileSystemTypeVersion: fsx.FileSystemTypeVersion_V_2_15,
})
```

**Note**: The `fileSystemTypeVersion` has a restrictions on the values that can be set based on the `deploymentType`.

* `V_2_10` is supported by the Scratch and `PERSISTENT_1` deployment types.
* `V_2_12` is supported by all Lustre deployment types.
* `V_2_15` is supported by all Lustre deployment types and is recommended for all new file systems.

**Note**: The default value of `fileSystemTypeVersion` is `V_2_10` except for `PERSISTENT_2` deployment type where the default value is `V_2_12`.

### Connecting

To control who can access the file system, use the `.connections` attribute. FSx has a fixed default port, so you don't
need to specify the port. This example allows an EC2 instance to connect to a file system:

```go
var fileSystem lustreFileSystem
var instance instance


fileSystem.Connections.AllowDefaultPortFrom(instance)
```

### Mounting

The LustreFileSystem Construct exposes both the DNS name of the file system as well as its mount name, which can be
used to mount the file system on an EC2 instance. The following example shows how to bring up a file system and EC2
instance, and then use User Data to mount the file system on the instance at start-up:

```go
import iam "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc

lustreConfiguration := map[string]lustreDeploymentType{
	"deploymentType": fsx.lustreDeploymentType_SCRATCH_2,
}

fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	LustreConfiguration: LustreConfiguration,
	StorageCapacityGiB: jsii.Number(1200),
	Vpc: Vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
})

inst := ec2.NewInstance(this, jsii.String("inst"), &InstanceProps{
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T2, ec2.InstanceSize_LARGE),
	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
	}),
	Vpc: Vpc,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
})
fs.Connections.AllowDefaultPortFrom(inst)

// Need to give the instance access to read information about FSx to determine the file system's mount name.
inst.Role.AddManagedPolicy(iam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("AmazonFSxReadOnlyAccess")))

mountPath := "/mnt/fsx"
dnsName := fs.DnsName
mountName := fs.MountName

inst.UserData.AddCommands(jsii.String("set -eux"), jsii.String("yum update -y"), jsii.String("amazon-linux-extras install -y lustre2.10"),
// Set up the directory to mount the file system to and change the owner to the AL2 default ec2-user.
fmt.Sprintf("mkdir -p %v", mountPath),
fmt.Sprintf("chmod 777 %v", mountPath),
fmt.Sprintf("chown ec2-user:ec2-user %v", mountPath),
// Set the file system up to mount automatically on start up and mount it.
fmt.Sprintf("echo \"%v@tcp:/%v %v lustre defaults,noatime,flock,_netdev 0 0\" >> /etc/fstab", dnsName, mountName, mountPath), jsii.String("mount -a"))
```

### Importing an existing Lustre filesystem

An FSx for Lustre file system can be imported with `fromLustreFileSystemAttributes(this, id, attributes)`. The
following example lays out how you could import the SecurityGroup a file system belongs to, use that to import the file
system, and then also import the VPC the file system is in and add an EC2 instance to it, giving it access to the file
system.

```go
sg := ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("FsxSecurityGroup"), jsii.String("{SECURITY-GROUP-ID}"))
fs := fsx.LustreFileSystem_FromLustreFileSystemAttributes(this, jsii.String("FsxLustreFileSystem"), &FileSystemAttributes{
	DnsName: jsii.String("{FILE-SYSTEM-DNS-NAME}"),
	FileSystemId: jsii.String("{FILE-SYSTEM-ID}"),
	SecurityGroup: sg,
})

vpc := ec2.Vpc_FromVpcAttributes(this, jsii.String("Vpc"), &VpcAttributes{
	AvailabilityZones: []*string{
		jsii.String("us-west-2a"),
		jsii.String("us-west-2b"),
	},
	PublicSubnetIds: []*string{
		jsii.String("{US-WEST-2A-SUBNET-ID}"),
		jsii.String("{US-WEST-2B-SUBNET-ID}"),
	},
	VpcId: jsii.String("{VPC-ID}"),
})

inst := ec2.NewInstance(this, jsii.String("inst"), &InstanceProps{
	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_T2, ec2.InstanceSize_LARGE),
	MachineImage: ec2.NewAmazonLinuxImage(&AmazonLinuxImageProps{
		Generation: ec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
	}),
	Vpc: Vpc,
	VpcSubnets: &SubnetSelection{
		SubnetType: ec2.SubnetType_PUBLIC,
	},
})

fs.Connections.AllowDefaultPortFrom(inst)
```

### Lustre Data Repository Association support

The LustreFilesystem Construct supports one [Data Repository Association](https://docs.aws.amazon.com/fsx/latest/LustreGuide/fsx-data-repositories.html) (DRA) to an S3 bucket.  This allows Lustre hierarchical storage management to S3 buckets, which in turn makes it possible to use S3 as a permanent backing store, and use FSx for Lustre as a temporary high performance cache.

Note: CloudFormation does not currently support for `PERSISTENT_2` filesystems, and so neither does CDK.

The following example illustrates setting up a DRA to an S3 bucket, including automated metadata import whenever a file is changed, created or deleted in the S3 bucket:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var vpc vpc
var bucket bucket


lustreConfiguration := map[string]interface{}{
	"deploymentType": fsx.LustreDeploymentType_SCRATCH_2,
	"exportPath": bucket.s3UrlForObject(),
	"importPath": bucket.s3UrlForObject(),
	"autoImportPolicy": fsx.LustreAutoImportPolicy_NEW_CHANGED_DELETED,
}

fs := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	Vpc: vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
	StorageCapacityGiB: jsii.Number(1200),
	LustreConfiguration: LustreConfiguration,
})
```

### Compression

By default, transparent compression of data within FSx for Lustre is switched off.  To enable it, add the following to your `lustreConfiguration`:

```go
lustreConfiguration := map[string]lustreDataCompressionType{
	// ...
	"dataCompressionType": fsx.lustreDataCompressionType_LZ4,
}
```

When you turn data compression on for an existing file system, only newly written files are compressed.  Existing files are not compressed. For more information, see [Compressing previously written files](https://docs.aws.amazon.com/fsx/latest/LustreGuide/data-compression.html#migrate-compression).

### Backups

You can take daily automatic backups by setting `automaticBackupRetention` to a non-zero day in the `lustreConfiguration`.

Additionally, you can set the backup window by specifying the `dailyAutomaticBackupStartTime`.

```go
import cdk "github.com/aws/aws-cdk-go/awscdk"


lustreConfiguration := map[string]interface{}{
	// ...
	"automaticBackupRetention": cdk.Duration_days(jsii.Number(3)),
	 // backup retention
	"copyTagsToBackups": jsii.Boolean(true),
	 // if true, tags are copied to backups
	"dailyAutomaticBackupStartTime": fsx.NewDailyAutomaticBackupStartTime(&DailyAutomaticBackupStartTimeProps{
		"hour": jsii.Number(11),
		"minute": jsii.Number(30),
	}),
}
```

For more information, see [Working with backups
](https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html).

### Storage Type

By default, FSx for Lustre uses SSD storage. To use HDD storage, specify `storageType`:

```go
var vpc vpc


fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	LustreConfiguration: &LustreConfiguration{
		DeploymentType: fsx.LustreDeploymentType_PERSISTENT_1,
	},
	StorageCapacityGiB: jsii.Number(1200),
	Vpc: Vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
	StorageType: fsx.StorageType_HDD,
})
```

**Note:** The HDD storage type is only supported for `PERSISTENT_1` deployment types.

To improve the performance of frequently accessed files by caching up to 20% of the total storage capacity of the file system, set `driveCacheType` to `READ`:

```go
var vpc vpc


fileSystem := fsx.NewLustreFileSystem(this, jsii.String("FsxLustreFileSystem"), &LustreFileSystemProps{
	LustreConfiguration: &LustreConfiguration{
		DeploymentType: fsx.LustreDeploymentType_PERSISTENT_1,
		DriveCacheType: fsx.DriveCacheType_READ,
	},
	StorageCapacityGiB: jsii.Number(1200),
	Vpc: Vpc,
	VpcSubnet: vpc.PrivateSubnets[jsii.Number(0)],
	StorageType: fsx.StorageType_HDD,
})
```

## FSx for Windows File Server

The L2 construct for the FSx for Windows File Server has not yet been implemented. To instantiate an FSx for Windows
file system, the L1 constructs can be used as defined by CloudFormation.

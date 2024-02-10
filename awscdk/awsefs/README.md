# Amazon Elastic File System Construct Library

[Amazon Elastic File System](https://docs.aws.amazon.com/efs/latest/ug/whatisefs.html) (Amazon EFS) provides a simple, scalable,
fully managed elastic NFS file system for use with AWS Cloud services and on-premises resources.
Amazon EFS provides file storage in the AWS Cloud. With Amazon EFS, you can create a file system,
mount the file system on an Amazon EC2 instance, and then read and write data to and from your file system.

This module is part of the [AWS Cloud Development Kit](https://github.com/aws/aws-cdk) project.

## File Systems

Amazon EFS provides elastic, shared file storage that is POSIX-compliant. The file system you create
supports concurrent read and write access from multiple Amazon EC2 instances and is accessible from
all of the Availability Zones in the AWS Region where it is created. Learn more about [EFS file systems](https://docs.aws.amazon.com/efs/latest/ug/creating-using.html)

### Create an Amazon EFS file system

A Virtual Private Cloud (VPC) is required to create an Amazon EFS file system.
The following example creates a file system that is encrypted at rest, running in `General Purpose`
performance mode, and `Bursting` throughput mode and does not transition files to the Infrequent
Access (IA) storage class.

```go
fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
	LifecyclePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
	 // files are not transitioned to infrequent access (IA) storage by default
	PerformanceMode: efs.PerformanceMode_GENERAL_PURPOSE,
	 // default
	OutOfInfrequentAccessPolicy: efs.OutOfInfrequentAccessPolicy_AFTER_1_ACCESS,
	 // files are not transitioned back from (infrequent access) IA to primary storage by default
	TransitionToArchivePolicy: efs.LifecyclePolicy_AFTER_14_DAYS,
	 // files are not transitioned to Archive by default
	ReplicationOverwriteProtection: efs.ReplicationOverwriteProtection_ENABLED,
})
```

⚠️ An Amazon EFS file system's performance mode can't be MAX_IO when its throughputMode is ELASTIC.

⚠️ An Amazon EFS file system's performance mode can't be changed after the file system has been created.
Updating this property will replace the file system.

Any file system that has been created outside the stack can be imported into your CDK app.

Use the `fromFileSystemAttributes()` API to import an existing file system.
Here is an example of giving a role write permissions on a file system.

```go
import iam "github.com/aws/aws-cdk-go/awscdk"


importedFileSystem := efs.FileSystem_FromFileSystemAttributes(this, jsii.String("existingFS"), &FileSystemAttributes{
	FileSystemId: jsii.String("fs-12345678"),
	 // You can also use fileSystemArn instead of fileSystemId.
	SecurityGroup: ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("SG"), jsii.String("sg-123456789"), &SecurityGroupImportOptions{
		AllowAllOutbound: jsii.Boolean(false),
	}),
})
```

### One Zone file system

To initialize a One Zone file system use the `oneZone` property:

```go
var vpc vpc


efs.NewFileSystem(this, jsii.String("OneZoneFileSystem"), &FileSystemProps{
	Vpc: Vpc,
	OneZone: jsii.Boolean(true),
})
```

⚠️ One Zone file systems are not compatible with the MAX_IO performance mode.

⚠️ When `oneZone` is enabled, the file system is automatically placed in the first availability zone of the VPC.
It is not currently possible to specify a different availability zone.

⚠️ When `oneZone` is enabled, mount targets will be created only in the specified availability zone.
This is to prevent deployment failures due to cross-AZ configurations.

⚠️ When `oneZone` is enabled, `vpcSubnets` cannot be specified.

### IAM to control file system data access

You can use both IAM identity policies and resource policies to control client access to Amazon EFS resources in a way that is scalable and optimized for cloud environments. Using IAM, you can permit clients to perform specific actions on a file system, including read-only, write, and root access.

```go
import "github.com/aws/aws-cdk-go/awscdk"


myFileSystemPolicy := iam.NewPolicyDocument(&PolicyDocumentProps{
	Statements: []policyStatement{
		iam.NewPolicyStatement(&PolicyStatementProps{
			Actions: []*string{
				jsii.String("elasticfilesystem:ClientWrite"),
				jsii.String("elasticfilesystem:ClientMount"),
			},
			Principals: []iPrincipal{
				iam.NewAccountRootPrincipal(),
			},
			Resources: []*string{
				jsii.String("*"),
			},
			Conditions: map[string]interface{}{
				"Bool": map[string]*string{
					"elasticfilesystem:AccessedViaMountTarget": jsii.String("true"),
				},
			},
		}),
	},
})

fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
	FileSystemPolicy: myFileSystemPolicy,
})
```

Alternatively, a resource policy can be added later using `addToResourcePolicy(statement)`. Note that this will not work with imported FileSystem.

```go
import iam "github.com/aws/aws-cdk-go/awscdk"

var statement policyStatement

fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
})

fileSystem.addToResourcePolicy(statement)
```

### Permissions

If you need to grant file system permissions to another resource, you can use the `.grant()` API.
As an example, the following code gives `elasticfilesystem:Backup` permissions to an IAM role.

```go
role := iam.NewRole(this, jsii.String("Role"), &RoleProps{
	AssumedBy: iam.NewAnyPrincipal(),
})

fileSystem.grant(role, jsii.String("elasticfilesystem:Backup"))
```

APIs for clients also include `.grantRead()`, `.grantReadWrite()`, and `.grantRootAccess()`. Using these APIs grants access to clients.
Also, by default, the file system policy is updated to only allow access to clients using IAM authentication and deny access to anonymous clients.

```go
role := iam.NewRole(this, jsii.String("ClientRole"), &RoleProps{
	AssumedBy: iam.NewAnyPrincipal(),
})

fileSystem.grantRead(role)
```

You can control this behavior with `allowAnonymousAccess`. The following example continues to allow anonymous client access.

```go
import "github.com/aws/aws-cdk-go/awscdk"


role := iam.NewRole(this, jsii.String("ClientRole"), &RoleProps{
	AssumedBy: iam.NewAnyPrincipal(),
})
fileSystem := efs.NewFileSystem(this, jsii.String("MyEfsFileSystem"), &FileSystemProps{
	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
	AllowAnonymousAccess: jsii.Boolean(true),
})

fileSystem.grantRead(role)
```

### Access Point

An access point is an application-specific view into an EFS file system that applies an operating
system user and group, and a file system path, to any file system request made through the access
point. The operating system user and group override any identity information provided by the NFS
client. The file system path is exposed as the access point's root directory. Applications using
the access point can only access data in its own directory and below. To learn more, see [Mounting a File System Using EFS Access Points](https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html).

Use the `addAccessPoint` API to create an access point from a fileSystem.

```go
fileSystem.AddAccessPoint(jsii.String("AccessPoint"))
```

By default, when you create an access point, the root(`/`) directory is exposed to the client
connecting to the access point. You can specify a custom path with the `path` property.

If `path` does not exist, it will be created with the settings defined in the `creationInfo`.
See [Creating Access Points](https://docs.aws.amazon.com/efs/latest/ug/create-access-point.html) for more details.

Any access point that has been created outside the stack can be imported into your CDK app.

Use the `fromAccessPointAttributes()` API to import an existing access point.

```go
efs.AccessPoint_FromAccessPointAttributes(this, jsii.String("ap"), &AccessPointAttributes{
	AccessPointId: jsii.String("fsap-1293c4d9832fo0912"),
	FileSystem: efs.FileSystem_FromFileSystemAttributes(this, jsii.String("efs"), &FileSystemAttributes{
		FileSystemId: jsii.String("fs-099d3e2f"),
		SecurityGroup: ec2.SecurityGroup_FromSecurityGroupId(this, jsii.String("sg"), jsii.String("sg-51530134")),
	}),
})
```

⚠️ Notice: When importing an Access Point using `fromAccessPointAttributes()`, you must make sure
the mount targets are deployed and their lifecycle state is `available`. Otherwise, you may encounter
the following error when deploying:

> EFS file system <ARN of efs> referenced by access point <ARN of access point of EFS> has
> mount targets created in all availability zones the function will execute in, but not all
> are in the available life cycle state yet. Please wait for them to become available and
> try the request again.

### Connecting

To control who can access the EFS, use the `.connections` attribute. EFS has
a fixed default port, so you don't need to specify the port:

```go
fileSystem.Connections.AllowDefaultPortFrom(instance)
```

Learn more about [managing file system network accessibility](https://docs.aws.amazon.com/efs/latest/ug/manage-fs-access.html)

### Mounting the file system using User Data

After you create a file system, you can create mount targets. Then you can mount the file system on
EC2 instances, containers, and Lambda functions in your virtual private cloud (VPC).

The following example automatically mounts a file system during instance launch.

```go
fileSystem.Connections.AllowDefaultPortFrom(instance)

instance.UserData.AddCommands(jsii.String("yum check-update -y"), jsii.String("yum upgrade -y"), jsii.String("yum install -y amazon-efs-utils"), jsii.String("yum install -y nfs-utils"), jsii.String("file_system_id_1=" + fileSystem.FileSystemId), jsii.String("efs_mount_point_1=/mnt/efs/fs1"), jsii.String("mkdir -p \"${efs_mount_point_1}\""), jsii.String("test -f \"/sbin/mount.efs\" && echo \"${file_system_id_1}:/ ${efs_mount_point_1} efs defaults,_netdev\" >> /etc/fstab || " + "echo \"${file_system_id_1}.efs." + awscdk.stack_Of(this).Region + ".amazonaws.com:/ ${efs_mount_point_1} nfs4 nfsvers=4.1,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport,_netdev 0 0\" >> /etc/fstab"), jsii.String("mount -a -t efs,nfs4 defaults"))
```

Learn more about [mounting EFS file systems](https://docs.aws.amazon.com/efs/latest/ug/mounting-fs.html)

### Deleting

Since file systems are stateful resources, by default the file system will not be deleted when your
stack is deleted.

You can configure the file system to be destroyed on stack deletion by setting a `removalPolicy`

```go
fileSystem := efs.NewFileSystem(this, jsii.String("EfsFileSystem"), &FileSystemProps{
	Vpc: ec2.NewVpc(this, jsii.String("VPC")),
	RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
})
```

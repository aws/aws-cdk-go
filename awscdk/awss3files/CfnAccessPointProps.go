package awss3files


// Properties for defining a `CfnAccessPoint`.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//
//   vpc := ec2.NewVpc(this, jsii.String("Vpc"))
//
//   // Versioning is required — S3 Files relies on object versions for consistency.
//   bucket := s3.NewBucket(this, jsii.String("Bucket"), &BucketProps{
//   	Versioned: jsii.Boolean(true),
//   })
//
//   // S3 Files assumes this role to sync data between S3 and the file system.
//   role := iam.NewRole(this, jsii.String("S3FilesRole"), &RoleProps{
//   	AssumedBy: iam.NewServicePrincipal(jsii.String("elasticfilesystem.amazonaws.com")),
//   })
//
//   // S3 permissions: read/write access to the bucket and objects
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:ListBucket*"),
//   	},
//   	Resources: []*string{
//   		bucket.bucketArn,
//   	},
//   }))
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("s3:AbortMultipartUpload"),
//   		jsii.String("s3:DeleteObject"),
//   		jsii.String("s3:GetObject*"),
//   		jsii.String("s3:List*"),
//   		jsii.String("s3:PutObject*"),
//   	},
//   	Resources: []*string{
//   		bucket.ArnForObjects(jsii.String("*")),
//   	},
//   }))
//
//   // EventBridge permissions: S3 Files creates rules prefixed "DO-NOT-DELETE-S3-Files"
//   // to detect S3 object changes and trigger data synchronization.
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("events:DeleteRule"),
//   		jsii.String("events:DisableRule"),
//   		jsii.String("events:EnableRule"),
//   		jsii.String("events:PutRule"),
//   		jsii.String("events:PutTargets"),
//   		jsii.String("events:RemoveTargets"),
//   	},
//   	Resources: []*string{
//   		fmt.Sprintf("arn:%v:events:*:*:rule/DO-NOT-DELETE-S3-Files*", cdk.Aws_PARTITION()),
//   	},
//   	Conditions: map[string]interface{}{
//   		"StringEquals": map[string]*string{
//   			"events:ManagedBy": jsii.String("elasticfilesystem.amazonaws.com"),
//   		},
//   	},
//   }))
//   role.AddToPolicy(iam.NewPolicyStatement(&PolicyStatementProps{
//   	Actions: []*string{
//   		jsii.String("events:DescribeRule"),
//   		jsii.String("events:ListRuleNamesByTarget"),
//   		jsii.String("events:ListRules"),
//   		jsii.String("events:ListTargetsByRule"),
//   	},
//   	Resources: []*string{
//   		fmt.Sprintf("arn:%v:events:*:*:rule/*", cdk.Aws_PARTITION()),
//   	},
//   }))
//
//   fileSystem := s3files.NewCfnFileSystem(this, jsii.String("S3FilesFs"), &CfnFileSystemProps{
//   	Bucket: bucket.bucketArn,
//   	RoleArn: role.roleArn,
//   })
//
//   sg := ec2.NewSecurityGroup(this, jsii.String("MountTargetSG"), &SecurityGroupProps{
//   	Vpc: Vpc,
//   })
//
//   // Create a mount target in each private subnet so Lambda can reach the file system via NFS.
//   vpc.PrivateSubnets.forEach((subnet, i) =>
//     new s3files.CfnMountTarget(this, `MountTarget${i}`, {
//       fileSystemId: fileSystem.attrFileSystemId,
//       subnetId: subnet.subnetId,
//       securityGroups: [sg.securityGroupId],
//     }))
//
//   // The access point defines the POSIX identity and root path Lambda uses on the file system.
//   accessPoint := s3files.NewCfnAccessPoint(this, jsii.String("AccessPoint"), &CfnAccessPointProps{
//   	FileSystemId: fileSystem.attrFileSystemId,
//   	RootDirectory: &RootDirectoryProperty{
//   		Path: jsii.String("/export/lambda"),
//   		CreationPermissions: &CreationPermissionsProperty{
//   			OwnerGid: jsii.String("1001"),
//   			OwnerUid: jsii.String("1001"),
//   			Permissions: jsii.String("750"),
//   		},
//   	},
//   	PosixUser: &PosixUserProperty{
//   		Gid: jsii.String("1001"),
//   		Uid: jsii.String("1001"),
//   	},
//   })
//
//   fn := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_LATEST(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	Vpc: Vpc,
//   	Filesystem: lambda.FileSystem_FromS3FilesAccessPoint(accessPoint, jsii.String("/mnt/s3files")),
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html
//
type CfnAccessPointProps struct {
	// The ID of the S3 Files file system that the access point provides access to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// (optional) A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent creation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-clienttoken
	//
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-posixuser
	//
	PosixUser interface{} `field:"optional" json:"posixUser" yaml:"posixUser"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-rootdirectory
	//
	RootDirectory interface{} `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-s3files-accesspoint.html#cfn-s3files-accesspoint-tags
	//
	Tags *[]*CfnAccessPoint_AccessPointTagProperty `field:"optional" json:"tags" yaml:"tags"`
}


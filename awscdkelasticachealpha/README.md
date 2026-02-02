# ElastiCache CDK Construct Library

<!--BEGIN STABILITY BANNER-->---


![cdk-constructs: Experimental](https://img.shields.io/badge/cdk--constructs-experimental-important.svg?style=for-the-badge)

> The APIs of higher level constructs in this module are experimental and under active development.
> They are subject to non-backward compatible changes or removal in any future version. These are
> not subject to the [Semantic Versioning](https://semver.org/) model and breaking changes will be
> announced in the release notes. This means that while you may use them, you may need to update
> your source code when upgrading to a newer version of this package.

---
<!--END STABILITY BANNER-->

This module has constructs for [Amazon ElastiCache](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/WhatIs.html).

* The `ServerlessCache` construct facilitates the creation and management of serverless cache.
* The `User` and `UserGroup` constructs facilitate the creation and management of users for the cache.

## Serverless Cache

Amazon ElastiCache Serverless is a serverless option that automatically scales cache capacity based on application traffic patterns. You can create a serverless cache using the `ServerlessCache` construct:

```go
vpc := ec2.NewVpc(this, jsii.String("VPC"))

cache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Vpc: Vpc,
})
```

### Connecting to serverless cache

To control who can access the serverless cache by the security groups, use the `.connections` attribute.

The serverless cache has a default port `6379`.

This example allows an EC2 instance to connect to the serverless cache:

```go
var serverlessCache ServerlessCache
var instance Instance


// allow the EC2 instance to connect to serverless cache on default port 6379
serverlessCache.Connections.AllowDefaultPortFrom(instance)
```

### Cache usage limits

You can configure usage limits on both cache data storage and ECPU/second for your cache to control costs and ensure predictable performance.

**Configuration options:**

* **Maximum limits**: Ensure your cache usage never exceeds the configured maximum
* **Minimum limits**: Reserve a baseline level of resources for consistent performance
* **Both**: Define a range where your cache usage will operate

For more infomation, see [Setting scaling limits to manage costs](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Scaling.html#Pre-Scaling).

```go
var vpc Vpc


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Engine: elasticache.CacheEngine_VALKEY_LATEST,
	Vpc: Vpc,
	CacheUsageLimits: &CacheUsageLimitsProperty{
		// cache data storage limits (GB)
		DataStorageMinimumSize: awscdk.Size_Gibibytes(jsii.Number(2)),
		 // minimum: 1GB
		DataStorageMaximumSize: awscdk.Size_*Gibibytes(jsii.Number(3)),
		 // maximum: 5000GB
		// rate limits (ECPU/second)
		RequestRateLimitMinimum: jsii.Number(1000),
		 // minimum: 1000
		RequestRateLimitMaximum: jsii.Number(10000),
	},
})
```

### Backups and restore

You can enable automatic backups for serverless cache.
When automatic backups are enabled, ElastiCache creates a backup of the cache on a daily basis.

Also you can set the backup window for any time when it's most convenient.
If you don't specify a backup window, ElastiCache assigns one automatically.

For more information, see [Scheduling automatic backups](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/backups-automatic.html).

To enable automatic backups, set the `backupRetentionLimit` property. You can also specify the snapshot creation time by setting `backupTime` property:

```go
var vpc Vpc


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Backup: &BackupSettings{
		// enable automatic backups and set the retention period to 6 days
		BackupRetentionLimit: jsii.Number(6),
		// set the backup window to 9:00 AM UTC
		BackupTime: events.Schedule_Cron(&CronOptions{
			Hour: jsii.String("9"),
			Minute: jsii.String("0"),
		}),
	},
	Vpc: Vpc,
})
```

You can create a final backup by setting `backupNameBeforeDeletion` property.

```go
var vpc Vpc


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Engine: elasticache.CacheEngine_VALKEY_LATEST,
	Backup: &BackupSettings{
		// set a backup name before deleting a cache
		BackupNameBeforeDeletion: jsii.String("my-final-backup-name"),
	},
	Vpc: Vpc,
})
```

You can restore from backups by setting snapshot ARNs to `backupArnsToRestore` property:

```go
var vpc Vpc


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Engine: elasticache.CacheEngine_VALKEY_LATEST,
	Backup: &BackupSettings{
		// set the backup(s) to restore
		BackupArnsToRestore: []*string{
			jsii.String("arn:aws:elasticache:us-east-1:123456789012:serverlesscachesnapshot:my-final-backup-name"),
		},
	},
	Vpc: Vpc,
})
```

### Encryption at rest

At-rest encryption is always enabled for Serverless Cache. There are two encryption options:

* **Default**: When no `kmsKey` is specified (left as `undefined`), AWS owned KMS keys are used automatically
* **Customer Managed Key**: Create a KMS key first, then pass it to the cache via the `kmsKey` property

### Customer Managed Key for encryption at rest

ElastiCache supports symmetric Customer Managed key (CMK) for encryption at rest.

For more information, see [Using customer managed keys from AWS KMS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/at-rest-encryption.html#using-customer-managed-keys-for-elasticache-security).

To use CMK, set your CMK to the `kmsKey` property:

```go
import "github.com/aws/aws-cdk-go/awscdk"

var kmsKey Key
var vpc Vpc


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Engine: elasticache.CacheEngine_VALKEY_LATEST,
	ServerlessCacheName: jsii.String("my-serverless-cache"),
	Vpc: Vpc,
	// set Customer Managed Key
	KmsKey: KmsKey,
})
```

### Metrics and monitoring

You can monitor your serverless cache using CloudWatch Metrics via the `metric` method.

For more information about serverless cache metrics, see [Serverless metrics and events for Valkey and Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/serverless-metrics-events-redis.html) and [Serverless metrics and events for Memcached](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/serverless-metrics-events.memcached.html).

```go
var serverlessCache ServerlessCache


// The 5 minutes average of the total number of successful read-only key lookups in the cache.
cacheHits := serverlessCache.MetricCacheHitCount()

// The 5 minutes average of the total number of bytes used by the data stored in the cache.
bytesUsedForCache := serverlessCache.MetricDataStored()

// The 5 minutes average of the total number of ElastiCacheProcessingUnits (ECPUs) consumed by the requests executed on the cache.
elastiCacheProcessingUnits := serverlessCache.MetricProcessingUnitsConsumed()

// Create an alarm for ECPUs.
elastiCacheProcessingUnits.CreateAlarm(this, jsii.String("ElastiCacheProcessingUnitsAlarm"), &CreateAlarmOptions{
	Threshold: jsii.Number(50),
	EvaluationPeriods: jsii.Number(1),
})
```

### Import an existing serverless cache

To import an existing ServerlessCache, use the `ServerlessCache.fromServerlessCacheAttributes` method:

```go
var securityGroup SecurityGroup


importedServerlessCache := elasticache.ServerlessCache_FromServerlessCacheAttributes(this, jsii.String("ImportedServerlessCache"), &ServerlessCacheAttributes{
	ServerlessCacheName: jsii.String("my-serverless-cache"),
	SecurityGroups: []ISecurityGroup{
		securityGroup,
	},
})
```

## User and User Group

Setup required properties and create:

```go
newDefaultUser := elasticache.NewNoPasswordUser(this, jsii.String("NoPasswordUser"), &NoPasswordUserProps{
	UserId: jsii.String("default"),
	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
})

userGroup := elasticache.NewUserGroup(this, jsii.String("UserGroup"), &UserGroupProps{
	Users: []IUser{
		newDefaultUser,
	},
})
```

### RBAC

In Valkey 7.2 and onward and Redis OSS 6.0 onward you can use a feature called Role-Based Access Control (RBAC). RBAC is also the only way to control access to serverless caches.

RBAC enables you to control cache access through user groups. These user groups are designed as a way to organize access to caches.

For more information, see [Role-Based Access Control (RBAC)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html).

To enable RBAC for ElastiCache with Valkey or Redis OSS, you take the following steps:

* Create users.
* Create a user group and add users to the user group.
* Assign the user group to a cache.

### Create users

First, you need to create users by using `IamUser`, `PasswordUser` or `NoPasswordUser` construct.

With RBAC, you create users and assign them specific permissions by using `accessString` property.

For more information, see [Specifying Permissions Using an Access String](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html#Access-string).

You can create an IAM-enabled user by using `IamUser` construct:

```go
user := elasticache.NewIamUser(this, jsii.String("User"), &IamUserProps{
	// set user engine
	Engine: elasticache.UserEngine_REDIS,

	// set user id
	UserId: jsii.String("my-user"),

	// set username
	UserName: jsii.String("my-user"),

	// set access string
	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
})
```

> NOTE: IAM-enabled users must have matching user id and username. For more information, see [Limitations](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/auth-iam.html). The construct can set automatically the username to be the same as the user id.

If you want to create a password authenticated user, use `PasswordUser` construct:

```go
user := elasticache.NewPasswordUser(this, jsii.String("User"), &PasswordUserProps{
	// set user engine
	Engine: elasticache.UserEngine_VALKEY,

	// set user id
	UserId: jsii.String("my-user-id"),

	// set access string
	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),

	// set username
	UserName: jsii.String("my-user-name"),

	// set up to two passwords
	Passwords: []SecretValue{
		awscdk.SecretValue_SecretsManager(jsii.String("SecretIdForPassword")),
		awscdk.SecretValue_*SecretsManager(jsii.String("AnotherSecretIdForPassword")),
	},
})
```

You can also create a no password required user by using `NoPasswordUser` construct:

```go
user := elasticache.NewNoPasswordUser(this, jsii.String("User"), &NoPasswordUserProps{
	// set user engine
	Engine: elasticache.UserEngine_REDIS,

	// set user id
	UserId: jsii.String("my-user-id"),

	// set access string
	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),

	// set username
	UserName: jsii.String("my-user-name"),
})
```

> NOTE: `NoPasswordUser` is only available for Redis Cache.

### Default user

ElastiCache automatically creates a default user with both a user ID and username set to `default`. This default user cannot be modified or deleted. The user is created as a no password authentication user.

This user is intended for compatibility with the default behavior of previous Redis OSS versions and has an access string that permits it to call all commands and access all keys.

To use this automatically created default user in CDK, you can import it using `NoPasswordUser.fromUserAttributes` method. For more information on import methods, see the [Import an existing user and user group](#import-an-existing-user-and-user-group) section.

To add proper access control to a cache, replace the default user with a new one that is either disabled by setting the `accessString` to `off -@all` or secured with a strong password.

To change the default user, create a new default user with the username set to `default`. You can then swap it with the original default user.

For more information, see [Applying RBAC to a Cache for ElastiCache with Valkey or Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/Clusters.RBAC.html#rbac-using).

If you want to create a new default user, `userName` must be `default` and `userId` must not be `default` by using `NoPasswordUser` or `PasswordUser`:

```go
// use the original `default` user by using import method
defaultUser := elasticache.NoPasswordUser_FromUserAttributes(this, jsii.String("DefaultUser"), &UserBaseAttributes{
	// userId and userName must be 'default'
	UserId: jsii.String("default"),
})

// create a new default user
newDefaultUser := elasticache.NewNoPasswordUser(this, jsii.String("NewDefaultUser"), &NoPasswordUserProps{
	// new default user id must not be 'default'
	UserId: jsii.String("new-default"),
	// new default username must be 'default'
	UserName: jsii.String("default"),
	// set access string
	AccessControl: elasticache.AccessControl_FromAccessString(jsii.String("on ~* +@all")),
})
```

> NOTE: You can't create a new default user using `IamUser` because an IAM-enabled user's username and user ID cannot be different.

### Add users to the user group

Next, use the `UserGroup` construct to create a user group and add users to it.
Ensure that you include either the original default user or a new default user:

```go
var newDefaultUser IUser
var user IUser
var anotherUser IUser


userGroup := elasticache.NewUserGroup(this, jsii.String("UserGroup"), &UserGroupProps{
	// add users including default user
	Users: []IUser{
		newDefaultUser,
		user,
	},
})

// you can also add a user by using addUser method
userGroup.AddUser(anotherUser)
```

### Assign user group

Finally, assign a user group to cache:

```go
var vpc Vpc
var userGroup UserGroup


serverlessCache := elasticache.NewServerlessCache(this, jsii.String("ServerlessCache"), &ServerlessCacheProps{
	Engine: elasticache.CacheEngine_VALKEY_LATEST,
	ServerlessCacheName: jsii.String("my-serverless-cache"),
	Vpc: Vpc,
	// assign User Group
	UserGroup: UserGroup,
})
```

### Grant permissions to IAM-enabled users

If you create IAM-enabled users, `"elasticache:Connect"` action must be allowed for the users and cache.

> NOTE: You don't need grant permissions to no password required users or password authentication users.

For more information, see [Authenticating with IAM](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/auth-iam.html).

To grant permissions, you can use the `grantConnect` method in `IamUser` and `ServerlessCache` constructs:

```go
var user IamUser
var serverlessCache ServerlessCache
var role Role


// grant "elasticache:Connect" action permissions to role
user.GrantConnect(role)
serverlessCache.Grants.Connect(role)
```

### Import an existing user and user group

You can import an existing user and user group by using import methods:

```go
stack := awscdk.Newstack()

importedIamUser := elasticache.IamUser_FromUserId(this, jsii.String("ImportedIamUser"), jsii.String("my-iam-user-id"))

importedPasswordUser := elasticache.PasswordUser_FromUserAttributes(stack, jsii.String("ImportedPasswordUser"), &UserBaseAttributes{
	UserId: jsii.String("my-password-user-id"),
})

importedNoPasswordUser := elasticache.NoPasswordUser_FromUserAttributes(stack, jsii.String("ImportedNoPasswordUser"), &UserBaseAttributes{
	UserId: jsii.String("my-no-password-user-id"),
})

importedUserGroup := elasticache.UserGroup_FromUserGroupAttributes(this, jsii.String("ImportedUserGroup"), &UserGroupAttributes{
	UserGroupName: jsii.String("my-user-group-name"),
})
```

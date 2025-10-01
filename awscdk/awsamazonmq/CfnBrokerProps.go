package awsamazonmq


// Properties for defining a `CfnBroker`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBrokerProps := &CfnBrokerProps{
//   	BrokerName: jsii.String("brokerName"),
//   	DeploymentMode: jsii.String("deploymentMode"),
//   	EngineType: jsii.String("engineType"),
//   	HostInstanceType: jsii.String("hostInstanceType"),
//   	PubliclyAccessible: jsii.Boolean(false),
//
//   	// the properties below are optional
//   	AuthenticationStrategy: jsii.String("authenticationStrategy"),
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	Configuration: &ConfigurationIdProperty{
//   		Id: jsii.String("id"),
//   		Revision: jsii.Number(123),
//   	},
//   	DataReplicationMode: jsii.String("dataReplicationMode"),
//   	DataReplicationPrimaryBrokerArn: jsii.String("dataReplicationPrimaryBrokerArn"),
//   	EncryptionOptions: &EncryptionOptionsProperty{
//   		UseAwsOwnedKey: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	EngineVersion: jsii.String("engineVersion"),
//   	LdapServerMetadata: &LdapServerMetadataProperty{
//   		Hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		RoleBase: jsii.String("roleBase"),
//   		RoleSearchMatching: jsii.String("roleSearchMatching"),
//   		ServiceAccountUsername: jsii.String("serviceAccountUsername"),
//   		UserBase: jsii.String("userBase"),
//   		UserSearchMatching: jsii.String("userSearchMatching"),
//
//   		// the properties below are optional
//   		RoleName: jsii.String("roleName"),
//   		RoleSearchSubtree: jsii.Boolean(false),
//   		ServiceAccountPassword: jsii.String("serviceAccountPassword"),
//   		UserRoleName: jsii.String("userRoleName"),
//   		UserSearchSubtree: jsii.Boolean(false),
//   	},
//   	Logs: &LogListProperty{
//   		Audit: jsii.Boolean(false),
//   		General: jsii.Boolean(false),
//   	},
//   	MaintenanceWindowStartTime: &MaintenanceWindowProperty{
//   		DayOfWeek: jsii.String("dayOfWeek"),
//   		TimeOfDay: jsii.String("timeOfDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   	SecurityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	StorageType: jsii.String("storageType"),
//   	SubnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	Tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Users: []interface{}{
//   		&UserProperty{
//   			Password: jsii.String("password"),
//   			Username: jsii.String("username"),
//
//   			// the properties below are optional
//   			ConsoleAccess: jsii.Boolean(false),
//   			Groups: []*string{
//   				jsii.String("groups"),
//   			},
//   			ReplicationUser: jsii.Boolean(false),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html
//
type CfnBrokerProps struct {
	// The name of the broker.
	//
	// This value must be unique in your AWS account , 1-50 characters long, must contain only letters, numbers, dashes, and underscores, and must not contain white spaces, brackets, wildcard characters, or special characters.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names. Broker names are accessible to other AWS services, including C CloudWatch Logs . Broker names are not intended to be used for private or sensitive data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-brokername
	//
	BrokerName *string `field:"required" json:"brokerName" yaml:"brokerName"`
	// The deployment mode of the broker. Available values:.
	//
	// - `SINGLE_INSTANCE`
	// - `ACTIVE_STANDBY_MULTI_AZ`
	// - `CLUSTER_MULTI_AZ`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-deploymentmode
	//
	DeploymentMode *string `field:"required" json:"deploymentMode" yaml:"deploymentMode"`
	// The type of broker engine.
	//
	// Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-enginetype
	//
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The broker's instance type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-hostinstancetype
	//
	HostInstanceType *string `field:"required" json:"hostInstanceType" yaml:"hostInstanceType"`
	// Enables connections from applications outside of the VPC that hosts the broker's subnets.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-publiclyaccessible
	//
	PubliclyAccessible interface{} `field:"required" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Optional.
	//
	// The authentication strategy used to secure the broker. The default is `SIMPLE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-authenticationstrategy
	//
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// Enables automatic upgrades to new minor versions for brokers, as new broker engine versions are released and supported by Amazon MQ.
	//
	// Automatic upgrades occur during the scheduled maintenance window of the broker or after a manual broker reboot.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-autominorversionupgrade
	//
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// A list of information about the configuration.
	//
	// Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Defines whether this broker is a part of a data replication pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-datareplicationmode
	//
	DataReplicationMode *string `field:"optional" json:"dataReplicationMode" yaml:"dataReplicationMode"`
	// The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker.
	//
	// Must be set when dataReplicationMode is set to CRDR.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-datareplicationprimarybrokerarn
	//
	DataReplicationPrimaryBrokerArn *string `field:"optional" json:"dataReplicationPrimaryBrokerArn" yaml:"dataReplicationPrimaryBrokerArn"`
	// Encryption options for the broker.
	//
	// Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-encryptionoptions
	//
	EncryptionOptions interface{} `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [Engine](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) in the *Amazon MQ Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-engineversion
	//
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Optional.
	//
	// The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-ldapservermetadata
	//
	LdapServerMetadata interface{} `field:"optional" json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
	// Enables Amazon CloudWatch logging for brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-logs
	//
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-maintenancewindowstarttime
	//
	MaintenanceWindowStartTime interface{} `field:"optional" json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-securitygroups
	//
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The broker's storage type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-storagetype
	//
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.
	//
	// If you specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create VPC endpoints for your broker with multiple subnets in the same Availability Zone. A SINGLE_INSTANCE deployment requires one subnet (for example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ deployment (ACTIVEMQ) requires two subnets. A CLUSTER_MULTI_AZ deployment (RABBITMQ) has no subnet requirements when deployed with public accessibility, deployment without public accessibility requires at least one subnet.
	//
	// > If you specify subnets in a shared VPC for a RabbitMQ broker, the associated VPC to which the specified subnets belong must be owned by your AWS account . Amazon MQ will not be able to create VPC enpoints in VPCs that are not owned by your AWS account .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-subnetids
	//
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// An array of key-value pairs.
	//
	// For more information, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the *Billing and Cost Management User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-tags
	//
	Tags *[]*CfnBroker_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
	// The list of broker users (persons or applications) who can access queues and topics.
	//
	// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent RabbitMQ users are created by via the RabbitMQ web console or by using the RabbitMQ management API.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-amazonmq-broker.html#cfn-amazonmq-broker-users
	//
	Users interface{} `field:"optional" json:"users" yaml:"users"`
}


package awsamazonmq


// Properties for defining a `CfnBroker`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnBrokerProps := &cfnBrokerProps{
//   	autoMinorVersionUpgrade: jsii.Boolean(false),
//   	brokerName: jsii.String("brokerName"),
//   	deploymentMode: jsii.String("deploymentMode"),
//   	engineType: jsii.String("engineType"),
//   	engineVersion: jsii.String("engineVersion"),
//   	hostInstanceType: jsii.String("hostInstanceType"),
//   	publiclyAccessible: jsii.Boolean(false),
//   	users: []interface{}{
//   		&userProperty{
//   			password: jsii.String("password"),
//   			username: jsii.String("username"),
//
//   			// the properties below are optional
//   			consoleAccess: jsii.Boolean(false),
//   			groups: []*string{
//   				jsii.String("groups"),
//   			},
//   		},
//   	},
//
//   	// the properties below are optional
//   	authenticationStrategy: jsii.String("authenticationStrategy"),
//   	configuration: &configurationIdProperty{
//   		id: jsii.String("id"),
//   		revision: jsii.Number(123),
//   	},
//   	encryptionOptions: &encryptionOptionsProperty{
//   		useAwsOwnedKey: jsii.Boolean(false),
//
//   		// the properties below are optional
//   		kmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	ldapServerMetadata: &ldapServerMetadataProperty{
//   		hosts: []*string{
//   			jsii.String("hosts"),
//   		},
//   		roleBase: jsii.String("roleBase"),
//   		roleSearchMatching: jsii.String("roleSearchMatching"),
//   		serviceAccountPassword: jsii.String("serviceAccountPassword"),
//   		serviceAccountUsername: jsii.String("serviceAccountUsername"),
//   		userBase: jsii.String("userBase"),
//   		userSearchMatching: jsii.String("userSearchMatching"),
//
//   		// the properties below are optional
//   		roleName: jsii.String("roleName"),
//   		roleSearchSubtree: jsii.Boolean(false),
//   		userRoleName: jsii.String("userRoleName"),
//   		userSearchSubtree: jsii.Boolean(false),
//   	},
//   	logs: &logListProperty{
//   		audit: jsii.Boolean(false),
//   		general: jsii.Boolean(false),
//   	},
//   	maintenanceWindowStartTime: &maintenanceWindowProperty{
//   		dayOfWeek: jsii.String("dayOfWeek"),
//   		timeOfDay: jsii.String("timeOfDay"),
//   		timeZone: jsii.String("timeZone"),
//   	},
//   	securityGroups: []*string{
//   		jsii.String("securityGroups"),
//   	},
//   	storageType: jsii.String("storageType"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []tagsEntryProperty{
//   		&tagsEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnBrokerProps struct {
	// Enables automatic upgrades to new minor versions for brokers, as new broker engine versions are released and supported by Amazon MQ.
	//
	// Automatic upgrades occur during the scheduled maintenance window of the broker or after a manual broker reboot.
	AutoMinorVersionUpgrade interface{} `field:"required" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The name of the broker.
	//
	// This value must be unique in your AWS account , 1-50 characters long, must contain only letters, numbers, dashes, and underscores, and must not contain white spaces, brackets, wildcard characters, or special characters.
	//
	// > Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names. Broker names are accessible to other AWS services, including C CloudWatch Logs . Broker names are not intended to be used for private or sensitive data.
	BrokerName *string `field:"required" json:"brokerName" yaml:"brokerName"`
	// The deployment mode of the broker. Available values:.
	//
	// - `SINGLE_INSTANCE`
	// - `ACTIVE_STANDBY_MULTI_AZ`
	// - `CLUSTER_MULTI_AZ`.
	DeploymentMode *string `field:"required" json:"deploymentMode" yaml:"deploymentMode"`
	// The type of broker engine.
	//
	// Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ` .
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// The version of the broker engine.
	//
	// For a list of supported engine versions, see [Engine](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-engine.html) in the *Amazon MQ Developer Guide* .
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// The broker's instance type.
	HostInstanceType *string `field:"required" json:"hostInstanceType" yaml:"hostInstanceType"`
	// Enables connections from applications outside of the VPC that hosts the broker's subnets.
	PubliclyAccessible interface{} `field:"required" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// The list of broker users (persons or applications) who can access queues and topics.
	//
	// For Amazon MQ for RabbitMQ brokers, one and only one administrative user is accepted and created when a broker is first provisioned. All subsequent RabbitMQ users are created by via the RabbitMQ web console or by using the RabbitMQ management API.
	Users interface{} `field:"required" json:"users" yaml:"users"`
	// Optional.
	//
	// The authentication strategy used to secure the broker. The default is `SIMPLE` .
	AuthenticationStrategy *string `field:"optional" json:"authenticationStrategy" yaml:"authenticationStrategy"`
	// A list of information about the configuration.
	//
	// Does not apply to RabbitMQ brokers.
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Encryption options for the broker.
	//
	// Does not apply to RabbitMQ brokers.
	EncryptionOptions interface{} `field:"optional" json:"encryptionOptions" yaml:"encryptionOptions"`
	// Optional.
	//
	// The metadata of the LDAP server used to authenticate and authorize connections to the broker. Does not apply to RabbitMQ brokers.
	LdapServerMetadata interface{} `field:"optional" json:"ldapServerMetadata" yaml:"ldapServerMetadata"`
	// Enables Amazon CloudWatch logging for brokers.
	Logs interface{} `field:"optional" json:"logs" yaml:"logs"`
	// The scheduled time period relative to UTC during which Amazon MQ begins to apply pending updates or patches to the broker.
	MaintenanceWindowStartTime interface{} `field:"optional" json:"maintenanceWindowStartTime" yaml:"maintenanceWindowStartTime"`
	// The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// The broker's storage type.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.
	//
	// If you specify more than one subnet, the subnets must be in different Availability Zones. Amazon MQ will not be able to create VPC endpoints for your broker with multiple subnets in the same Availability Zone. A SINGLE_INSTANCE deployment requires one subnet (for example, the default subnet). An ACTIVE_STANDBY_MULTI_AZ deployment (ACTIVEMQ) requires two subnets. A CLUSTER_MULTI_AZ deployment (RABBITMQ) has no subnet requirements when deployed with public accessibility, deployment without public accessibility requires at least one subnet.
	//
	// > If you specify subnets in a shared VPC for a RabbitMQ broker, the associated VPC to which the specified subnets belong must be owned by your AWS account . Amazon MQ will not be able to create VPC enpoints in VPCs that are not owned by your AWS account .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// An array of key-value pairs.
	//
	// For more information, see [Using Cost Allocation Tags](https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html) in the *Billing and Cost Management User Guide* .
	Tags *[]*CfnBroker_TagsEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}


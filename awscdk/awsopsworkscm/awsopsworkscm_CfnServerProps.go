package awsopsworkscm

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnServer`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnServerProps := &cfnServerProps{
//   	instanceProfileArn: jsii.String("instanceProfileArn"),
//   	instanceType: jsii.String("instanceType"),
//   	serviceRoleArn: jsii.String("serviceRoleArn"),
//
//   	// the properties below are optional
//   	associatePublicIpAddress: jsii.Boolean(false),
//   	backupId: jsii.String("backupId"),
//   	backupRetentionCount: jsii.Number(123),
//   	customCertificate: jsii.String("customCertificate"),
//   	customDomain: jsii.String("customDomain"),
//   	customPrivateKey: jsii.String("customPrivateKey"),
//   	disableAutomatedBackup: jsii.Boolean(false),
//   	engine: jsii.String("engine"),
//   	engineAttributes: []interface{}{
//   		&engineAttributeProperty{
//   			name: jsii.String("name"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	engineModel: jsii.String("engineModel"),
//   	engineVersion: jsii.String("engineVersion"),
//   	keyPair: jsii.String("keyPair"),
//   	preferredBackupWindow: jsii.String("preferredBackupWindow"),
//   	preferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	securityGroupIds: []*string{
//   		jsii.String("securityGroupIds"),
//   	},
//   	serverName: jsii.String("serverName"),
//   	subnetIds: []*string{
//   		jsii.String("subnetIds"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnServerProps struct {
	// The ARN of the instance profile that your Amazon EC2 instances use.
	InstanceProfileArn *string `field:"required" json:"instanceProfileArn" yaml:"instanceProfileArn"`
	// The Amazon EC2 instance type to use.
	//
	// For example, `m5.large` .
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// The service role that the AWS OpsWorks CM service backend uses to work with your account.
	//
	// Although the AWS OpsWorks management console typically creates the service role for you, if you are using the AWS CLI or API commands, run the service-role-creation.yaml AWS CloudFormation template, located at https://s3.amazonaws.com/opsworks-cm-us-east-1-prod-default-assets/misc/opsworks-cm-roles.yaml. This template creates a CloudFormation stack that includes the service role and instance profile that you need.
	ServiceRoleArn *string `field:"required" json:"serviceRoleArn" yaml:"serviceRoleArn"`
	// Associate a public IP address with a server that you are launching.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// If you specify this field, AWS OpsWorks CM creates the server by using the backup represented by BackupId.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The number of automated backups that you want to keep.
	//
	// Whenever a new backup is created, AWS OpsWorks CM deletes the oldest backups if this number is exceeded. The default value is `1` .
	BackupRetentionCount *float64 `field:"optional" json:"backupRetentionCount" yaml:"backupRetentionCount"`
	// Supported on servers running Chef Automate 2.0 only. A PEM-formatted HTTPS certificate. The value can be be a single, self-signed certificate, or a certificate chain. If you specify a custom certificate, you must also specify values for `CustomDomain` and `CustomPrivateKey` . The following are requirements for the `CustomCertificate` value:.
	//
	// - You can provide either a self-signed, custom certificate, or the full certificate chain.
	// - The certificate must be a valid X509 certificate, or a certificate chain in PEM format.
	// - The certificate must be valid at the time of upload. A certificate can't be used before its validity period begins (the certificate's `NotBefore` date), or after it expires (the certificate's `NotAfter` date).
	// - The certificate’s common name or subject alternative names (SANs), if present, must match the value of `CustomDomain` .
	// - The certificate must match the value of `CustomPrivateKey` .
	CustomCertificate *string `field:"optional" json:"customCertificate" yaml:"customCertificate"`
	// Supported on servers running Chef Automate 2.0 only. An optional public endpoint of a server, such as `https://aws.my-company.com` . To access the server, create a CNAME DNS record in your preferred DNS service that points the custom domain to the endpoint that is generated when the server is created (the value of the CreateServer Endpoint attribute). You cannot access the server by using the generated `Endpoint` value if the server is using a custom domain. If you specify a custom domain, you must also specify values for `CustomCertificate` and `CustomPrivateKey` .
	CustomDomain *string `field:"optional" json:"customDomain" yaml:"customDomain"`
	// Supported on servers running Chef Automate 2.0 only. A private key in PEM format for connecting to the server by using HTTPS. The private key must not be encrypted; it cannot be protected by a password or passphrase. If you specify a custom private key, you must also specify values for `CustomDomain` and `CustomCertificate` .
	CustomPrivateKey *string `field:"optional" json:"customPrivateKey" yaml:"customPrivateKey"`
	// Enable or disable scheduled backups.
	//
	// Valid values are `true` or `false` . The default value is `true` .
	DisableAutomatedBackup interface{} `field:"optional" json:"disableAutomatedBackup" yaml:"disableAutomatedBackup"`
	// The configuration management engine to use.
	//
	// Valid values include `ChefAutomate` and `Puppet` .
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Optional engine attributes on a specified server.
	//
	// **Attributes accepted in a Chef createServer request:** - `CHEF_AUTOMATE_PIVOTAL_KEY` : A base64-encoded RSA public key. The corresponding private key is required to access the Chef API. When no CHEF_AUTOMATE_PIVOTAL_KEY is set, a private key is generated and returned in the response. When you are specifying the value of CHEF_AUTOMATE_PIVOTAL_KEY as a parameter in the AWS CloudFormation console, you must add newline ( `\n` ) characters at the end of each line of the pivotal key value.
	// - `CHEF_AUTOMATE_ADMIN_PASSWORD` : The password for the administrative user in the Chef Automate web-based dashboard. The password length is a minimum of eight characters, and a maximum of 32. The password can contain letters, numbers, and special characters (!/@#$%^&+=_). The password must contain at least one lower case letter, one upper case letter, one number, and one special character. When no CHEF_AUTOMATE_ADMIN_PASSWORD is set, one is generated and returned in the response.
	//
	// **Attributes accepted in a Puppet createServer request:** - `PUPPET_ADMIN_PASSWORD` : To work with the Puppet Enterprise console, a password must use ASCII characters.
	// - `PUPPET_R10K_REMOTE` : The r10k remote is the URL of your control repository (for example, ssh://git@your.git-repo.com:user/control-repo.git). Specifying an r10k remote opens TCP port 8170.
	// - `PUPPET_R10K_PRIVATE_KEY` : If you are using a private Git repository, add PUPPET_R10K_PRIVATE_KEY to specify a PEM-encoded private SSH key.
	EngineAttributes interface{} `field:"optional" json:"engineAttributes" yaml:"engineAttributes"`
	// The engine model of the server.
	//
	// Valid values in this release include `Monolithic` for Puppet and `Single` for Chef.
	EngineModel *string `field:"optional" json:"engineModel" yaml:"engineModel"`
	// The major release version of the engine that you want to use.
	//
	// For a Chef server, the valid value for EngineVersion is currently `2` . For a Puppet server, valid values are `2019` or `2017` .
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// The Amazon EC2 key pair to set for the instance.
	//
	// This parameter is optional; if desired, you may specify this parameter to connect to your instances by using SSH.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// The start time for a one-hour period during which AWS OpsWorks CM backs up application-level data on your server if automated backups are enabled.
	//
	// Valid values must be specified in one of the following formats:
	//
	// - `HH:MM` for daily backups
	// - `DDD:HH:MM` for weekly backups
	//
	// `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random, daily start time.
	//
	// *Example:* `08:00` , which represents a daily start time of 08:00 UTC.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// The start time for a one-hour period each week during which AWS OpsWorks CM performs maintenance on the instance.
	//
	// Valid values must be specified in the following format: `DDD:HH:MM` . `MM` must be specified as `00` . The specified time is in coordinated universal time (UTC). The default value is a random one-hour period on Tuesday, Wednesday, or Friday. See `TimeWindowDefinition` for more information.
	//
	// *Example:* `Mon:08:00` , which represents a start time of every Monday at 08:00 UTC. (8:00 a.m.)
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// A list of security group IDs to attach to the Amazon EC2 instance.
	//
	// If you add this parameter, the specified security groups must be within the VPC that is specified by `SubnetIds` .
	//
	// If you do not specify this parameter, AWS OpsWorks CM creates one new security group that uses TCP ports 22 and 443, open to 0.0.0.0/0 (everyone).
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The name of the server.
	//
	// The server name must be unique within your AWS account, within each region. Server names must start with a letter; then letters, numbers, or hyphens (-) are allowed, up to a maximum of 40 characters.
	ServerName *string `field:"optional" json:"serverName" yaml:"serverName"`
	// The IDs of subnets in which to launch the server EC2 instance.
	//
	// Amazon EC2-Classic customers: This field is required. All servers must run within a VPC. The VPC must have "Auto Assign Public IP" enabled.
	//
	// EC2-VPC customers: This field is optional. If you do not specify subnet IDs, your EC2 instances are created in a default subnet that is selected by Amazon EC2. If you specify subnet IDs, the VPC must have "Auto Assign Public IP" enabled.
	//
	// For more information about supported Amazon EC2 platforms, see [Supported Platforms](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-platforms.html) .
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// A map that contains tag keys and tag values to attach to an AWS OpsWorks for Chef Automate or OpsWorks for Puppet Enterprise server.
	//
	// - The key cannot be empty.
	// - The key can be a maximum of 127 characters, and can contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - The value can be a maximum 255 characters, and contain only Unicode letters, numbers, or separators, or the following special characters: `+ - = . _ : / @`
	// - Leading and trailing spaces are trimmed from both the key and value.
	// - A maximum of 50 user-applied tags is allowed for any AWS OpsWorks CM server.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


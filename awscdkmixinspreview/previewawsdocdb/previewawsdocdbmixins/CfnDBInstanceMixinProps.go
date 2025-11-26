package previewawsdocdbmixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnDBInstancePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnDBInstanceMixinProps := &CfnDBInstanceMixinProps{
//   	AutoMinorVersionUpgrade: jsii.Boolean(false),
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	CaCertificateIdentifier: jsii.String("caCertificateIdentifier"),
//   	CertificateRotationRestart: jsii.Boolean(false),
//   	DbClusterIdentifier: jsii.String("dbClusterIdentifier"),
//   	DbInstanceClass: jsii.String("dbInstanceClass"),
//   	DbInstanceIdentifier: jsii.String("dbInstanceIdentifier"),
//   	EnablePerformanceInsights: jsii.Boolean(false),
//   	PreferredMaintenanceWindow: jsii.String("preferredMaintenanceWindow"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html
//
type CfnDBInstanceMixinProps struct {
	// This parameter does not apply to Amazon DocumentDB.
	//
	// Amazon DocumentDB does not perform minor version upgrades regardless of the value set.
	//
	// Default: `false`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-autominorversionupgrade
	//
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// The Amazon EC2 Availability Zone that the instance is created in.
	//
	// Default: A random, system-chosen Availability Zone in the endpoint's AWS Region .
	//
	// Example: `us-east-1d`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The identifier of the CA certificate for this DB instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-cacertificateidentifier
	//
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// Specifies whether the DB instance is restarted when you rotate your SSL/TLS certificate.
	//
	// By default, the DB instance is restarted when you rotate your SSL/TLS certificate. The certificate is not updated until the DB instance is restarted.
	//
	// > Set this parameter only if you are *not* using SSL/TLS to connect to the DB instance.
	//
	// If you are using SSL/TLS to connect to the DB instance, see [Updating Your Amazon DocumentDB TLS Certificates](https://docs.aws.amazon.com/documentdb/latest/developerguide/ca_cert_rotation.html) and [Encrypting Data in Transit](https://docs.aws.amazon.com/documentdb/latest/developerguide/security.encryption.ssl.html) in the *Amazon DocumentDB Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-certificaterotationrestart
	//
	CertificateRotationRestart interface{} `field:"optional" json:"certificateRotationRestart" yaml:"certificateRotationRestart"`
	// The identifier of the cluster that the instance will belong to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbclusteridentifier
	//
	DbClusterIdentifier *string `field:"optional" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The compute and memory capacity of the instance;
	//
	// for example, `db.m4.large` . If you change the class of an instance there can be some interruption in the cluster's service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbinstanceclass
	//
	DbInstanceClass *string `field:"optional" json:"dbInstanceClass" yaml:"dbInstanceClass"`
	// The instance identifier. This parameter is stored as a lowercase string.
	//
	// Constraints:
	//
	// - Must contain from 1 to 63 letters, numbers, or hyphens.
	// - The first character must be a letter.
	// - Cannot end with a hyphen or contain two consecutive hyphens.
	//
	// Example: `mydbinstance`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbinstanceidentifier
	//
	DbInstanceIdentifier *string `field:"optional" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// A value that indicates whether to enable Performance Insights for the DB Instance.
	//
	// For more information, see [Using Amazon Performance Insights](https://docs.aws.amazon.com/documentdb/latest/developerguide/performance-insights.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-enableperformanceinsights
	//
	EnablePerformanceInsights interface{} `field:"optional" json:"enablePerformanceInsights" yaml:"enablePerformanceInsights"`
	// The time range each week during which system maintenance can occur, in Universal Coordinated Time (UTC).
	//
	// Format: `ddd:hh24:mi-ddd:hh24:mi`
	//
	// The default is a 30-minute window selected at random from an 8-hour block of time for each AWS Region , occurring on a random day of the week.
	//
	// Valid days: Mon, Tue, Wed, Thu, Fri, Sat, Sun
	//
	// Constraints: Minimum 30-minute window.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-preferredmaintenancewindow
	//
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// The tags to be assigned to the instance.
	//
	// You can assign up to 10 tags to an instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


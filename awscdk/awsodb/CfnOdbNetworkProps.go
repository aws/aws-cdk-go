package awsodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnOdbNetwork`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnOdbNetworkProps := &CfnOdbNetworkProps{
//   	AvailabilityZone: jsii.String("availabilityZone"),
//   	AvailabilityZoneId: jsii.String("availabilityZoneId"),
//   	BackupSubnetCidr: jsii.String("backupSubnetCidr"),
//   	ClientSubnetCidr: jsii.String("clientSubnetCidr"),
//   	CustomDomainName: jsii.String("customDomainName"),
//   	DefaultDnsPrefix: jsii.String("defaultDnsPrefix"),
//   	DeleteAssociatedResources: jsii.Boolean(false),
//   	DisplayName: jsii.String("displayName"),
//   	S3Access: jsii.String("s3Access"),
//   	S3PolicyDocument: jsii.String("s3PolicyDocument"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ZeroEtlAccess: jsii.String("zeroEtlAccess"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html
//
type CfnOdbNetworkProps struct {
	// The Availability Zone (AZ) where the ODB network is located.
	//
	// Required when creating an ODB network. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-availabilityzone
	//
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The AZ ID of the AZ where the ODB network is located.
	//
	// Required when creating an ODB network. Specify either AvailabilityZone or AvailabilityZoneId to define the location of the network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-availabilityzoneid
	//
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The CIDR range of the backup subnet in the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-backupsubnetcidr
	//
	BackupSubnetCidr *string `field:"optional" json:"backupSubnetCidr" yaml:"backupSubnetCidr"`
	// The CIDR range of the client subnet in the ODB network.
	//
	// Required when creating an ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-clientsubnetcidr
	//
	ClientSubnetCidr *string `field:"optional" json:"clientSubnetCidr" yaml:"clientSubnetCidr"`
	// The domain name for the resources in the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-customdomainname
	//
	CustomDomainName *string `field:"optional" json:"customDomainName" yaml:"customDomainName"`
	// The DNS prefix to the default DNS domain name.
	//
	// The default DNS domain name is oraclevcn.com.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-defaultdnsprefix
	//
	DefaultDnsPrefix *string `field:"optional" json:"defaultDnsPrefix" yaml:"defaultDnsPrefix"`
	// Specifies whether to delete associated OCI networking resources along with the ODB network.
	//
	// Required when creating an ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-deleteassociatedresources
	//
	DeleteAssociatedResources interface{} `field:"optional" json:"deleteAssociatedResources" yaml:"deleteAssociatedResources"`
	// The user-friendly name of the ODB network.
	//
	// Required when creating an ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// The configuration for Amazon S3 access from the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-s3access
	//
	S3Access *string `field:"optional" json:"s3Access" yaml:"s3Access"`
	// Specifies the endpoint policy for Amazon S3 access from the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-s3policydocument
	//
	S3PolicyDocument *string `field:"optional" json:"s3PolicyDocument" yaml:"s3PolicyDocument"`
	// Tags to assign to the Odb Network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The configuration for Zero-ETL access from the ODB network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-odb-odbnetwork.html#cfn-odb-odbnetwork-zeroetlaccess
	//
	ZeroEtlAccess *string `field:"optional" json:"zeroEtlAccess" yaml:"zeroEtlAccess"`
}


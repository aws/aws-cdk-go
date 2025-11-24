package mixinsawsodb


// The managed services configuration for the ODB network.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   managedServicesProperty := &ManagedServicesProperty{
//   	ManagedS3BackupAccess: &ManagedS3BackupAccessProperty{
//   		Ipv4Addresses: []*string{
//   			jsii.String("ipv4Addresses"),
//   		},
//   		Status: jsii.String("status"),
//   	},
//   	ManagedServicesIpv4Cidrs: []*string{
//   		jsii.String("managedServicesIpv4Cidrs"),
//   	},
//   	ResourceGatewayArn: jsii.String("resourceGatewayArn"),
//   	S3Access: &S3AccessProperty{
//   		DomainName: jsii.String("domainName"),
//   		Ipv4Addresses: []*string{
//   			jsii.String("ipv4Addresses"),
//   		},
//   		S3PolicyDocument: jsii.String("s3PolicyDocument"),
//   		Status: jsii.String("status"),
//   	},
//   	ServiceNetworkArn: jsii.String("serviceNetworkArn"),
//   	ServiceNetworkEndpoint: &ServiceNetworkEndpointProperty{
//   		VpcEndpointId: jsii.String("vpcEndpointId"),
//   		VpcEndpointType: jsii.String("vpcEndpointType"),
//   	},
//   	ZeroEtlAccess: &ZeroEtlAccessProperty{
//   		Cidr: jsii.String("cidr"),
//   		Status: jsii.String("status"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html
//
type CfnOdbNetworkPropsMixin_ManagedServicesProperty struct {
	// The managed Amazon S3 backup access configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-manageds3backupaccess
	//
	ManagedS3BackupAccess interface{} `field:"optional" json:"managedS3BackupAccess" yaml:"managedS3BackupAccess"`
	// The IPv4 CIDR blocks for the managed services.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-managedservicesipv4cidrs
	//
	ManagedServicesIpv4Cidrs *[]*string `field:"optional" json:"managedServicesIpv4Cidrs" yaml:"managedServicesIpv4Cidrs"`
	// The Amazon Resource Name (ARN) of the resource gateway.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-resourcegatewayarn
	//
	ResourceGatewayArn *string `field:"optional" json:"resourceGatewayArn" yaml:"resourceGatewayArn"`
	// The Amazon S3 access configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-s3access
	//
	S3Access interface{} `field:"optional" json:"s3Access" yaml:"s3Access"`
	// The Amazon Resource Name (ARN) of the service network.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-servicenetworkarn
	//
	ServiceNetworkArn *string `field:"optional" json:"serviceNetworkArn" yaml:"serviceNetworkArn"`
	// The service network endpoint configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-servicenetworkendpoint
	//
	ServiceNetworkEndpoint interface{} `field:"optional" json:"serviceNetworkEndpoint" yaml:"serviceNetworkEndpoint"`
	// The Zero-ETL access configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-odb-odbnetwork-managedservices.html#cfn-odb-odbnetwork-managedservices-zeroetlaccess
	//
	ZeroEtlAccess interface{} `field:"optional" json:"zeroEtlAccess" yaml:"zeroEtlAccess"`
}


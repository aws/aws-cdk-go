package awsrds

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The CA certificate used for a DB instance.
//
// Example:
//   var vpc Vpc
//
//
//   cluster := docdb.NewDatabaseCluster(this, jsii.String("Database"), &DatabaseClusterProps{
//   	MasterUser: &Login{
//   		Username: jsii.String("myuser"),
//   	},
//   	InstanceType: ec2.InstanceType_Of(ec2.InstanceClass_MEMORY5, ec2.InstanceSize_LARGE),
//   	VpcSubnets: &SubnetSelection{
//   		SubnetType: ec2.SubnetType_PUBLIC,
//   	},
//   	Vpc: Vpc,
//   	CaCertificate: docdb.CaCertificate_RDS_CA_RSA4096_G1(),
//   })
//
// See: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html
//
type CaCertificate interface {
	// Returns the CA certificate identifier as a string.
	ToString() *string
}

// The jsii proxy struct for CaCertificate
type jsiiProxy_CaCertificate struct {
	_ byte // padding
}

// Custom CA certificate.
func CaCertificate_Of(identifier *string) CaCertificate {
	_init_.Initialize()

	if err := validateCaCertificate_OfParameters(identifier); err != nil {
		panic(err)
	}
	var returns CaCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"of",
		[]interface{}{identifier},
		&returns,
	)

	return returns
}

func CaCertificate_RDS_CA_2019() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_2019",
		&returns,
	)
	return returns
}

func CaCertificate_RDS_CA_ECC384_G1() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_ECC384_G1",
		&returns,
	)
	return returns
}

func CaCertificate_RDS_CA_RDS2048_G1() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_RDS2048_G1",
		&returns,
	)
	return returns
}

func CaCertificate_RDS_CA_RDS4096_G1() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_RDS4096_G1",
		&returns,
	)
	return returns
}

func CaCertificate_RDS_CA_RSA2048_G1() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_RSA2048_G1",
		&returns,
	)
	return returns
}

func CaCertificate_RDS_CA_RSA4096_G1() CaCertificate {
	_init_.Initialize()
	var returns CaCertificate
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_rds.CaCertificate",
		"RDS_CA_RSA4096_G1",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CaCertificate) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


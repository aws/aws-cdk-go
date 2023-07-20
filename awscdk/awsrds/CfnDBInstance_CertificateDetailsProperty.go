package awsrds


// Returns the details of the DB instance’s server certificate.
//
// For more information, see [Using SSL/TLS to encrypt a connection to a DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html) in the *Amazon RDS User Guide* and [Using SSL/TLS to encrypt a connection to a DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html) in the *Amazon Aurora User Guide* .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   certificateDetailsProperty := &CertificateDetailsProperty{
//   	CaIdentifier: jsii.String("caIdentifier"),
//   	ValidTill: jsii.String("validTill"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-certificatedetails.html
//
type CfnDBInstance_CertificateDetailsProperty struct {
	// The CA identifier of the CA certificate used for the DB instance's server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-certificatedetails.html#cfn-rds-dbinstance-certificatedetails-caidentifier
	//
	CaIdentifier *string `field:"optional" json:"caIdentifier" yaml:"caIdentifier"`
	// The expiration date of the DB instance’s server certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbinstance-certificatedetails.html#cfn-rds-dbinstance-certificatedetails-validtill
	//
	ValidTill *string `field:"optional" json:"validTill" yaml:"validTill"`
}


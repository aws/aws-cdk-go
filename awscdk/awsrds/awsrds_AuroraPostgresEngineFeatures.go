package awsrds


// Features supported by this version of the Aurora Postgres cluster engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraPostgresEngineFeatures := &auroraPostgresEngineFeatures{
//   	s3Export: jsii.Boolean(false),
//   	s3Import: jsii.Boolean(false),
//   }
//
// Experimental.
type AuroraPostgresEngineFeatures struct {
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data export feature.
	// Experimental.
	S3Export *bool `field:"optional" json:"s3Export" yaml:"s3Export"`
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data import feature.
	// Experimental.
	S3Import *bool `field:"optional" json:"s3Import" yaml:"s3Import"`
}


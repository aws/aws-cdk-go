package awsrds


// Features supported by the Postgres database engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   postgresEngineFeatures := &PostgresEngineFeatures{
//   	S3Export: jsii.Boolean(false),
//   	S3Import: jsii.Boolean(false),
//   }
//
type PostgresEngineFeatures struct {
	// Whether this version of the Postgres engine supports the S3 data export feature.
	// Default: false.
	//
	S3Export *bool `field:"optional" json:"s3Export" yaml:"s3Export"`
	// Whether this version of the Postgres engine supports the S3 data import feature.
	// Default: false.
	//
	S3Import *bool `field:"optional" json:"s3Import" yaml:"s3Import"`
}


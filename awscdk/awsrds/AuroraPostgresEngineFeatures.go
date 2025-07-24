package awsrds


// Features supported by this version of the Aurora Postgres cluster engine.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   auroraPostgresEngineFeatures := &AuroraPostgresEngineFeatures{
//   	S3Export: jsii.Boolean(false),
//   	S3Import: jsii.Boolean(false),
//   	ServerlessV2AutoPauseSupported: jsii.Boolean(false),
//   }
//
type AuroraPostgresEngineFeatures struct {
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data export feature.
	// Default: false.
	//
	S3Export *bool `field:"optional" json:"s3Export" yaml:"s3Export"`
	// Whether this version of the Aurora Postgres cluster engine supports the S3 data import feature.
	// Default: false.
	//
	S3Import *bool `field:"optional" json:"s3Import" yaml:"s3Import"`
	// Whether this version of the Aurora Postgres cluster engine supports the Aurora SeverlessV2 auto-pause feature.
	// See: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless-v2-auto-pause.html#auto-pause-prereqs
	//
	// Default: false.
	//
	ServerlessV2AutoPauseSupported *bool `field:"optional" json:"serverlessV2AutoPauseSupported" yaml:"serverlessV2AutoPauseSupported"`
}


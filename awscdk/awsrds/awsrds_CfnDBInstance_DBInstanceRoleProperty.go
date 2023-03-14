package awsrds


// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBInstanceRoleProperty := &dBInstanceRoleProperty{
//   	featureName: jsii.String("featureName"),
//   	roleArn: jsii.String("roleArn"),
//   }
//
type CfnDBInstance_DBInstanceRoleProperty struct {
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB instance grant permission for the DB instance to access other AWS services on your behalf. For the list of supported feature names, see the `SupportedFeatureNames` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference* .
	FeatureName *string `field:"required" json:"featureName" yaml:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB instance.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
}


package awsrds


// Describes an AWS Identity and Access Management (IAM) role that is associated with a DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   dBClusterRoleProperty := &dBClusterRoleProperty{
//   	roleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	featureName: jsii.String("featureName"),
//   }
//
type CfnDBCluster_DBClusterRoleProperty struct {
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The name of the feature associated with the AWS Identity and Access Management (IAM) role.
	//
	// IAM roles that are associated with a DB cluster grant permission for the DB cluster to access other AWS services on your behalf. For the list of supported feature names, see the `SupportedFeatureNames` description in [DBEngineVersion](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DBEngineVersion.html) in the *Amazon RDS API Reference* .
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
}


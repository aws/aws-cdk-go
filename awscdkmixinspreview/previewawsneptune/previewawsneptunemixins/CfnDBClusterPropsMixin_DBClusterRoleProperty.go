package previewawsneptunemixins


// Describes an Amazon Identity and Access Management (IAM) role that is associated with a DB cluster.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dBClusterRoleProperty := &DBClusterRoleProperty{
//   	FeatureName: jsii.String("featureName"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-dbclusterrole.html
//
type CfnDBClusterPropsMixin_DBClusterRoleProperty struct {
	// The name of the feature associated with the Amazon Identity and Access Management (IAM) role.
	//
	// For the list of supported feature names, see [DescribeDBEngineVersions](https://docs.aws.amazon.com/neptune/latest/userguide/api-other-apis.html#DescribeDBEngineVersions) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-dbclusterrole.html#cfn-neptune-dbcluster-dbclusterrole-featurename
	//
	FeatureName *string `field:"optional" json:"featureName" yaml:"featureName"`
	// The Amazon Resource Name (ARN) of the IAM role that is associated with the DB cluster.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-neptune-dbcluster-dbclusterrole.html#cfn-neptune-dbcluster-dbclusterrole-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


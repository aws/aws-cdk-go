package awsquicksight


// A structure that grants Amazon QuickSight access to your cluster and make a call to the `redshift:GetClusterCredentials` API.
//
// For more information on the `redshift:GetClusterCredentials` API, see [`GetClusterCredentials`](https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   redshiftIAMParametersProperty := &RedshiftIAMParametersProperty{
//   	RoleArn: jsii.String("roleArn"),
//
//   	// the properties below are optional
//   	AutoCreateDatabaseUser: jsii.Boolean(false),
//   	DatabaseGroups: []*string{
//   		jsii.String("databaseGroups"),
//   	},
//   	DatabaseUser: jsii.String("databaseUser"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html
//
type CfnDataSource_RedshiftIAMParametersProperty struct {
	// Use the `RoleArn` structure to allow Amazon QuickSight to call `redshift:GetClusterCredentials` on your cluster.
	//
	// The calling principal must have `iam:PassRole` access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Automatically creates a database user.
	//
	// If your database doesn't have a `DatabaseUser` , set this parameter to `True` . If there is no `DatabaseUser` , Amazon QuickSight can't connect to your cluster. The `RoleArn` that you use for this operation must grant access to `redshift:CreateClusterUser` to successfully create the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-autocreatedatabaseuser
	//
	// Default: - false.
	//
	AutoCreateDatabaseUser interface{} `field:"optional" json:"autoCreateDatabaseUser" yaml:"autoCreateDatabaseUser"`
	// A list of groups whose permissions will be granted to Amazon QuickSight to access the cluster.
	//
	// These permissions are combined with the permissions granted to Amazon QuickSight by the `DatabaseUser` . If you choose to include this parameter, the `RoleArn` must grant access to `redshift:JoinGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databasegroups
	//
	DatabaseGroups *[]*string `field:"optional" json:"databaseGroups" yaml:"databaseGroups"`
	// The user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster.
	//
	// If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of `AutoCreateDatabaseUser` to `True` to create a new user with PUBLIC permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
}


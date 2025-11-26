package previewawsquicksightmixins


// A structure that grants Quick Sight access to your cluster and make a call to the `redshift:GetClusterCredentials` API.
//
// For more information on the `redshift:GetClusterCredentials` API, see [`GetClusterCredentials`](https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   redshiftIAMParametersProperty := &RedshiftIAMParametersProperty{
//   	AutoCreateDatabaseUser: jsii.Boolean(false),
//   	DatabaseGroups: []*string{
//   		jsii.String("databaseGroups"),
//   	},
//   	DatabaseUser: jsii.String("databaseUser"),
//   	RoleArn: jsii.String("roleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html
//
type CfnDataSourcePropsMixin_RedshiftIAMParametersProperty struct {
	// Automatically creates a database user.
	//
	// If your database doesn't have a `DatabaseUser` , set this parameter to `True` . If there is no `DatabaseUser` , Quick Sight can't connect to your cluster. The `RoleArn` that you use for this operation must grant access to `redshift:CreateClusterUser` to successfully create the user.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-autocreatedatabaseuser
	//
	// Default: - false.
	//
	AutoCreateDatabaseUser interface{} `field:"optional" json:"autoCreateDatabaseUser" yaml:"autoCreateDatabaseUser"`
	// A list of groups whose permissions will be granted to Quick Sight to access the cluster.
	//
	// These permissions are combined with the permissions granted to Quick Sight by the `DatabaseUser` . If you choose to include this parameter, the `RoleArn` must grant access to `redshift:JoinGroup` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databasegroups
	//
	DatabaseGroups *[]*string `field:"optional" json:"databaseGroups" yaml:"databaseGroups"`
	// The user whose permissions and group memberships will be used by Quick Sight to access the cluster.
	//
	// If this user already exists in your database, Amazon Quick Sight is granted the same permissions that the user has. If the user doesn't exist, set the value of `AutoCreateDatabaseUser` to `True` to create a new user with PUBLIC permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
	// Use the `RoleArn` structure to allow Quick Sight to call `redshift:GetClusterCredentials` on your cluster.
	//
	// The calling principal must have `iam:PassRole` access to pass the role to Quick Sight. The role's trust policy must allow the Quick Sight service principal to assume the role.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}


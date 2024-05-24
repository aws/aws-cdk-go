package awsquicksight


// <p>A structure that grants Amazon QuickSight access to your cluster and make a call to the <code>redshift:GetClusterCredentials</code> API.
//
// For more information on the <code>redshift:GetClusterCredentials</code> API, see <a href="https://docs.aws.amazon.com/redshift/latest/APIReference/API_GetClusterCredentials.html">
//                <code>GetClusterCredentials</code>
//             </a>.</p>
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
	// <p>Use the <code>RoleArn</code> structure to allow Amazon QuickSight to call <code>redshift:GetClusterCredentials</code> on your cluster.
	//
	// The calling principal must have <code>iam:PassRole</code> access to pass the role to Amazon QuickSight. The role's trust policy must allow the Amazon QuickSight service principal to assume the role.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-rolearn
	//
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// <p>Automatically creates a database user.
	//
	// If your database doesn't have a <code>DatabaseUser</code>, set this parameter to <code>True</code>. If there is no <code>DatabaseUser</code>, Amazon QuickSight can't connect to your cluster. The <code>RoleArn</code> that you use for this operation must grant access to <code>redshift:CreateClusterUser</code> to successfully create the user.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-autocreatedatabaseuser
	//
	// Default: - false.
	//
	AutoCreateDatabaseUser interface{} `field:"optional" json:"autoCreateDatabaseUser" yaml:"autoCreateDatabaseUser"`
	// <p>A list of groups whose permissions will be granted to Amazon QuickSight to access the cluster.
	//
	// These permissions are combined with the permissions granted to Amazon QuickSight by the <code>DatabaseUser</code>. If you choose to include this parameter, the <code>RoleArn</code> must grant access to <code>redshift:JoinGroup</code>.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databasegroups
	//
	DatabaseGroups *[]*string `field:"optional" json:"databaseGroups" yaml:"databaseGroups"`
	// <p>The user whose permissions and group memberships will be used by Amazon QuickSight to access the cluster.
	//
	// If this user already exists in your database, Amazon QuickSight is granted the same permissions that the user has. If the user doesn't exist, set the value of <code>AutoCreateDatabaseUser</code> to <code>True</code> to create a new user with PUBLIC permissions.</p>
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-datasource-redshiftiamparameters.html#cfn-quicksight-datasource-redshiftiamparameters-databaseuser
	//
	DatabaseUser *string `field:"optional" json:"databaseUser" yaml:"databaseUser"`
}


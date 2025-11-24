package mixinsawskendra


// Provides information about the column that should be used for filtering the query response by groups.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   aclConfigurationProperty := &AclConfigurationProperty{
//   	AllowedGroupsColumnName: jsii.String("allowedGroupsColumnName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-aclconfiguration.html
//
type CfnDataSourcePropsMixin_AclConfigurationProperty struct {
	// A list of groups, separated by semi-colons, that filters a query response based on user context.
	//
	// The document is only returned to users that are in one of the groups specified in the `UserContext` field of the [Query](https://docs.aws.amazon.com/kendra/latest/dg/API_Query.html) operation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-aclconfiguration.html#cfn-kendra-datasource-aclconfiguration-allowedgroupscolumnname
	//
	AllowedGroupsColumnName *string `field:"optional" json:"allowedGroupsColumnName" yaml:"allowedGroupsColumnName"`
}


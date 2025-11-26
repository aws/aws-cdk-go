package previewawslakeformationmixins


// A list of LF-tag conditions that define a resource's LF-tag policy.
//
// A structure that allows an admin to grant user permissions on certain conditions. For example, granting a role access to all columns that do not have the LF-tag 'PII' in tables that have the LF-tag 'Prod'.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   lFTagPolicyResourceProperty := &LFTagPolicyResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	Expression: []interface{}{
//   		&LFTagProperty{
//   			TagKey: jsii.String("tagKey"),
//   			TagValues: []*string{
//   				jsii.String("tagValues"),
//   			},
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagpolicyresource.html
//
type CfnPrincipalPermissionsPropsMixin_LFTagPolicyResourceProperty struct {
	// The identifier for the Data Catalog .
	//
	// The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagpolicyresource.html#cfn-lakeformation-principalpermissions-lftagpolicyresource-catalogid
	//
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
	// A list of LF-tag conditions that apply to the resource's LF-tag policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagpolicyresource.html#cfn-lakeformation-principalpermissions-lftagpolicyresource-expression
	//
	Expression interface{} `field:"optional" json:"expression" yaml:"expression"`
	// The resource type for which the LF-tag policy applies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagpolicyresource.html#cfn-lakeformation-principalpermissions-lftagpolicyresource-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}


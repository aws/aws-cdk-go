package awslakeformation


// A structure containing an LF-tag key and values for a resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagKeyResourceProperty := &LFTagKeyResourceProperty{
//   	CatalogId: jsii.String("catalogId"),
//   	TagKey: jsii.String("tagKey"),
//   	TagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagkeyresource.html
//
type CfnPrincipalPermissions_LFTagKeyResourceProperty struct {
	// The identifier for the Data Catalog where the location is registered with Data Catalog .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagkeyresource.html#cfn-lakeformation-principalpermissions-lftagkeyresource-catalogid
	//
	CatalogId *string `field:"required" json:"catalogId" yaml:"catalogId"`
	// The key-name for the LF-tag.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagkeyresource.html#cfn-lakeformation-principalpermissions-lftagkeyresource-tagkey
	//
	TagKey *string `field:"required" json:"tagKey" yaml:"tagKey"`
	// A list of possible values for the corresponding `TagKey` of an LF-tag key-value pair.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-principalpermissions-lftagkeyresource.html#cfn-lakeformation-principalpermissions-lftagkeyresource-tagvalues
	//
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}


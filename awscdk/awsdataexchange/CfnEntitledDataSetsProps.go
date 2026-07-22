package awsdataexchange


// Properties for defining a `CfnEntitledDataSets`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEntitledDataSetsProps := &CfnEntitledDataSetsProps{
//   	AssetType: jsii.String("assetType"),
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-entitleddatasets.html
//
type CfnEntitledDataSetsProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-entitleddatasets.html#cfn-dataexchange-entitleddatasets-assettype
	//
	AssetType *string `field:"optional" json:"assetType" yaml:"assetType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-entitleddatasets.html#cfn-dataexchange-entitleddatasets-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dataexchange-entitleddatasets.html#cfn-dataexchange-entitleddatasets-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
}


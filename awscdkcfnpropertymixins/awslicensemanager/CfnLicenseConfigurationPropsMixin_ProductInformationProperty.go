package awslicensemanager


// Describes product information for a license configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   productInformationProperty := &ProductInformationProperty{
//   	ProductInformationFilterList: []interface{}{
//   		&ProductInformationFilterProperty{
//   			ProductInformationFilterComparator: jsii.String("productInformationFilterComparator"),
//   			ProductInformationFilterName: jsii.String("productInformationFilterName"),
//   			ProductInformationFilterValue: []*string{
//   				jsii.String("productInformationFilterValue"),
//   			},
//   		},
//   	},
//   	ResourceType: jsii.String("resourceType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformation.html
//
type CfnLicenseConfigurationPropsMixin_ProductInformationProperty struct {
	// A product information filter list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformation.html#cfn-licensemanager-licenseconfiguration-productinformation-productinformationfilterlist
	//
	ProductInformationFilterList interface{} `field:"optional" json:"productInformationFilterList" yaml:"productInformationFilterList"`
	// Resource type.
	//
	// The possible values are SSM_MANAGED | RDS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformation.html#cfn-licensemanager-licenseconfiguration-productinformation-resourcetype
	//
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}


package awslicensemanager


// Describes a product information filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   productInformationFilterProperty := &ProductInformationFilterProperty{
//   	ProductInformationFilterComparator: jsii.String("productInformationFilterComparator"),
//   	ProductInformationFilterName: jsii.String("productInformationFilterName"),
//   	ProductInformationFilterValue: []*string{
//   		jsii.String("productInformationFilterValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html
//
type CfnLicenseConfigurationPropsMixin_ProductInformationFilterProperty struct {
	// Logical operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltercomparator
	//
	ProductInformationFilterComparator *string `field:"optional" json:"productInformationFilterComparator" yaml:"productInformationFilterComparator"`
	// Filter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltername
	//
	ProductInformationFilterName *string `field:"optional" json:"productInformationFilterName" yaml:"productInformationFilterName"`
	// Filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltervalue
	//
	ProductInformationFilterValue *[]*string `field:"optional" json:"productInformationFilterValue" yaml:"productInformationFilterValue"`
}


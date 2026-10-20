package awslicensemanager


// Describes a product information filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   productInformationFilterProperty := &ProductInformationFilterProperty{
//   	ProductInformationFilterComparator: jsii.String("productInformationFilterComparator"),
//   	ProductInformationFilterName: jsii.String("productInformationFilterName"),
//
//   	// the properties below are optional
//   	ProductInformationFilterValue: []*string{
//   		jsii.String("productInformationFilterValue"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html
//
type CfnLicenseConfiguration_ProductInformationFilterProperty struct {
	// Logical operator.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltercomparator
	//
	ProductInformationFilterComparator *string `field:"required" json:"productInformationFilterComparator" yaml:"productInformationFilterComparator"`
	// Filter name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltername
	//
	ProductInformationFilterName *string `field:"required" json:"productInformationFilterName" yaml:"productInformationFilterName"`
	// Filter value.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-licensemanager-licenseconfiguration-productinformationfilter.html#cfn-licensemanager-licenseconfiguration-productinformationfilter-productinformationfiltervalue
	//
	ProductInformationFilterValue *[]*string `field:"optional" json:"productInformationFilterValue" yaml:"productInformationFilterValue"`
}


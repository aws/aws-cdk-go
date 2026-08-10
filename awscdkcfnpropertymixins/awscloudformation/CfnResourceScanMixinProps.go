package awscloudformation


// Properties for CfnResourceScanPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnResourceScanMixinProps := &CfnResourceScanMixinProps{
//   	ScanFilters: []interface{}{
//   		&ScanFilterProperty{
//   			Types: []*string{
//   				jsii.String("types"),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourcescan.html
//
type CfnResourceScanMixinProps struct {
	// The scan filters to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cloudformation-resourcescan.html#cfn-cloudformation-resourcescan-scanfilters
	//
	ScanFilters interface{} `field:"optional" json:"scanFilters" yaml:"scanFilters"`
}


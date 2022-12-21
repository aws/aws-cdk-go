package awswafv2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aWSManagedRulesBotControlRuleSetProperty := &aWSManagedRulesBotControlRuleSetProperty{
//   	inspectionLevel: jsii.String("inspectionLevel"),
//   }
//
type CfnWebACL_AWSManagedRulesBotControlRuleSetProperty struct {
	// `CfnWebACL.AWSManagedRulesBotControlRuleSetProperty.InspectionLevel`.
	InspectionLevel *string `field:"required" json:"inspectionLevel" yaml:"inspectionLevel"`
}


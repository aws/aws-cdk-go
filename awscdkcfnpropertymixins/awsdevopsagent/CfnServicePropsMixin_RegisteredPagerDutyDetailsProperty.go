package awsdevopsagent


// PagerDuty service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   registeredPagerDutyDetailsProperty := &RegisteredPagerDutyDetailsProperty{
//   	Scopes: []*string{
//   		jsii.String("scopes"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredpagerdutydetails.html
//
type CfnServicePropsMixin_RegisteredPagerDutyDetailsProperty struct {
	// The scopes assigned to the service.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registeredpagerdutydetails.html#cfn-devopsagent-service-registeredpagerdutydetails-scopes
	//
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
}


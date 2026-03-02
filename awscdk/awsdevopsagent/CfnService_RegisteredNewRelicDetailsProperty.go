package awsdevopsagent


// New Relic service details returned after registration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   registeredNewRelicDetailsProperty := &RegisteredNewRelicDetailsProperty{
//   	AccountId: jsii.String("accountId"),
//   	Region: jsii.String("region"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html
//
type CfnService_RegisteredNewRelicDetailsProperty struct {
	// New Relic account ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-accountid
	//
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// New Relic region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-region
	//
	Region *string `field:"required" json:"region" yaml:"region"`
	// Optional user description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-devopsagent-service-registerednewrelicdetails.html#cfn-devopsagent-service-registerednewrelicdetails-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}


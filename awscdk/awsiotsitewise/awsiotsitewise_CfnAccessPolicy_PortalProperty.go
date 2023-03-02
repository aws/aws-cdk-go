package awsiotsitewise


// The `Portal` property type specifies the AWS IoT SiteWise Monitor portal for an [AWS::IoTSiteWise::AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   portalProperty := &portalProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnAccessPolicy_PortalProperty struct {
	// The ID of the portal.
	Id *string `field:"optional" json:"id" yaml:"id"`
}


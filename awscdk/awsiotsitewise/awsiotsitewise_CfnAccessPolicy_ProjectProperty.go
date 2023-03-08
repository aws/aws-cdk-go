package awsiotsitewise


// The `Project` property type specifies the AWS IoT SiteWise Monitor project for an [AWS::IoTSiteWise::AccessPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iotsitewise-accesspolicy.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   projectProperty := &projectProperty{
//   	id: jsii.String("id"),
//   }
//
type CfnAccessPolicy_ProjectProperty struct {
	// The ID of the project.
	Id *string `field:"optional" json:"id" yaml:"id"`
}


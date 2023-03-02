package awsiotsitewise


// The AWS IoT SiteWise Monitor resource for this access policy.
//
// Choose either a portal or a project.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   accessPolicyResourceProperty := &accessPolicyResourceProperty{
//   	portal: &portalProperty{
//   		id: jsii.String("id"),
//   	},
//   	project: &projectProperty{
//   		id: jsii.String("id"),
//   	},
//   }
//
type CfnAccessPolicy_AccessPolicyResourceProperty struct {
	// The AWS IoT SiteWise Monitor portal for this access policy.
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// The AWS IoT SiteWise Monitor project for this access policy.
	Project interface{} `field:"optional" json:"project" yaml:"project"`
}


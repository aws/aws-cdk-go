package awssupportauthz


// The set of resources a support permit applies to.
//
// Exactly one of AllResourcesInRegion or Resources must be provided.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var allResourcesInRegion interface{}
//
//   resourceSetProperty := &ResourceSetProperty{
//   	AllResourcesInRegion: allResourcesInRegion,
//   	Resources: []*string{
//   		jsii.String("resources"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-resourceset.html
//
type CfnSupportPermitPropsMixin_ResourceSetProperty struct {
	// Applies to all resources in the region.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-resourceset.html#cfn-supportauthz-supportpermit-resourceset-allresourcesinregion
	//
	AllResourcesInRegion interface{} `field:"optional" json:"allResourcesInRegion" yaml:"allResourcesInRegion"`
	// An explicit list of resource ARNs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-supportauthz-supportpermit-resourceset.html#cfn-supportauthz-supportpermit-resourceset-resources
	//
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}


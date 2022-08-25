package awslakeformation


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   lFTagProperty := &lFTagProperty{
//   	tagKey: jsii.String("tagKey"),
//   	tagValues: []*string{
//   		jsii.String("tagValues"),
//   	},
//   }
//
type CfnPrincipalPermissions_LFTagProperty struct {
	// `CfnPrincipalPermissions.LFTagProperty.TagKey`.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// `CfnPrincipalPermissions.LFTagProperty.TagValues`.
	TagValues *[]*string `field:"optional" json:"tagValues" yaml:"tagValues"`
}


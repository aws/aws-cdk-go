package awsec2


// Properties for looking up an existing managed prefix list.
//
// Example:
//   var alb ApplicationLoadBalancer
//
//
//   cfOriginFacing := ec2.PrefixList_FromLookup(this, jsii.String("CloudFrontOriginFacing"), &PrefixListLookupOptions{
//   	PrefixListName: jsii.String("com.amazonaws.global.cloudfront.origin-facing"),
//   })
//   alb.Connections.AllowFrom(cfOriginFacing, ec2.Port_HTTP())
//
type PrefixListLookupOptions struct {
	// The name of the managed prefix list.
	PrefixListName *string `field:"required" json:"prefixListName" yaml:"prefixListName"`
	// The address family of the managed prefix list.
	// Default: - Don't filter on addressFamily.
	//
	AddressFamily AddressFamily `field:"optional" json:"addressFamily" yaml:"addressFamily"`
	// The ID of the AWS account that owns the managed prefix list.
	// Default: - Don't filter on ownerId.
	//
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
}


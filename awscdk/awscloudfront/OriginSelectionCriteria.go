package awscloudfront


// The selection criteria for the origin group.
//
// Example:
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewOriginGroup(&OriginGroupProps{
//   			PrimaryOrigin: origins.NewHttpOrigin(jsii.String("<AWS Elemental MediaPackageV2 origin 1>")),
//   			FallbackOrigin: origins.NewHttpOrigin(jsii.String("<AWS Elemental MediaPackageV2 origin 2>")),
//   			FallbackStatusCodes: []*f64{
//   				jsii.Number(404),
//   			},
//   			SelectionCriteria: cloudfront.OriginSelectionCriteria_MEDIA_QUALITY_BASED,
//   		}),
//   	},
//   })
//
type OriginSelectionCriteria string

const (
	// Default selection behavior.
	OriginSelectionCriteria_DEFAULT OriginSelectionCriteria = "DEFAULT"
	// Selection based on media quality.
	//
	// This option is only valid for AWS Elemental MediaPackage v2 Origins.
	OriginSelectionCriteria_MEDIA_QUALITY_BASED OriginSelectionCriteria = "MEDIA_QUALITY_BASED"
)


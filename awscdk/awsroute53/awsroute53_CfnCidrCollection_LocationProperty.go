package awsroute53


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   locationProperty := &locationProperty{
//   	cidrList: []*string{
//   		jsii.String("cidrList"),
//   	},
//   	locationName: jsii.String("locationName"),
//   }
//
type CfnCidrCollection_LocationProperty struct {
	// `CfnCidrCollection.LocationProperty.CidrList`.
	CidrList *[]*string `field:"required" json:"cidrList" yaml:"cidrList"`
	// `CfnCidrCollection.LocationProperty.LocationName`.
	LocationName *string `field:"required" json:"locationName" yaml:"locationName"`
}


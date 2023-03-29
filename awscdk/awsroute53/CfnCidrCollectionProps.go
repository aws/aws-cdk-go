package awsroute53


// Properties for defining a `CfnCidrCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCidrCollectionProps := &CfnCidrCollectionProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Locations: []interface{}{
//   		&LocationProperty{
//   			CidrList: []*string{
//   				jsii.String("cidrList"),
//   			},
//   			LocationName: jsii.String("locationName"),
//   		},
//   	},
//   }
//
type CfnCidrCollectionProps struct {
	// The name of a CIDR collection.
	Name *string `field:"required" json:"name" yaml:"name"`
	// A complex type that contains information about the list of CIDR locations.
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}


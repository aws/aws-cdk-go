package awsroute53


// Properties for defining a `CfnCidrCollection`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCidrCollectionProps := &cfnCidrCollectionProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	locations: []interface{}{
//   		&locationProperty{
//   			cidrList: []*string{
//   				jsii.String("cidrList"),
//   			},
//   			locationName: jsii.String("locationName"),
//   		},
//   	},
//   }
//
type CfnCidrCollectionProps struct {
	// `AWS::Route53::CidrCollection.Name`.
	Name *string `field:"required" json:"name" yaml:"name"`
	// `AWS::Route53::CidrCollection.Locations`.
	Locations interface{} `field:"optional" json:"locations" yaml:"locations"`
}


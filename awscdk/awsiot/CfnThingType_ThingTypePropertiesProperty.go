package awsiot


// The ThingTypeProperties contains information about the thing type including: a thing type description, and a list of searchable thing attribute names.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   thingTypePropertiesProperty := &ThingTypePropertiesProperty{
//   	SearchableAttributes: []*string{
//   		jsii.String("searchableAttributes"),
//   	},
//   	ThingTypeDescription: jsii.String("thingTypeDescription"),
//   }
//
type CfnThingType_ThingTypePropertiesProperty struct {
	// A list of searchable thing attribute names.
	SearchableAttributes *[]*string `field:"optional" json:"searchableAttributes" yaml:"searchableAttributes"`
	// The description of the thing type.
	ThingTypeDescription *string `field:"optional" json:"thingTypeDescription" yaml:"thingTypeDescription"`
}


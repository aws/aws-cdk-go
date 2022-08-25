package awscustomerprofiles


// An object that defines the Key element of a ProfileObject.
//
// A Key is a special element that can be used to search for a customer profile.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   objectTypeKeyProperty := &objectTypeKeyProperty{
//   	fieldNames: []*string{
//   		jsii.String("fieldNames"),
//   	},
//   	standardIdentifiers: []*string{
//   		jsii.String("standardIdentifiers"),
//   	},
//   }
//
type CfnObjectType_ObjectTypeKeyProperty struct {
	// The reference for the key name of the fields map.
	FieldNames *[]*string `field:"optional" json:"fieldNames" yaml:"fieldNames"`
	// The types of keys that a ProfileObject can have.
	//
	// Each ProfileObject can have only 1 UNIQUE key but multiple PROFILE keys. PROFILE means that this key can be used to tie an object to a PROFILE. UNIQUE means that it can be used to uniquely identify an object. If a key a is marked as SECONDARY, it will be used to search for profiles after all other PROFILE keys have been searched. A LOOKUP_ONLY key is only used to match a profile but is not persisted to be used for searching of the profile. A NEW_ONLY key is only used if the profile does not already exist before the object is ingested, otherwise it is only used for matching objects to profiles.
	StandardIdentifiers *[]*string `field:"optional" json:"standardIdentifiers" yaml:"standardIdentifiers"`
}


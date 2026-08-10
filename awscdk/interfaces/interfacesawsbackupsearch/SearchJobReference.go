package interfacesawsbackupsearch


// A reference to a SearchJob resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   searchJobReference := &SearchJobReference{
//   	SearchJobArn: jsii.String("searchJobArn"),
//   }
//
type SearchJobReference struct {
	// The SearchJobArn of the SearchJob resource.
	SearchJobArn *string `field:"required" json:"searchJobArn" yaml:"searchJobArn"`
}


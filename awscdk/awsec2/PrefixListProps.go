package awsec2


// Properties for creating a prefix list.
//
// Example:
//   ec2.NewPrefixList(this, jsii.String("EmptyPrefixList"), &PrefixListProps{
//   	MaxEntries: jsii.Number(100),
//   })
//
type PrefixListProps struct {
	// The maximum number of entries for the prefix list.
	// Default: Automatically-calculated.
	//
	MaxEntries *float64 `field:"optional" json:"maxEntries" yaml:"maxEntries"`
	// The address family of the prefix list.
	// Default: AddressFamily.IP_V4
	//
	AddressFamily AddressFamily `field:"optional" json:"addressFamily" yaml:"addressFamily"`
	// The list of entries for the prefix list.
	// Default: [].
	//
	Entries *[]*CfnPrefixList_EntryProperty `field:"optional" json:"entries" yaml:"entries"`
	// The name of the prefix list.
	// Default: None.
	//
	PrefixListName *string `field:"optional" json:"prefixListName" yaml:"prefixListName"`
}


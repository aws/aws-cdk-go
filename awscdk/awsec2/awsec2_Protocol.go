package awsec2


// Protocol for use in Connection Rules.
//
// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
// Experimental.
type Protocol string

const (
	// Experimental.
	Protocol_ALL Protocol = "ALL"
	// Experimental.
	Protocol_HOPOPT Protocol = "HOPOPT"
	// Experimental.
	Protocol_ICMP Protocol = "ICMP"
	// Experimental.
	Protocol_IGMP Protocol = "IGMP"
	// Experimental.
	Protocol_GGP Protocol = "GGP"
	// Experimental.
	Protocol_IPV4 Protocol = "IPV4"
	// Experimental.
	Protocol_ST Protocol = "ST"
	// Experimental.
	Protocol_TCP Protocol = "TCP"
	// Experimental.
	Protocol_CBT Protocol = "CBT"
	// Experimental.
	Protocol_EGP Protocol = "EGP"
	// Experimental.
	Protocol_IGP Protocol = "IGP"
	// Experimental.
	Protocol_BBN_RCC_MON Protocol = "BBN_RCC_MON"
	// Experimental.
	Protocol_NVP_II Protocol = "NVP_II"
	// Experimental.
	Protocol_PUP Protocol = "PUP"
	// Experimental.
	Protocol_EMCON Protocol = "EMCON"
	// Experimental.
	Protocol_XNET Protocol = "XNET"
	// Experimental.
	Protocol_CHAOS Protocol = "CHAOS"
	// Experimental.
	Protocol_UDP Protocol = "UDP"
	// Experimental.
	Protocol_MUX Protocol = "MUX"
	// Experimental.
	Protocol_DCN_MEAS Protocol = "DCN_MEAS"
	// Experimental.
	Protocol_HMP Protocol = "HMP"
	// Experimental.
	Protocol_PRM Protocol = "PRM"
	// Experimental.
	Protocol_XNS_IDP Protocol = "XNS_IDP"
	// Experimental.
	Protocol_TRUNK_1 Protocol = "TRUNK_1"
	// Experimental.
	Protocol_TRUNK_2 Protocol = "TRUNK_2"
	// Experimental.
	Protocol_LEAF_1 Protocol = "LEAF_1"
	// Experimental.
	Protocol_LEAF_2 Protocol = "LEAF_2"
	// Experimental.
	Protocol_RDP Protocol = "RDP"
	// Experimental.
	Protocol_IRTP Protocol = "IRTP"
	// Experimental.
	Protocol_ISO_TP4 Protocol = "ISO_TP4"
	// Experimental.
	Protocol_NETBLT Protocol = "NETBLT"
	// Experimental.
	Protocol_MFE_NSP Protocol = "MFE_NSP"
	// Experimental.
	Protocol_MERIT_INP Protocol = "MERIT_INP"
	// Experimental.
	Protocol_DCCP Protocol = "DCCP"
	// Experimental.
	Protocol_THREEPC Protocol = "THREEPC"
	// Experimental.
	Protocol_IDPR Protocol = "IDPR"
	// Experimental.
	Protocol_XTP Protocol = "XTP"
	// Experimental.
	Protocol_DDP Protocol = "DDP"
	// Experimental.
	Protocol_IDPR_CMTP Protocol = "IDPR_CMTP"
	// Experimental.
	Protocol_TPPLUSPLUS Protocol = "TPPLUSPLUS"
	// Experimental.
	Protocol_IL Protocol = "IL"
	// Experimental.
	Protocol_IPV6 Protocol = "IPV6"
	// Experimental.
	Protocol_SDRP Protocol = "SDRP"
	// Experimental.
	Protocol_IPV6_ROUTE Protocol = "IPV6_ROUTE"
	// Experimental.
	Protocol_IPV6_FRAG Protocol = "IPV6_FRAG"
	// Experimental.
	Protocol_IDRP Protocol = "IDRP"
	// Experimental.
	Protocol_RSVP Protocol = "RSVP"
	// Experimental.
	Protocol_GRE Protocol = "GRE"
	// Experimental.
	Protocol_DSR Protocol = "DSR"
	// Experimental.
	Protocol_BNA Protocol = "BNA"
	// Experimental.
	Protocol_ESP Protocol = "ESP"
	// Experimental.
	Protocol_AH Protocol = "AH"
	// Experimental.
	Protocol_I_NLSP Protocol = "I_NLSP"
	// Experimental.
	Protocol_SWIPE Protocol = "SWIPE"
	// Experimental.
	Protocol_NARP Protocol = "NARP"
	// Experimental.
	Protocol_MOBILE Protocol = "MOBILE"
	// Experimental.
	Protocol_TLSP Protocol = "TLSP"
	// Experimental.
	Protocol_SKIP Protocol = "SKIP"
	// Experimental.
	Protocol_ICMPV6 Protocol = "ICMPV6"
	// Experimental.
	Protocol_IPV6_NONXT Protocol = "IPV6_NONXT"
	// Experimental.
	Protocol_IPV6_OPTS Protocol = "IPV6_OPTS"
	// Experimental.
	Protocol_CFTP Protocol = "CFTP"
	// Experimental.
	Protocol_ANY_LOCAL Protocol = "ANY_LOCAL"
	// Experimental.
	Protocol_SAT_EXPAK Protocol = "SAT_EXPAK"
	// Experimental.
	Protocol_KRYPTOLAN Protocol = "KRYPTOLAN"
	// Experimental.
	Protocol_RVD Protocol = "RVD"
	// Experimental.
	Protocol_IPPC Protocol = "IPPC"
	// Experimental.
	Protocol_ANY_DFS Protocol = "ANY_DFS"
	// Experimental.
	Protocol_SAT_MON Protocol = "SAT_MON"
	// Experimental.
	Protocol_VISA Protocol = "VISA"
	// Experimental.
	Protocol_IPCV Protocol = "IPCV"
	// Experimental.
	Protocol_CPNX Protocol = "CPNX"
	// Experimental.
	Protocol_CPHB Protocol = "CPHB"
	// Experimental.
	Protocol_WSN Protocol = "WSN"
	// Experimental.
	Protocol_PVP Protocol = "PVP"
	// Experimental.
	Protocol_BR_SAT_MON Protocol = "BR_SAT_MON"
	// Experimental.
	Protocol_SUN_ND Protocol = "SUN_ND"
	// Experimental.
	Protocol_WB_MON Protocol = "WB_MON"
	// Experimental.
	Protocol_WB_EXPAK Protocol = "WB_EXPAK"
	// Experimental.
	Protocol_ISO_IP Protocol = "ISO_IP"
	// Experimental.
	Protocol_VMTP Protocol = "VMTP"
	// Experimental.
	Protocol_SECURE_VMTP Protocol = "SECURE_VMTP"
	// Experimental.
	Protocol_VINES Protocol = "VINES"
	// Experimental.
	Protocol_TTP Protocol = "TTP"
	// Experimental.
	Protocol_IPTM Protocol = "IPTM"
	// Experimental.
	Protocol_NSFNET_IGP Protocol = "NSFNET_IGP"
	// Experimental.
	Protocol_DGP Protocol = "DGP"
	// Experimental.
	Protocol_TCF Protocol = "TCF"
	// Experimental.
	Protocol_EIGRP Protocol = "EIGRP"
	// Experimental.
	Protocol_OSPFIGP Protocol = "OSPFIGP"
	// Experimental.
	Protocol_SPRITE_RPC Protocol = "SPRITE_RPC"
	// Experimental.
	Protocol_LARP Protocol = "LARP"
	// Experimental.
	Protocol_MTP Protocol = "MTP"
	// Experimental.
	Protocol_AX_25 Protocol = "AX_25"
	// Experimental.
	Protocol_IPIP Protocol = "IPIP"
	// Experimental.
	Protocol_MICP Protocol = "MICP"
	// Experimental.
	Protocol_SCC_SP Protocol = "SCC_SP"
	// Experimental.
	Protocol_ETHERIP Protocol = "ETHERIP"
	// Experimental.
	Protocol_ENCAP Protocol = "ENCAP"
	// Experimental.
	Protocol_ANY_ENC Protocol = "ANY_ENC"
	// Experimental.
	Protocol_GMTP Protocol = "GMTP"
	// Experimental.
	Protocol_IFMP Protocol = "IFMP"
	// Experimental.
	Protocol_PNNI Protocol = "PNNI"
	// Experimental.
	Protocol_PIM Protocol = "PIM"
	// Experimental.
	Protocol_ARIS Protocol = "ARIS"
	// Experimental.
	Protocol_SCPS Protocol = "SCPS"
	// Experimental.
	Protocol_QNX Protocol = "QNX"
	// Experimental.
	Protocol_A_N Protocol = "A_N"
	// Experimental.
	Protocol_IPCOMP Protocol = "IPCOMP"
	// Experimental.
	Protocol_SNP Protocol = "SNP"
	// Experimental.
	Protocol_COMPAQ_PEER Protocol = "COMPAQ_PEER"
	// Experimental.
	Protocol_IPX_IN_IP Protocol = "IPX_IN_IP"
	// Experimental.
	Protocol_VRRP Protocol = "VRRP"
	// Experimental.
	Protocol_PGM Protocol = "PGM"
	// Experimental.
	Protocol_ANY_0_HOP Protocol = "ANY_0_HOP"
	// Experimental.
	Protocol_L2_T_P Protocol = "L2_T_P"
	// Experimental.
	Protocol_DDX Protocol = "DDX"
	// Experimental.
	Protocol_IATP Protocol = "IATP"
	// Experimental.
	Protocol_STP Protocol = "STP"
	// Experimental.
	Protocol_SRP Protocol = "SRP"
	// Experimental.
	Protocol_UTI Protocol = "UTI"
	// Experimental.
	Protocol_SMP Protocol = "SMP"
	// Experimental.
	Protocol_SM Protocol = "SM"
	// Experimental.
	Protocol_PTP Protocol = "PTP"
	// Experimental.
	Protocol_ISIS_IPV4 Protocol = "ISIS_IPV4"
	// Experimental.
	Protocol_FIRE Protocol = "FIRE"
	// Experimental.
	Protocol_CRTP Protocol = "CRTP"
	// Experimental.
	Protocol_CRUDP Protocol = "CRUDP"
	// Experimental.
	Protocol_SSCOPMCE Protocol = "SSCOPMCE"
	// Experimental.
	Protocol_IPLT Protocol = "IPLT"
	// Experimental.
	Protocol_SPS Protocol = "SPS"
	// Experimental.
	Protocol_PIPE Protocol = "PIPE"
	// Experimental.
	Protocol_SCTP Protocol = "SCTP"
	// Experimental.
	Protocol_FC Protocol = "FC"
	// Experimental.
	Protocol_RSVP_E2E_IGNORE Protocol = "RSVP_E2E_IGNORE"
	// Experimental.
	Protocol_MOBILITY_HEADER Protocol = "MOBILITY_HEADER"
	// Experimental.
	Protocol_UDPLITE Protocol = "UDPLITE"
	// Experimental.
	Protocol_MPLS_IN_IP Protocol = "MPLS_IN_IP"
	// Experimental.
	Protocol_MANET Protocol = "MANET"
	// Experimental.
	Protocol_HIP Protocol = "HIP"
	// Experimental.
	Protocol_SHIM6 Protocol = "SHIM6"
	// Experimental.
	Protocol_WESP Protocol = "WESP"
	// Experimental.
	Protocol_ROHC Protocol = "ROHC"
	// Experimental.
	Protocol_ETHERNET Protocol = "ETHERNET"
	// Experimental.
	Protocol_EXPERIMENT_1 Protocol = "EXPERIMENT_1"
	// Experimental.
	Protocol_EXPERIMENT_2 Protocol = "EXPERIMENT_2"
	// Experimental.
	Protocol_RESERVED Protocol = "RESERVED"
)


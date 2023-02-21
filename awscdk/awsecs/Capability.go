package awsecs


// A Linux capability.
type Capability string

const (
	Capability_ALL Capability = "ALL"
	Capability_AUDIT_CONTROL Capability = "AUDIT_CONTROL"
	Capability_AUDIT_WRITE Capability = "AUDIT_WRITE"
	Capability_BLOCK_SUSPEND Capability = "BLOCK_SUSPEND"
	Capability_CHOWN Capability = "CHOWN"
	Capability_DAC_OVERRIDE Capability = "DAC_OVERRIDE"
	Capability_DAC_READ_SEARCH Capability = "DAC_READ_SEARCH"
	Capability_FOWNER Capability = "FOWNER"
	Capability_FSETID Capability = "FSETID"
	Capability_IPC_LOCK Capability = "IPC_LOCK"
	Capability_IPC_OWNER Capability = "IPC_OWNER"
	Capability_KILL Capability = "KILL"
	Capability_LEASE Capability = "LEASE"
	Capability_LINUX_IMMUTABLE Capability = "LINUX_IMMUTABLE"
	Capability_MAC_ADMIN Capability = "MAC_ADMIN"
	Capability_MAC_OVERRIDE Capability = "MAC_OVERRIDE"
	Capability_MKNOD Capability = "MKNOD"
	Capability_NET_ADMIN Capability = "NET_ADMIN"
	Capability_NET_BIND_SERVICE Capability = "NET_BIND_SERVICE"
	Capability_NET_BROADCAST Capability = "NET_BROADCAST"
	Capability_NET_RAW Capability = "NET_RAW"
	Capability_SETFCAP Capability = "SETFCAP"
	Capability_SETGID Capability = "SETGID"
	Capability_SETPCAP Capability = "SETPCAP"
	Capability_SETUID Capability = "SETUID"
	Capability_SYS_ADMIN Capability = "SYS_ADMIN"
	Capability_SYS_BOOT Capability = "SYS_BOOT"
	Capability_SYS_CHROOT Capability = "SYS_CHROOT"
	Capability_SYS_MODULE Capability = "SYS_MODULE"
	Capability_SYS_NICE Capability = "SYS_NICE"
	Capability_SYS_PACCT Capability = "SYS_PACCT"
	Capability_SYS_PTRACE Capability = "SYS_PTRACE"
	Capability_SYS_RAWIO Capability = "SYS_RAWIO"
	Capability_SYS_RESOURCE Capability = "SYS_RESOURCE"
	Capability_SYS_TIME Capability = "SYS_TIME"
	Capability_SYS_TTY_CONFIG Capability = "SYS_TTY_CONFIG"
	Capability_SYSLOG Capability = "SYSLOG"
	Capability_WAKE_ALARM Capability = "WAKE_ALARM"
)


package common

// NameReg regex
const (
	NameReg      = "^[a-zA-Z/0-9_.-]{2,64}$"
	PwdReg       = "^[a-zA-Z0-9@#!~&]{6,16}$"
	DescReg      = "^[\\w-_., ;/：:?？=&\\p{Han}，。；【】\n]{0,1000}$"
	AliasNameReg = "^[\\w-_.,\\p{Han}]{4,64}$"

	VersionTagReg = "^[a-zA-Z/0-9.-]{4,32}$"
	// VersionReg = "^[a-f0-9]{7}$"
	VersionReg = "^[a-zA-Z/0-9.-]{4,32}$"

	IPRegex            = "(\\d|[1-9]\\d|1\\d\\d|2[0-4]\\d|25[0-5])"
	ProjectManifestReg = "^(http|https|ssh)://[0-9A-Za-z./:#@]{4,64}$"
	EmailReg           = "^([a-zA-Z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$"
)

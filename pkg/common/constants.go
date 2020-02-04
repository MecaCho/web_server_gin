package common

// ContextKey ...
type ContextKey string

// RequestUserNameContextKey holds the key to store the request parameters
const (
	RequestUserNameContextKey   ContextKey = "user_name"
	RequestRolesContextKey      ContextKey = "roles"
	RequestGroupsContextKey     ContextKey = "groups"
	RequestURLContextKey        ContextKey = "url"
	RequestActionContextKey     ContextKey = "action"
	CasbinEnforcer              ContextKey = "enforcer"
	CasbinEnforcerVersion       ContextKey = "casbin_version"
	AuditDBHandleContextKey     ContextKey = "audit_db"
	RequestResourceIDContextKey ContextKey = "resource_id"
	RequestDurationContextKey   ContextKey = "request_duration"
	RequestRequestIDContextKey  ContextKey = "request_id"
	RequestAuditSwitch          ContextKey = "request_audit"
)

// UserUnauthorized unauthorized
const (
	UserUnauthorized = iota + 40100
	UserUnauthorizedExpired
	UserUnauthorizedInvilid
	UserUnauthorizedNotFound
	UserNotFound
	RoleNotFound
	UserDeleteConflict
	HostNotFound
)

// RoleUnauthorized ...
const (
	RoleUnauthorized = iota + 40300
)

// invalid request body parameters status codes
const (
	NameDuplicate = iota + 10000
	InvalidJSONFormat
	StringMissing
	IntMissing
	ArrayMissing
	EnumsNotSupport
	DescInvalid
	NameInvalid
	AliasNameINvalid
	ExcessiveBindingNum
	RedundantKey
	VersionInvalid
	EmailInvalid
	PwdInvalid
	RequestBodyInvalid
	IPInvalid
)

// InternalServerError resp status code
const (
	DBInternalError = iota
)

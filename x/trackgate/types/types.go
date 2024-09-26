package types

// StructDef represents a structure definition
type StructDef struct {
	Fields map[string]interface{} // Changed to interface{} to allow nested definitions
}

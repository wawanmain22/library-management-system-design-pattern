package prototype

import "fmt"

// PrototypeManager manages prototypes for easy cloning
type PrototypeManager struct {
	prototypes map[string]Prototype
}

// NewPrototypeManager creates a new PrototypeManager
func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Prototype),
	}
}

// RegisterPrototype registers a prototype with a key
func (pm *PrototypeManager) RegisterPrototype(key string, proto Prototype) {
	pm.prototypes[key] = proto
}

// GetPrototype retrieves and clones a prototype by key
func (pm *PrototypeManager) GetPrototype(key string) (Prototype, error) {
	proto, exists := pm.prototypes[key]
	if !exists {
		return nil, fmt.Errorf("prototype with key '%s' not found", key)
	}
	return proto.Clone(), nil
}

// ListPrototypes returns all registered prototype keys
func (pm *PrototypeManager) ListPrototypes() []string {
	keys := make([]string, 0, len(pm.prototypes))
	for k := range pm.prototypes {
		keys = append(keys, k)
	}
	return keys
}

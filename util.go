package advantageair

func setRecursively(dest map[string]any, value any, keys ...string) map[string]any {
	if len(keys) < 1 {
		return dest
	}

	if len(keys) == 1 {
		dest[keys[0]] = value
		return dest
	}

	child, ok := dest[keys[0]]
	if !ok {
		child = make(map[string]any)
	}

	sub, ok := child.(map[string]any)
	if !ok {
		return dest
	}

	dest[keys[0]] = setRecursively(sub, value, keys[1:]...)
	return dest
}

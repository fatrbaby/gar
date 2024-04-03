package ent

func (k *Keyword) IntoString() string {
	if len(k.Word) > 0 {
		return k.Field + "\001" + k.Word
	}

	return ""
}

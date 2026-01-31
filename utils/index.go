package utils

type Index map[string][]int

// Add indexes any slice of Searchable items
func (idx Index) Add(docs []Searchable) {
	for _, doc := range docs {
		for _, token := range Analyze(doc.GetSearchText()) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.GetID() {
				continue
			}
			idx[token] = append(ids, doc.GetID())
		}
	}
}

func Intersection(a []int, b []int) []int {
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	r := make([]int, 0, maxLen)
	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func (idx Index) Search(text string) []int {
	var r []int
	for _, token := range Analyze(text) {
		if ids, ok := idx[token]; ok {
			if r == nil {
				r = ids
			} else {
				r = Intersection(r, ids)
			}
		} else {
			return nil
		}
	}
	return r
}

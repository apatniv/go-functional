package template

type FilterIter struct {
	iter   Iter
	filter filterFunc
}

func Filter(iter Iter, filter filterFunc) FilterIter {
	return FilterIter{iter, filter}
}

func (iter FilterIter) Next() Result {
	for {
		option := iter.iter.Next()
		if !option.Present() || iter.filter(fromT(option.Value())) {
			return option
		}
	}
}

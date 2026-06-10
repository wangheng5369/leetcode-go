package graph

type UnitConverter struct {
	graph map[string]map[string]float64
}

func NewUnitConverter() *UnitConverter {
	return &UnitConverter{graph: make(map[string]map[string]float64)}
}

func (uc *UnitConverter) AddFact(fromUnit string, toUnit string, value float64) {
	if _, ok := uc.graph[fromUnit]; !ok {
		uc.graph[fromUnit] = make(map[string]float64)
	}
	uc.graph[fromUnit][toUnit] = value
	if _, ok := uc.graph[toUnit]; !ok {
		uc.graph[toUnit] = make(map[string]float64)
	}
	uc.graph[toUnit][fromUnit] = 1.0 / value
}

func (uc *UnitConverter) Convert(value float64, fromUnit string, toUnit string) (float64, bool) {
	if fromUnit == toUnit {
		return value, true
	}

	if _, ok := uc.graph[fromUnit]; !ok {
		return 0, false
	}

	type queueItem struct {
		unit  string
		value float64
	}
	queue := []queueItem{{fromUnit, value}}
	visited := make(map[string]bool)
	visited[fromUnit] = true
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for k, v := range uc.graph[item.unit] {
			if k == toUnit {
				return value * item.value * v, true
			}
			if !visited[k] {
				visited[k] = true
				queue = append(queue, queueItem{k, item.value * v})
			}
		}
	}
	return 0, false
}

package discs

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Disc struct {
	Name   string
	Weight int
	Parent *Disc
	Kids   []*Disc
}

func (d *Disc) NameOrDefault(defaultValue string) string {
	if d != nil && d.Name != "" {
		return (*d).Name
	}
	return defaultValue
}

func (d *Disc) TotalWeight() int {
	totalWeight := d.Weight
	for _, child := range d.Kids {
		totalWeight += child.TotalWeight()
	}
	return totalWeight
}

func (d *Disc) IsBalanced() bool {
	if d.Kids != nil && len(d.Kids) > 0 {
		first := d.Kids[0].TotalWeight()
		for _, c := range d.Kids {
			if c.TotalWeight() != first {
				return false
			}
		}
	}
	// either no kids or all are balanced - am assuming 0 kids == balanced
	return true
}

func (d *Disc) Discrepancy() (int, error) {
	if d.Kids != nil && len(d.Kids) > 0 {
		weights := make(map[int]int, 2)
		for _, c := range d.Kids {
			weights[c.TotalWeight()] += 1
		}
		different, common := 0, 0
		for key, value := range weights {
			if value == 1 {
				different = key
			}
			if value > 1 {
				common = key
			}
		}

		for _, c := range d.Kids {
			if c.TotalWeight() == different {
				if different > common {
					return c.Weight + (common - different), nil
				} else {
					return c.Weight - (different - common), nil
				}
			}
		}
	}
	return 0, errors.New("something went wrong calculating discrepancy")
}

func ParseStrings(data []string) map[string]*Disc {
	re := regexp.MustCompile(`([A-Za-z]+) \((\d+)\)`)
	reKids := regexp.MustCompile(`.* -> ([A-Za-z, ]+)`)

	results := make(map[string]*Disc)
	for i := 0; i < len(data); i++ {
		parts := re.FindStringSubmatch(data[i])
		kids := reKids.FindStringSubmatch(data[i])

		if len(parts) > 0 {
			name := parts[1]
			weight, _ := strconv.Atoi(parts[2])
			children := ""

			if len(kids) > 0 {
				children = kids[1]
			}
			upsertDisc(&results, name, weight, children)
		}
	}
	return results
}

func upsertDisc(discs *map[string]*Disc, name string, weight int, children string) {
	existing, shouldUpdate := (*discs)[name]
	if shouldUpdate { // if exists, update weight
		existing.Weight = weight
		(*discs)[name] = existing
		upsertChild(discs, existing, children)
	} else { // doesn't exist, create
		newDisc := Disc{Name: name, Weight: weight, Parent: nil}
		(*discs)[name] = &newDisc
		upsertChild(discs, &newDisc, children)
	}
}

func upsertChild(discs *map[string]*Disc, newDisc *Disc, children string) {
	for _, child := range strings.Split(children, ", ") {
		if len(child) > 0 {
			val, ok := (*discs)[child]
			if ok { // if exists, set parent
				val.Parent = newDisc
				newDisc.Kids = append(newDisc.Kids, val)
			} else { // doesn't exist, set name and parent
				childDisc := &Disc{Name: child, Weight: 0, Parent: newDisc}
				(*discs)[child] = childDisc
				newDisc.Kids = append(newDisc.Kids, childDisc)
			}
		}
	}
}

// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package deck

import "sort"

// A card set; sortable by value.
type ValueSet []Card

func (s ValueSet) Sort()              { sort.Sort(s) }
func (s ValueSet) Len() int           { return len(s) }
func (s ValueSet) Less(i, j int) bool { return s[i].Value() < s[j].Value() }
func (s ValueSet) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// A card set; sortable by suit.
type SuitSet []Card

func (s SuitSet) Sort()              { sort.Sort(s) }
func (s SuitSet) Len() int           { return len(s) }
func (s SuitSet) Less(i, j int) bool { return s[i].Suit() < s[j].Suit() }
func (s SuitSet) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

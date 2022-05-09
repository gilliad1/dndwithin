package rules

import (
	"fmt"

	"github.com/jackc/pgx/v4"
)

type Spell struct {
	ID              int
	Name            string
	School          string
	SubSchool       string
	Source          string
	Description     string
	Effect          string
	Range           string
	Components      string
	Area            string
	Target          string
	Duration        string
	SavingThrow     string
	SpellResistance bool
	Is30            bool
	IsEvil          bool
}

//Given a Rows query on the spell table, return the corresponding []Spell
func SpellFromQueryRows(rows pgx.Rows) ([]Spell, error) {
	retSpell := make([]Spell, 0, 25)
	defer rows.Close()
	for rows.Next() {
		thisSpell := Spell{}
		if rows.Err() != nil {
			return retSpell, fmt.Errorf("ERROR PROCESSING ROW: %w", rows.Err())
		}
		if err := rows.Scan(
			&thisSpell.ID,
			&thisSpell.Name,
			&thisSpell.School,
			&thisSpell.SubSchool,
			&thisSpell.Source,
			&thisSpell.Description,
			&thisSpell.Effect,
			&thisSpell.Range,
			&thisSpell.Components,
			&thisSpell.Area,
			&thisSpell.Target,
			&thisSpell.Duration,
			&thisSpell.SavingThrow,
			&thisSpell.SpellResistance,
			&thisSpell.Is30,
			&thisSpell.IsEvil,
		); err != nil {
			return retSpell, fmt.Errorf("ERROR TRANSLATING SPELL QUERY TO STRUCT: %w", err)
		}
		if _, present := sourceName[thisSpell.Source]; present {
			thisSpell.Source = sourceName[thisSpell.Source]
		}
		retSpell = append(retSpell, thisSpell)
	}
	return retSpell, nil
}

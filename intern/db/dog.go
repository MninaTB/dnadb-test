package db

import (
  "github.de/mninatb/dnadb/intern/model"
)

const (
  dogByID = `
SELECT
    id, name, birthdate, race, color, mother, father
FROM prices WHERE id = ?
`
)

type DogRepo struct {
}

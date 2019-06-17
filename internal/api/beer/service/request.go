package service

import (
	"errors"
	"fmt"
	"github.com/Maoltr/alco/model"
	"github.com/satori/go.uuid"
)

// CreateBeerRequest holds data for creating beer
type CreateBeerRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Strength    uint   `json:"strength"`
	UserID      string `json:"user_id"`
}

// Converts create beer request into beer model
func (c CreateBeerRequest) ConvertToBeer() (model.Beer, error) {
	var result model.Beer
	if len(c.Name) < 6 || len(c.Name) > 100 {
		return result, errors.New(fmt.Sprintf("beer name must be from 6 to 100 chars, you provided: %d", len(c.Name)))
	}

	if len(c.Description) < 20 || len(c.Description) > 1000 {
		return result, errors.New(fmt.Sprintf("beer description must be from 20 to 1-00 chars, you provided: %d", len(c.Description)))
	}

	if c.Strength > 100 {
		return result, errors.New(fmt.Sprintf("beer strength can not be more than 100 percents, you provided: %d", c.Strength))
	}

	result.ID = uuid.NewV4().String()
	result.Strength = c.Strength
	result.Name = c.Name
	result.Description = c.Description
	result.AddedBy = c.UserID
	return result, nil
}

// UpdateBeerRequest holds data for updating beer
type UpdateBeerRequest struct {
	ID          string  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Strength    *int    `json:"strength"`
}

func (u UpdateBeerRequest) IsValid() error {
	var atLeastOneChanged bool
	if u.Name != nil {
		if len(*u.Name) < 6 || len(*u.Name) > 100 {
			return errors.New(fmt.Sprintf("beer name must be from 6 to 100 chars, you provided: %d", len(c.Name)))
		}
		atLeastOneChanged = true
	}

	if u.Description != nil {
		if len(*u.Description) < 20 || len(*u.Description) > 1000 {
			return errors.New(fmt.Sprintf("beer description must be from 20 to 1-00 chars, you provided: %d", len(c.Description)))
		}
		atLeastOneChanged = true
	}

	if u.Strength != nil {
		if *u.Strength > 100 {
			return errors.New(fmt.Sprintf("beer strength can not be more than 100 percents, you provided: %d", c.Strength))
		}
		atLeastOneChanged = true
	}

	if !atLeastOneChanged {
		return errors.New("at least one field must be updated")
	}

	return nil
}
